package goesi

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestNewConfig_Success(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)

	require.NoError(t, err, "NewConfig should not return error")
	require.NotNil(t, config, "Config should not be nil")
	require.Equal(t, "test-client-id", config.ClientID, "ClientID should match")
	require.Equal(t, "http://localhost/callback", config.RedirectURL, "RedirectURL should match")
	require.NotEmpty(t, config.verifier, "PKCE verifier should be generated")
	require.NotEmpty(t, config.challenge, "PKCE challenge should be generated")
	require.NotNil(t, config.oauth2Config, "OAuth2 config should be initialized")
	require.NotNil(t, config.JWTKeyFunc, "JWT KeyFunc should be set")
}

func TestNewConfig_MissingClientID(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("", "http://localhost/callback", &keyFunc)

	require.Error(t, err, "NewConfig should return error for missing client ID")
	require.Nil(t, config, "Config should be nil on error")
	require.Contains(t, err.Error(), "client ID is required", "Error message should mention client ID")
}

func TestNewConfig_MissingRedirectURL(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "", &keyFunc)

	require.Error(t, err, "NewConfig should return error for missing redirect URL")
	require.Nil(t, config, "Config should be nil on error")
	require.Contains(t, err.Error(), "redirect URL is required", "Error message should mention redirect URL")
}

func TestNewConfig_NilKeyFunc(t *testing.T) {
	config, err := NewConfig("test-client-id", "http://localhost/callback", nil)

	require.NoError(t, err, "NewConfig should not error with nil keyFunc")
	require.NotNil(t, config, "Config should not be nil")
	require.Nil(t, config.JWTKeyFunc, "JWTKeyFunc should be nil")
}

func TestAuthURL(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	state := "test-state-123"
	scopes := []string{"scope1", "scope2"}
	authURL := config.AuthURL(state, scopes)

	require.NotEmpty(t, authURL, "AuthURL should not be empty")
	require.Contains(t, authURL, "test-client-id", "AuthURL should contain client ID")
	require.Contains(t, authURL, "state=test-state-123", "AuthURL should contain state")
	require.Contains(t, authURL, "code_challenge=", "AuthURL should contain code_challenge")
	require.Contains(t, authURL, "code_challenge_method=S256", "AuthURL should contain code_challenge_method")
	require.Contains(t, authURL, esiAuthURL, "AuthURL should start with ESI auth URL")
	require.Contains(t, authURL, "scope=scope1", "AuthURL should contain scope")
}

func TestESIJWTClaims_CharacterID_Success(t *testing.T) {
	claims := &ESIJWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: "CHARACTER:EVE:123456789",
		},
	}

	charID, err := claims.CharacterID()
	require.NoError(t, err, "CharacterID should not return error")
	require.Equal(t, int64(123456789), charID, "Character ID should match")
}

func TestESIJWTClaims_CharacterID_NilClaims(t *testing.T) {
	var claims *ESIJWTClaims

	charID, err := claims.CharacterID()
	require.Error(t, err, "CharacterID should return error for nil claims")
	require.Equal(t, int64(0), charID, "Character ID should be 0")
	require.Contains(t, err.Error(), "no claims available", "Error message should mention claims")
}

func TestESIJWTClaims_CharacterID_InvalidFormat(t *testing.T) {
	testCases := []struct {
		name    string
		subject string
		errMsg  string
	}{
		{
			name:    "wrong prefix",
			subject: "CORPORATION:EVE:123456",
			errMsg:  "invalid subject format",
		},
		{
			name:    "wrong game",
			subject: "CHARACTER:OTHER:123456",
			errMsg:  "invalid subject format",
		},
		{
			name:    "missing parts",
			subject: "CHARACTER:123456",
			errMsg:  "invalid subject format",
		},
		{
			name:    "non-numeric ID",
			subject: "CHARACTER:EVE:abc",
			errMsg:  "invalid character ID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			claims := &ESIJWTClaims{
				RegisteredClaims: jwt.RegisteredClaims{
					Subject: tc.subject,
				},
			}

			charID, err := claims.CharacterID()
			require.Error(t, err, "CharacterID should return error")
			require.Equal(t, int64(0), charID, "Character ID should be 0")
			require.Contains(t, err.Error(), tc.errMsg, "Error message should match")
		})
	}
}

func TestIsExpired_NilToken(t *testing.T) {
	expired := IsExpired(nil)
	require.True(t, expired, "Nil token should be considered expired")
}

func TestIsExpired_ExpiredToken(t *testing.T) {
	token := &oauth2.Token{
		AccessToken: "test-token",
		Expiry:      time.Now().Add(-1 * time.Hour),
	}

	expired := IsExpired(token)
	require.True(t, expired, "Expired token should be considered expired")
}

