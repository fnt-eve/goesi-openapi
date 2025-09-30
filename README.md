# GoESI OpenAPI Client

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.24-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A Go client library for the [EVE Online ESI API](https://esi.evetech.net/), generated from the official OpenAPI specification.

## Installation

```bash
go get github.com/fnt-eve/goesi-openapi
```

## Quick Start

### Public API (No Authentication)

```go
package main

import (
    "context"
    "log"

    "github.com/fnt-eve/goesi-openapi"
)

func main() {
    ctx := context.Background()
    
    // Create client for public endpoints
    client := goesi.NewPublicESIClient("MyApp/1.0 (contact@example.com)")
    
    // Get system information
    system, _, err := client.UniverseAPI.GetUniverseSystemsSystemId(ctx, 30000142).Execute()
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("System: %s", system.GetName())
}
```

### Authenticated API

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/fnt-eve/goesi-openapi"
)

func main() {
    ctx := context.Background()
    
    // Set up OAuth2
    clientID := os.Getenv("ESI_CLIENT_ID")
    redirectURL := "http://localhost:8080/callback"
    scopes := []string{goesi.ScopeLocationReadLocationV1}
    
    // Create JWT key function for token validation
    keyFunc, err := goesi.ESIDefaultKeyfunc(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    // Create OAuth2 config
    config, err := goesi.NewConfig(clientID, redirectURL, scopes, &keyFunc)
    if err != nil {
        log.Fatal(err)
    }
    
    // Get authorization URL
    state := "random-state-string"
    authURL := config.AuthURL(state)
    log.Printf("Visit: %s", authURL)
    
    // Exchange code for token (after user authorization)
    // var code string // Get this from the callback
    // token, claims, err := config.Exchange(ctx, code, state, state)
    // if err != nil {
    //     log.Fatal(err)
    // }
    
    // Create authenticated client
    // client := goesi.NewAuthenticatedESIClient(ctx, config, token, "MyApp/1.0 (contact@example.com)")
}
```

## OAuth2 Authentication

ESI uses OAuth2 with PKCE. The authentication flow:

1. **Create config** with client ID, redirect URL, and scopes
2. **Generate auth URL** and redirect user to authorize
3. **Exchange code** for access token
4. **Use token** to make authenticated API calls
5. **Refresh token** when it expires

### Available Scopes

The library includes constants for all ESI scopes:

```go
scopes := []string{
    goesi.ScopeLocationReadLocationV1,
    goesi.ScopeAssetsReadAssetsV1, 
    goesi.ScopeSkillsReadSkillsV1,
    goesi.ScopeCorporationsReadBlueprintsV1,
    // ... many more
}
```

### JWT Token Information

Access tokens are JWTs containing character information. The `Exchange` method returns both the token and parsed claims:

```go
// Exchange authorization code for token and claims
token, claims, err := config.Exchange(ctx, code, state, state)
if err != nil {
    log.Fatal(err)
}

// Extract character information from claims
characterID, err := claims.CharacterID()    // int32
characterName := claims.CharacterName()     // string  
scopes := claims.TokenScopes()              // []string
```

## API Usage

The generated client provides access to all ESI endpoints through API groups:

```go
client := goesi.NewPublicESIClient("MyApp/1.0 (contact@example.com)")

// Character information
client.CharacterAPI.GetCharactersCharacterId(ctx, characterID)

// Market data  
client.MarketAPI.GetMarketsRegionIdOrders(ctx, regionID)

// Universe data
client.UniverseAPI.GetUniverseSystemsSystemId(ctx, systemID)

// Corporation data
client.CorporationAPI.GetCorporationsCorporationId(ctx, corporationID)

// Alliance data
client.AllianceAPI.GetAlliancesAllianceId(ctx, allianceID)
```

## Error Handling

```go
data, response, err := client.CharacterAPI.GetCharactersCharacterId(ctx, characterID).Execute()
if err != nil {
    if apiErr, ok := err.(*esi.GenericOpenAPIError); ok {
        log.Printf("API Error: %s", apiErr.Error())
        log.Printf("Response: %s", apiErr.Body())
    } else {
        log.Printf("Other error: %v", err)
    }
    return
}
```

## Token Management

```go
// Check if token needs refresh
if goesi.IsExpired(token) {
    newToken, newClaims, err := config.RefreshToken(ctx, token)
    if err != nil {
        log.Fatal("Token refresh failed:", err)
    }
    token = newToken
    claims = newClaims
}

// Store/load tokens (stores only the oauth2.Token, not claims)
tokenJSON, _ := goesi.TokenToJSON(token)
// Store tokenJSON in database/file

// Later: restore token and parse claims
storedToken, _ := goesi.TokenFromJSON(tokenJSON)
claims, err := config.ParseClaims(storedToken)
if err != nil {
    log.Fatal("Failed to parse claims:", err)
}
```

## Rate Limits

ESI rate limits:
- **20 requests/second** for authenticated requests  
- **10 requests/second** for public requests

The client automatically includes required headers like `X-Compatibility-Date`.

## Examples

See [`examples/`](examples/) directory:
- [`basic-oauth2/`](examples/basic-oauth2/main.go) - Complete OAuth2 flow with authenticated client
- [`context-auth/`](examples/context-auth/main.go) - Context-based authentication pattern

## Development

### Code Generation

The client is generated from the ESI OpenAPI specification:

```bash
make generate
```

This downloads the latest ESI spec, generates the client code, and runs post-processing scripts.

### Building

```bash
go build ./...
go test ./...
go run examples/oauth2_example.go
```

## Library Configuration

### User Agent 

ESI requires all requests to include a User-Agent header with contact information:

```go
userAgent := "MyEVEApp/1.0 (contact@example.com)"
client := goesi.NewPublicESIClient(userAgent)
```

Format: `AppName/Version (contact-email)`

## Running Examples

The examples require an ESI application registration:

### 1. Register Your Application

Visit [EVE Developers](https://developers.eveonline.com/) to create an application and get a Client ID.

### 2. Set Environment Variables

```bash
export ESI_CLIENT_ID="your-client-id-from-developers-portal"
```

### 3. Run Examples

```bash
go run examples/basic-oauth2/main.go
go run examples/context-auth/main.go
```

## Dependencies

- `golang.org/x/oauth2` - OAuth2 client
- `github.com/golang-jwt/jwt/v5` - JWT token parsing
- `github.com/MicahParks/keyfunc/v3` - JWT key validation

## Requirements

- Go 1.24+
- Valid ESI application registration

## Resources

- [ESI Documentation](https://esi.evetech.net/ui/)
- [ESI OpenAPI Spec](https://esi.evetech.net/meta/openapi-3.0.json)
- [Developer Resources](https://developers.eveonline.com/)

## License

MIT License - see [LICENSE](LICENSE) file.

## Disclaimer

Not affiliated with CCP Games. EVE Online is a trademark of CCP hf.
