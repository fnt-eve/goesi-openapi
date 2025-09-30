//go:generate ./scripts/generate_scopes.sh
//go:generate ./scripts/update_compatibility_date.sh

package goesi

import (
	"context"
	"net/http"

	"github.com/fnt-eve/goesi-openapi/esi"
	"golang.org/x/oauth2"
)

const (
	// ESI compatibility date for API versioning
	ESICompatibilityDate = "2020-01-01"
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

	// Set compatibility date
	compatibilityDate := options.CompatibilityDate
	if compatibilityDate == "" {
		compatibilityDate = ESICompatibilityDate
	}
	cfg.AddDefaultHeader("X-Compatibility-Date", compatibilityDate)

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

	// Add required ESI headers
	cfg.AddDefaultHeader("X-Compatibility-Date", ESICompatibilityDate)

	return esi.NewAPIClient(cfg)
}