func TestIsExpired_ValidToken(t *testing.T) {
	token := &oauth2.Token{
		AccessToken: "test-token",
		Expiry:      time.Now().Add(1 * time.Hour),
	}

	expired := IsExpired(token)
	require.False(t, expired, "Valid token should not be considered expired")
}

func TestIsExpired_ExpiringSoon(t *testing.T) {
	token := &oauth2.Token{
		AccessToken: "test-token",
		Expiry:      time.Now().Add(15 * time.Second),
	}

	expired := IsExpired(token)
	require.True(t, expired, "Token expiring within 30 seconds should be considered expired")
}

func TestGeneratePKCE(t *testing.T) {
	verifier1, challenge1, err1 := generatePKCE()
	require.NoError(t, err1, "generatePKCE should not return error")
	require.NotEmpty(t, verifier1, "Verifier should not be empty")
	require.NotEmpty(t, challenge1, "Challenge should not be empty")

	verifier2, challenge2, err2 := generatePKCE()
	require.NoError(t, err2, "generatePKCE should not return error")
	require.NotEmpty(t, verifier2, "Verifier should not be empty")
	require.NotEmpty(t, challenge2, "Challenge should not be empty")

	require.NotEqual(t, verifier1, verifier2, "Each call should generate different verifier")
	require.NotEqual(t, challenge1, challenge2, "Each call should generate different challenge")

	require.GreaterOrEqual(t, len(verifier1), 43, "Verifier should be at least 43 characters")
	require.LessOrEqual(t, len(verifier1), 128, "Verifier should be at most 128 characters")
}

