//go:generate ./scripts/generate_scopes.sh
//go:generate ./scripts/generate_ratelimit_map.sh

package goesi

import (
	"context"
	"net/http"

	"github.com/fnt-eve/goesi-openapi/esi"
	"golang.org/x/oauth2"
)

// Context keys for ESI authentication
var (
	// ContextOAuth2 is the context key for GoESI authentication. Pass a tokenSource with this key to a context for an ESI API Call
	ContextOAuth2 = esi.ContextOAuth2
)

// NewAuthenticatedESIClient creates a new ESI API client with OAuth2 authentication
func NewAuthenticatedESIClient(ctx context.Context, config *Config, token *oauth2.Token, userAgent string) *esi.APIClient {
	httpClient := config.Client(ctx, token)
	return newESIClient(httpClient, userAgent)
}

// NewPublicESIClient creates a new ESI API client for public endpoints (no authentication required)
func NewPublicESIClient(userAgent string) *esi.APIClient {
	return newESIClient(http.DefaultClient, userAgent)
}

// ClientOptions holds configuration options for ESI client creation
type ClientOptions struct {
	UserAgent         string
	CompatibilityDate string
	BaseURL           string
}

// NewESIClientWithOptions creates a new ESI API client with custom options
func NewESIClientWithOptions(httpClient *http.Client, options ClientOptions) *esi.APIClient {
	cfg := esi.NewConfiguration()
	cfg.HTTPClient = httpClient

	// Set user agent
	if options.UserAgent != "" {
		cfg.UserAgent = options.UserAgent
	}

	// Set compatibility date - uses the one from Configuration if not specified
	if options.CompatibilityDate != "" {
		cfg.CompatibilityDate = options.CompatibilityDate
	}

	// Set base URL if provided
	if options.BaseURL != "" {
		cfg.Servers = esi.ServerConfigurations{
			{
				URL: options.BaseURL,
			},
		}
	}

	return esi.NewAPIClient(cfg)
}

// newESIClient creates a new ESI API client with proper configuration
func newESIClient(httpClient *http.Client, userAgent string) *esi.APIClient {
	cfg := esi.NewConfiguration()
	cfg.HTTPClient = httpClient
	cfg.UserAgent = userAgent

	return esi.NewAPIClient(cfg)
}
