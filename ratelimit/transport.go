package ratelimit

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Transport implements http.RoundTripper with rate limiting
type Transport struct {
	Base  http.RoundTripper
	Store *BucketStore
}

// NewTransport creates a new rate limiting transport
func NewTransport(base http.RoundTripper) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &Transport{
		Base:  base,
		Store: NewBucketStore(),
	}
}

// RoundTrip executes a single HTTP transaction
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	routeConfig := t.matchRoute(req)

	if routeConfig == nil {
		return t.Base.RoundTrip(req)
	}

	key := t.getBucketKey(req, routeConfig)
	bucket := t.Store.GetBucket(key, routeConfig.Limit, routeConfig.Window)

	// Default cost is 2 tokens per request (standard success cost)
	if err := bucket.Wait(req.Context(), 2); err != nil {
		return nil, err
	}

	resp, err := t.Base.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if remainStr := resp.Header.Get("X-RateLimit-Remaining"); remainStr != "" {
		if remain, err := strconv.Atoi(remainStr); err == nil {
			bucket.Sync(remain)
		}
	}

	return resp, nil
}

func (t *Transport) matchRoute(req *http.Request) *RouteConfig {
	path := req.URL.Path
	method := req.Method

	for _, rc := range RouteMapping {
		if rc.Method == method && rc.PathRegex.MatchString(path) {
			return &rc
		}
	}
	return nil
}

func (t *Transport) getBucketKey(req *http.Request, routeConfig *RouteConfig) BucketKey {
	azp, sub := t.extractClaims(req)

	var subject string
	if routeConfig.Auth {
		// Authenticated route: use azp:sub
		if azp != "" && sub != "" {
			subject = azp + ":" + sub
		} else if sub != "" {
			subject = sub
		} else {
			subject = "public"
		}
	} else {
		// Public route: use public:azp if available, else public
		if azp != "" {
			subject = "public:" + azp
		} else {
			subject = "public"
		}
	}

	return BucketKey(routeConfig.Group + ":" + subject)
}

func (t *Transport) extractClaims(req *http.Request) (string, string) {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return "", ""
	}

	parts := strings.Fields(auth)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ""
	}

	tokenString := parts[1]

	// Parse JWT without verification (we just need ID for rate limiting)
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		sub, _ := claims["sub"].(string)
		azp, _ := claims["azp"].(string)
		return azp, sub
	}

	return "", ""
}
