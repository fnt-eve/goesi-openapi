package goesi

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestNewPublicESIClient(t *testing.T) {
	userAgent := "TestAgent/1.0"
	client := NewPublicESIClient(userAgent)

	require.NotNil(t, client, "NewPublicESIClient should not return nil")

	cfg := client.GetConfig()
	require.Equal(t, userAgent, cfg.UserAgent, "UserAgent should match")
	require.NotEmpty(t, cfg.DefaultHeader, "Default headers should be set")
	require.Contains(t, cfg.DefaultHeader, "X-Compatibility-Date", "X-Compatibility-Date header should be set")
	require.Equal(t, ESICompatibilityDate, cfg.DefaultHeader["X-Compatibility-Date"], "Compatibility date should match")
}

func TestNewAuthenticatedESIClient(t *testing.T) {
	ctx := context.Background()
	userAgent := "TestAgent/1.0"

	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost:8080/callback", []string{"scope1"}, &keyFunc)
	require.NoError(t, err, "NewConfig should not return error")
	require.NotNil(t, config, "Config should not be nil")

	token := &oauth2.Token{
		AccessToken:  "test-access-token",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	client := NewAuthenticatedESIClient(ctx, config, token, userAgent)
	require.NotNil(t, client, "NewAuthenticatedESIClient should not return nil")

	cfg := client.GetConfig()
	require.Equal(t, userAgent, cfg.UserAgent, "UserAgent should match")
	require.NotNil(t, cfg.HTTPClient, "HTTPClient should be set")
}

func TestNewESIClientWithOptions_DefaultCompatibilityDate(t *testing.T) {
	options := ClientOptions{
		UserAgent: "TestAgent/1.0",
	}

	client := NewESIClientWithOptions(http.DefaultClient, options)
	require.NotNil(t, client, "NewESIClientWithOptions should not return nil")

	cfg := client.GetConfig()
	require.Contains(t, cfg.DefaultHeader, "X-Compatibility-Date", "X-Compatibility-Date header should be set")
	require.Equal(t, ESICompatibilityDate, cfg.DefaultHeader["X-Compatibility-Date"], "Should use default compatibility date")
}

func TestNewESIClientWithOptions_CustomCompatibilityDate(t *testing.T) {
	customDate := "2021-01-01"
	options := ClientOptions{
		UserAgent:         "TestAgent/1.0",
		CompatibilityDate: customDate,
	}

	client := NewESIClientWithOptions(http.DefaultClient, options)
	require.NotNil(t, client, "NewESIClientWithOptions should not return nil")

	cfg := client.GetConfig()
	require.Contains(t, cfg.DefaultHeader, "X-Compatibility-Date", "X-Compatibility-Date header should be set")
	require.Equal(t, customDate, cfg.DefaultHeader["X-Compatibility-Date"], "Should use custom compatibility date")
}

func TestNewESIClientWithOptions_CustomBaseURL(t *testing.T) {
	customURL := "https://custom.api.example.com"
	options := ClientOptions{
		UserAgent: "TestAgent/1.0",
		BaseURL:   customURL,
	}

	client := NewESIClientWithOptions(http.DefaultClient, options)
	require.NotNil(t, client, "NewESIClientWithOptions should not return nil")

	cfg := client.GetConfig()
	require.NotEmpty(t, cfg.Servers, "Servers should be configured")
	require.Equal(t, customURL, cfg.Servers[0].URL, "Base URL should match")
}

func TestNewESIClientWithOptions_AllOptions(t *testing.T) {
	customDate := "2022-06-15"
	customURL := "https://test.api.example.com"
	customAgent := "CustomAgent/2.0"

	options := ClientOptions{
		UserAgent:         customAgent,
		CompatibilityDate: customDate,
		BaseURL:           customURL,
	}

	client := NewESIClientWithOptions(http.DefaultClient, options)
	require.NotNil(t, client, "NewESIClientWithOptions should not return nil")

	cfg := client.GetConfig()
	require.Equal(t, customAgent, cfg.UserAgent, "UserAgent should match")
	require.Equal(t, customDate, cfg.DefaultHeader["X-Compatibility-Date"], "Compatibility date should match")
	require.NotEmpty(t, cfg.Servers, "Servers should be configured")
	require.Equal(t, customURL, cfg.Servers[0].URL, "Base URL should match")
}

func TestNewESIClientWithOptions_EmptyUserAgent(t *testing.T) {
	options := ClientOptions{}

	client := NewESIClientWithOptions(http.DefaultClient, options)
	require.NotNil(t, client, "NewESIClientWithOptions should not return nil")

	cfg := client.GetConfig()
	require.NotEmpty(t, cfg.UserAgent, "UserAgent should use default from generated config")
	require.Equal(t, "OpenAPI-Generator/1.0.0/go", cfg.UserAgent, "Should use default OpenAPI generator UserAgent")
}

func mockKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret"), nil
	}
}
