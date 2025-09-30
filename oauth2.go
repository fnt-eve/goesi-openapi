package goesi

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

const (
	// ESI OAuth2 endpoints
	esiAuthURL  = "https://login.eveonline.com/v2/oauth/authorize"
	esiTokenURL = "https://login.eveonline.com/v2/oauth/token"
	esiJWKSURL  = "https://login.eveonline.com/oauth/jwks"
)

var (
	ErrInvalidToken = errors.New("invalid or expired token")
	ErrInvalidState = errors.New("invalid OAuth2 state parameter")
	ErrMissingCode  = errors.New("missing authorization code")
	ErrTokenRefresh = errors.New("failed to refresh token")
)

// Config holds the OAuth2 configuration for ESI authentication
type Config struct {
	ClientID     string
	RedirectURL  string
	Scopes       []string
	oauth2Config *oauth2.Config
	verifier     string
	challenge    string
	JWTKeyFunc   jwt.Keyfunc
}

// ESIJWTClaims represents the claims in an ESI JWT token with embedded standard claims
type ESIJWTClaims struct {
	jwt.RegisteredClaims

	// ESI-specific claims
	Scopes []string `json:"scp"`
	Name   string   `json:"name"`
	Owner  string   `json:"owner"`
	Kid    string   `json:"kid"`
	Azp    string   `json:"azp"`
	Tenant string   `json:"tenant"`
	Tier   string   `json:"tier"`
	Region string   `json:"region"`
}

// CharacterID extracts the character ID from the JWT subject
func (claims *ESIJWTClaims) CharacterID() (int32, error) {
	if claims == nil {
		return 0, errors.New("no claims available")
	}

	// Subject format is "CHARACTER:EVE:123456789"
	parts := strings.Split(claims.Subject, ":")
	if len(parts) != 3 || parts[0] != "CHARACTER" || parts[1] != "EVE" {
		return 0, errors.New("invalid subject format")
	}

	id, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid character ID: %w", err)
	}

	return int32(id), nil
}

// CharacterName returns the character name from the JWT claims
func (claims *ESIJWTClaims) CharacterName() string {
	if claims == nil {
		return ""
	}
	return claims.Name
}

// TokenScopes returns the scopes from the JWT claims
func (claims *ESIJWTClaims) TokenScopes() []string {
	if claims == nil {
		return nil
	}
	return claims.Scopes
}

// NewConfig creates a new ESI OAuth2 configuration
func NewConfig(clientID, redirectURL string, scopes []string, keyFunc *jwt.Keyfunc) (*Config, error) {
	if clientID == "" {
		return nil, errors.New("client ID is required")
	}
	if redirectURL == "" {
		return nil, errors.New("redirect URL is required")
	}

	// Generate PKCE challenge and verifier
	verifier, challenge, err := generatePKCE()
	if err != nil {
		return nil, fmt.Errorf("failed to generate PKCE: %w", err)
	}

	oauth2Config := &oauth2.Config{
		ClientID:    clientID,
		RedirectURL: redirectURL,
		Scopes:      scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  esiAuthURL,
			TokenURL: esiTokenURL,
		},
	}

	// Keyfunc must be set
	var jwtKeyFunc jwt.Keyfunc
	if keyFunc != nil {
		jwtKeyFunc = *keyFunc
	}

	return &Config{
		ClientID:     clientID,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		oauth2Config: oauth2Config,
		verifier:     verifier,
		challenge:    challenge,
		JWTKeyFunc:   jwtKeyFunc,
	}, nil
}