func TestTokenToJSON_Success(t *testing.T) {
	token := &oauth2.Token{
		AccessToken:  "test-access-token",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token",
		Expiry:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	jsonStr, err := TokenToJSON(token)
	require.NoError(t, err, "TokenToJSON should not return error")
	require.NotEmpty(t, jsonStr, "JSON string should not be empty")

	var decoded oauth2.Token
	err = json.Unmarshal([]byte(jsonStr), &decoded)
	require.NoError(t, err, "JSON should be valid")
	require.Equal(t, token.AccessToken, decoded.AccessToken, "Access token should match")
	require.Equal(t, token.TokenType, decoded.TokenType, "Token type should match")
}

func TestTokenFromJSON_Success(t *testing.T) {
	jsonStr := `{
		"access_token": "test-access-token",
		"token_type": "Bearer",
		"refresh_token": "test-refresh-token",
		"expiry": "2024-01-01T00:00:00Z"
	}`

	token, err := TokenFromJSON(jsonStr)
	require.NoError(t, err, "TokenFromJSON should not return error")
	require.NotNil(t, token, "Token should not be nil")
	require.Equal(t, "test-access-token", token.AccessToken, "Access token should match")
	require.Equal(t, "Bearer", token.TokenType, "Token type should match")
	require.Equal(t, "test-refresh-token", token.RefreshToken, "Refresh token should match")
}

func TestTokenFromJSON_InvalidJSON(t *testing.T) {
	jsonStr := `{invalid json`

	token, err := TokenFromJSON(jsonStr)
	require.Error(t, err, "TokenFromJSON should return error for invalid JSON")
	require.Nil(t, token, "Token should be nil on error")
}

func TestTokenToJSON_TokenFromJSON_RoundTrip(t *testing.T) {
	original := &oauth2.Token{
		AccessToken:  "test-access-token",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token",
		Expiry:       time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
	}

	jsonStr, err := TokenToJSON(original)
	require.NoError(t, err, "TokenToJSON should not return error")

	restored, err := TokenFromJSON(jsonStr)
	require.NoError(t, err, "TokenFromJSON should not return error")

	require.Equal(t, original.AccessToken, restored.AccessToken, "Access token should match")
	require.Equal(t, original.TokenType, restored.TokenType, "Token type should match")
	require.Equal(t, original.RefreshToken, restored.RefreshToken, "Refresh token should match")
	require.True(t, original.Expiry.Equal(restored.Expiry), "Expiry should match")
}

func TestExchange_MissingCode(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	token, claims, err := config.Exchange(context.Background(), "", "state", "state")
	require.Error(t, err, "Exchange should return error for missing code")
	require.Nil(t, token, "Token should be nil")
	require.Nil(t, claims, "Claims should be nil")
	require.ErrorIs(t, err, ErrMissingCode, "Error should be ErrMissingCode")
}

func TestExchange_InvalidState(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	token, claims, err := config.Exchange(context.Background(), "code123", "wrong-state", "expected-state")
	require.Error(t, err, "Exchange should return error for invalid state")
	require.Nil(t, token, "Token should be nil")
	require.Nil(t, claims, "Claims should be nil")
	require.ErrorIs(t, err, ErrInvalidState, "Error should be ErrInvalidState")
}

func TestRefreshToken_MissingRefreshToken(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	token := &oauth2.Token{
		AccessToken: "test-token",
		Expiry:      time.Now().Add(-1 * time.Hour),
	}

	newToken, claims, err := config.RefreshToken(context.Background(), token)
	require.Error(t, err, "RefreshToken should return error for missing refresh token")
	require.Nil(t, newToken, "New token should be nil")
	require.Nil(t, claims, "Claims should be nil")
	require.ErrorIs(t, err, ErrTokenRefresh, "Error should be ErrTokenRefresh")
}

func TestParseTokenInfo_EmptyAccessToken(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	claims, err := config.parseTokenInfo("")
	require.Error(t, err, "parseTokenInfo should return error for empty access token")
	require.Nil(t, claims, "Claims should be nil")
	require.Contains(t, err.Error(), "no access token available", "Error message should mention access token")
}

func TestParseTokenInfo_NoKeyFunc(t *testing.T) {
	config, err := NewConfig("test-client-id", "http://localhost/callback", nil)
	require.NoError(t, err)

	claims, err := config.parseTokenInfo("test-token")
	require.Error(t, err, "parseTokenInfo should return error when no keyfunc is available")
	require.Nil(t, claims, "Claims should be nil")
	require.Contains(t, err.Error(), "no keyfunc available", "Error message should mention keyfunc")
}

func TestClient(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	token := &oauth2.Token{
		AccessToken: "test-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}

	client := config.Client(context.Background(), token)
	require.NotNil(t, client, "Client should not be nil")
}

func TestTokenSource(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	token := &oauth2.Token{
		AccessToken:  "test-token",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token",
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	tokenSource := config.TokenSource(context.Background(), token)
	require.NotNil(t, tokenSource, "TokenSource should not be nil")
}

func TestTokenSource_NilToken(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	tokenSource := config.TokenSource(context.Background(), nil)
	require.NotNil(t, tokenSource, "TokenSource should not be nil even with nil token")
}

func TestParseClaims_NilToken(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	claims, err := config.ParseClaims(nil)
	require.Error(t, err, "ParseClaims should return error for nil token")
	require.Nil(t, claims, "Claims should be nil")
	require.Contains(t, err.Error(), "token is nil", "Error message should mention nil token")
}

func TestParseClaims_EmptyAccessToken(t *testing.T) {
	keyFunc := mockKeyFunc()
	config, err := NewConfig("test-client-id", "http://localhost/callback", &keyFunc)
	require.NoError(t, err)

	token := &oauth2.Token{
		AccessToken: "",
		TokenType:   "Bearer",
	}

	claims, err := config.ParseClaims(token)
	require.Error(t, err, "ParseClaims should return error for empty access token")
	require.Nil(t, claims, "Claims should be nil")
	require.Contains(t, err.Error(), "no access token available", "Error message should mention access token")
}

func TestParseClaims_NoKeyFunc(t *testing.T) {
	config, err := NewConfig("test-client-id", "http://localhost/callback", nil)
	require.NoError(t, err)

	token := &oauth2.Token{
		AccessToken: "test-token",
		TokenType:   "Bearer",
	}

	claims, err := config.ParseClaims(token)
	require.Error(t, err, "ParseClaims should return error when no keyfunc is available")
	require.Nil(t, claims, "Claims should be nil")
	require.Contains(t, err.Error(), "no keyfunc available", "Error message should mention keyfunc")
}

func TestStringOrArray_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name        string
		jsonData    string
		wantLen     int
		wantScopes  []string
		expectError bool
	}{
		{
			name:        "array of scopes",
			jsonData:    `["scope1", "scope2", "scope3"]`,
			wantLen:     3,
			wantScopes:  []string{"scope1", "scope2", "scope3"},
			expectError: false,
		},
		{
			name:        "single string scope",
			jsonData:    `"single-scope"`,
			wantLen:     1,
			wantScopes:  []string{"single-scope"},
			expectError: false,
		},
		{
			name:        "empty array",
			jsonData:    `[]`,
			wantLen:     0,
			wantScopes:  []string{},
			expectError: false,
		},
		{
			name:        "empty string",
			jsonData:    `""`,
			wantLen:     1,
			wantScopes:  []string{""},
			expectError: false,
		},
		{
			name:        "invalid JSON object",
			jsonData:    `{invalid}`,
			expectError: true,
		},
		{
			name:        "number instead of string",
			jsonData:    `123`,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var scopes StringOrArray
			err := json.Unmarshal([]byte(tc.jsonData), &scopes)

			if tc.expectError {
				require.Error(t, err, "Should return error")
			} else {
				require.NoError(t, err, "Should not return error")
				require.Len(t, scopes, tc.wantLen, "Scope count should match")
				for i, expectedScope := range tc.wantScopes {
					require.Equal(t, expectedScope, scopes[i], "Scope at index %d should match", i)
				}
			}
		})
	}
}