// AuthURL generates the OAuth2 authorization URL with PKCE
func (c *Config) AuthURL(state string) string {
	return c.oauth2Config.AuthCodeURL(state,
		oauth2.SetAuthURLParam("code_challenge", c.challenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
}

// Exchange exchanges the authorization code for an access token and parses JWT claims
func (c *Config) Exchange(ctx context.Context, code, state string, expectedState string) (*oauth2.Token, *ESIJWTClaims, error) {
	if code == "" {
		return nil, nil, ErrMissingCode
	}
	if state != expectedState {
		return nil, nil, ErrInvalidState
	}

	token, err := c.oauth2Config.Exchange(ctx, code,
		oauth2.SetAuthURLParam("code_verifier", c.verifier),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	// Parse and validate JWT claims - this is required for ESI tokens
	claims, err := c.parseTokenInfo(token.AccessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse token claims: %w", err)
	}

	return token, claims, nil
}

// RefreshToken refreshes an expired access token and parses the new JWT claims
func (c *Config) RefreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, *ESIJWTClaims, error) {
	if token.RefreshToken == "" {
		return nil, nil, ErrTokenRefresh
	}

	newToken, err := c.oauth2Config.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrTokenRefresh, err)
	}

	// Parse and validate JWT claims from the refreshed token
	claims, err := c.parseTokenInfo(newToken.AccessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse refreshed token claims: %w", err)
	}

	return newToken, claims, nil
}

// Client returns an HTTP client that automatically includes authentication
func (c *Config) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return c.oauth2Config.Client(ctx, token)
}

// IsExpired checks if the token is expired or will expire soon
func IsExpired(token *oauth2.Token) bool {
	if token == nil {
		return true
	}
	// Consider token expired if it expires within the next 30 seconds
	return token.Expiry.Before(time.Now().Add(30 * time.Second))
}

// ESIDefaultKeyfunc creates a default keyfunc that fetches ESI JWT signing keys from JWKS
func ESIDefaultKeyfunc(ctx context.Context) (jwt.Keyfunc, error) {
	jwks, err := keyfunc.NewDefaultCtx(ctx, []string{esiJWKSURL})
	if err != nil {
		return nil, fmt.Errorf("failed to create JWKS from resource at the given URL: %w", err)
	}

	return jwks.Keyfunc, nil
}

// generatePKCE generates PKCE code verifier and challenge
func generatePKCE() (verifier, challenge string, err error) {
	// Generate 43-128 character random string for verifier
	bytes := make([]byte, 96) // 96 bytes = 128 base64url characters
	if _, err := rand.Read(bytes); err != nil {
		return "", "", err
	}

	verifier = base64.RawURLEncoding.EncodeToString(bytes)

	// Create challenge as SHA256 hash of verifier
	hash := sha256.Sum256([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(hash[:])

	return verifier, challenge, nil
}

// ParseClaims extracts and validates claims from an oauth2.Token's access token
func (c *Config) ParseClaims(token *oauth2.Token) (*ESIJWTClaims, error) {
	if token == nil {
		return nil, errors.New("token is nil")
	}
	return c.parseTokenInfo(token.AccessToken)
}

// parseTokenInfo extracts and validates character information from the JWT access token
func (c *Config) parseTokenInfo(accessToken string) (*ESIJWTClaims, error) {
	if accessToken == "" {
		return nil, errors.New("no access token available")
	}

	// Keyfunc must be set
	keyFunc := c.JWTKeyFunc
	if keyFunc == nil {
		return nil, errors.New("no keyfunc available for JWT validation")
	}

	// Parse and validate JWT token with signature verification
	jwtToken, err := jwt.ParseWithClaims(accessToken, &ESIJWTClaims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to validate JWT token: %w", err)
	}

	// Extract the custom claims
	if claims, ok := jwtToken.Claims.(*ESIJWTClaims); ok {
		return claims, nil
	}

	return nil, errors.New("failed to extract ESI JWT claims")
}

// TokenToJSON helper function to convert a token to a storable format.
func TokenToJSON(token *oauth2.Token) (string, error) {
	if d, err := json.Marshal(token); err != nil {
		return "", err
	} else {
		return string(d), nil
	}
}

// TokenFromJSON helper function to convert stored JSON to a token.
func TokenFromJSON(jsonStr string) (*oauth2.Token, error) {
	var token oauth2.Token
	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
		return nil, err
	}
	return &token, nil
}