func TestStringOrArray_MarshalJSON(t *testing.T) {
	testCases := []struct {
		name       string
		scopes     StringOrArray
		wantResult []string
	}{
		{
			name:       "multiple scopes",
			scopes:     StringOrArray{"scope1", "scope2", "scope3"},
			wantResult: []string{"scope1", "scope2", "scope3"},
		},
		{
			name:       "single scope",
			scopes:     StringOrArray{"single-scope"},
			wantResult: []string{"single-scope"},
		},
		{
			name:       "empty array",
			scopes:     StringOrArray{},
			wantResult: []string{},
		},
		{
			name:       "nil array",
			scopes:     nil,
			wantResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.scopes)
			require.NoError(t, err, "MarshalJSON should not return error")

			var result []string
			err = json.Unmarshal(jsonData, &result)
			require.NoError(t, err, "Result should be valid JSON array")

			if tc.wantResult == nil {
				require.Nil(t, result, "Result should be nil")
			} else {
				require.Equal(t, tc.wantResult, result, "Result should match expected")
			}
		})
	}
}

func TestESIJWTClaims_ScopesUnmarshaling(t *testing.T) {
	testCases := []struct {
		name        string
		jsonData    string
		wantScopes  []string
		expectError bool
	}{
		{
			name: "array of scopes",
			jsonData: `{
				"sub": "CHARACTER:EVE:123456789",
				"name": "Test Character",
				"scp": ["scope1", "scope2", "scope3"],
				"owner": "owner123",
				"exp": 1234567890
			}`,
			wantScopes:  []string{"scope1", "scope2", "scope3"},
			expectError: false,
		},
		{
			name: "single string scope",
			jsonData: `{
				"sub": "CHARACTER:EVE:123456789",
				"name": "Test Character",
				"scp": "single-scope",
				"owner": "owner123",
				"exp": 1234567890
			}`,
			wantScopes:  []string{"single-scope"},
			expectError: false,
		},
		{
			name: "empty array of scopes",
			jsonData: `{
				"sub": "CHARACTER:EVE:123456789",
				"name": "Test Character",
				"scp": [],
				"owner": "owner123",
				"exp": 1234567890
			}`,
			wantScopes:  []string{},
			expectError: false,
		},
		{
			name: "missing scp field",
			jsonData: `{
				"sub": "CHARACTER:EVE:123456789",
				"name": "Test Character",
				"owner": "owner123",
				"exp": 1234567890
			}`,
			wantScopes:  nil,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var claims ESIJWTClaims
			err := json.Unmarshal([]byte(tc.jsonData), &claims)

			if tc.expectError {
				require.Error(t, err, "Should return error")
			} else {
				require.NoError(t, err, "Should not return error")
				require.Equal(t, "Test Character", claims.Name, "Name should match")
				require.Equal(t, "CHARACTER:EVE:123456789", claims.Subject, "Subject should match")

				if tc.wantScopes == nil {
					require.Nil(t, claims.Scopes, "Scopes should be nil")
				} else {
					require.Equal(t, tc.wantScopes, []string(claims.Scopes), "Scopes should match")
				}
			}
		})
	}
}

func TestESIJWTClaims_MarshalUnmarshal_RoundTrip(t *testing.T) {
	testCases := []struct {
		name   string
		scopes StringOrArray
	}{
		{
			name:   "multiple scopes",
			scopes: StringOrArray{"scope1", "scope2"},
		},
		{
			name:   "single scope",
			scopes: StringOrArray{"single-scope"},
		},
		{
			name:   "empty scopes",
			scopes: StringOrArray{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			original := ESIJWTClaims{
				RegisteredClaims: jwt.RegisteredClaims{
					Subject:   "CHARACTER:EVE:123456789",
					ExpiresAt: jwt.NewNumericDate(time.Unix(1234567890, 0)),
				},
				Scopes: tc.scopes,
				Name:   "Test Character",
				Owner:  "owner123",
			}

			jsonData, err := json.Marshal(original)
			require.NoError(t, err, "Should marshal claims")

			var restored ESIJWTClaims
			err = json.Unmarshal(jsonData, &restored)
			require.NoError(t, err, "Should unmarshal claims")

			require.Equal(t, original.Subject, restored.Subject, "Subject should match")
			require.Equal(t, original.Name, restored.Name, "Name should match")
			require.Equal(t, original.Owner, restored.Owner, "Owner should match")
			require.Equal(t, []string(original.Scopes), []string(restored.Scopes), "Scopes should match")
		})
	}
}
