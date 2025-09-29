# GoESI OpenAPI Client

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A comprehensive Go client library for the [EVE Online ESI (EVE Swagger Interface) API](https://esi.evetech.net/), automatically generated from the official OpenAPI 3.0 specification.

## Overview

This library provides a type-safe, feature-complete Go client for accessing EVE Online game data through the ESI API. It includes:

- **Auto-generated client** from the official ESI OpenAPI specification
- **OAuth2 authentication** with PKCE support for secure API access
- **Complete API coverage** for all ESI endpoints
- **Type safety** with strongly-typed Go structs for all requests and responses
- **Built-in rate limiting** compliance with ESI requirements
- **Automatic header management** including required ESI headers

## Features

- **Complete ESI API Coverage**: Auto-generated from the official OpenAPI spec
- **OAuth2 Authentication**: Full OAuth2 flow with PKCE for security
- **Type Safety**: Strongly-typed structs for all API calls
- **Rate Limit Compliance**: Respects ESI's rate limiting requirements
- **Caching Headers**: Proper handling of ESI cache-control headers
- **Error Handling**: Comprehensive error types and handling
- **Go Modules**: Full Go modules support for easy dependency management

## Installation

```bash
go get github.com/fnt-eve/goesi-openapi
```

## Quick Start

### 1. Basic Setup

```go
package main

import (
    "context"
    "log"

    "github.com/fnt-eve/goesi-openapi"
)

func main() {
    // Create OAuth2 configuration
    config, err := goesi.NewConfig(
        "your-client-id",
        "http://localhost:8080/callback",
        []string{goesi.ScopeLocationReadLocationV1},
    )
    if err != nil {
        log.Fatal(err)
    }

    // Generate authorization URL
    authURL := config.AuthURL("random-state-string")
    log.Printf("Visit: %s", authURL)
    
    // ... handle OAuth2 flow (see examples/ for complete flow)
}
```

### 2. Making API Calls

```go
// After obtaining a token through OAuth2 flow, use the helper to create ESI client
userAgent := "MyApp/1.0 (contact@example.com)"
client := goesi.NewAuthenticatedESIClient(ctx, config, token, userAgent)

// Make API calls
characterInfo, _, err := client.CharacterAPI.GetCharactersCharacterId(ctx, characterID).Execute()
if err != nil {
    log.Fatal(err)
}

log.Printf("Character: %s", characterInfo.Name)
```

### 3. Public API Calls (No Authentication)

```go
// For public endpoints that don't require authentication
userAgent := "MyApp/1.0 (contact@example.com)"
client := goesi.NewPublicESIClient(userAgent)

// Make public API calls
systemInfo, _, err := client.UniverseAPI.GetUniverseSystemsSystemId(ctx, systemID).Execute()
if err != nil {
    log.Fatal(err)
}

log.Printf("System: %s", systemInfo.Name)
```

## Authentication

ESI uses OAuth2 with PKCE for authentication. This library provides a complete OAuth2 implementation:

### OAuth2 Scopes

The library provides idiomatic Go constants for all available ESI scopes:

```go
scopes := []string{
    goesi.ScopeLocationReadLocationV1,
    goesi.ScopeAssetsReadAssetsV1,
    goesi.ScopeSkillsReadSkillsV1,
    goesi.ScopeCorporationsReadBlueprintsV1,
    // ... many more available, all following CamelCase convention
}
```

### JWT Token Parsing

The library automatically parses JWT tokens returned by ESI and provides convenient methods:

```go
// After successful OAuth2 exchange, the token contains parsed JWT claims
token, err := config.Exchange(ctx, code, state, expectedState)

// Extract information from the JWT token
characterID, err := token.CharacterID()       // int32
characterName := token.CharacterName()        // string
scopes := token.TokenScopes()                // []string
```

### Complete OAuth2 Flow

```go
// 1. Create configuration
config, err := goesi.NewConfig(clientID, redirectURL, scopes)

// 2. Get authorization URL
authURL := config.AuthURL(state)

// 3. User authorizes and you receive the code
token, err := config.Exchange(ctx, code, state, expectedState)

// 4. Use token to make authenticated requests or create ESI client
userAgent := "MyApp/1.0 (contact@example.com)"
client := goesi.NewAuthenticatedESIClient(ctx, config, token, userAgent)

// 5. Refresh token when needed
if token.IsExpired() {
    token, err = config.RefreshToken(ctx, token)
}
```

## API Structure

The generated client organizes ESI endpoints into logical groups:

- **CharacterAPI**: Character information, skills, assets, etc.
- **CorporationAPI**: Corporation data and management
- **MarketAPI**: Market orders and history
- **UniverseAPI**: Static universe data (systems, types, etc.)
- **AllianceAPI**: Alliance information
- **And many more...**

## Examples

See the [`examples/`](examples/) directory for complete working examples:

- [`oauth2_example.go`](examples/oauth2_example.go): Complete OAuth2 authentication flow

## Configuration

### Environment Variables

Set up your ESI application credentials:

```bash
export ESI_CLIENT_ID="your-client-id-here"
```

### User Agent

ESI requires a proper User-Agent header with contact information:

```go
userAgent := "MyEVEApp/1.0 (contact@example.com)"
```

## Rate Limits

ESI has rate limits that this library respects:
- **20 requests/second** for authenticated requests
- **10 requests/second** for unauthenticated requests

The library automatically includes required headers and follows ESI best practices.

## Error Handling

The library provides comprehensive error handling:

```go
characterInfo, response, err := client.CharacterAPI.GetCharactersCharacterId(ctx, characterID).Execute()
if err != nil {
    // Handle different types of errors
    switch e := err.(type) {
    case *esi.GenericOpenAPIError:
        log.Printf("API Error: %s", e.Error())
        log.Printf("Response Body: %s", e.Body())
    default:
        log.Printf("Other error: %v", err)
    }
    return
}
```

## Caching

ESI responses include cache headers. Respect these to minimize API load:

```go
import "time"

// Check cache expiration
expires := esi.CacheExpires(response)
if time.Now().Before(expires) {
    // Use cached data
}
```

## Development

### Code Generation

This library is generated from the ESI OpenAPI specification. The `generate` command in the `Makefile` will regenerate the client from the latest ESI spec, fix the generated code, and run `go generate` to update scopes and compatibility dates.

```bash
# Regenerate client and update generated files
make generate
```

### Building

```bash
# Build the library
go build ./...

# Run tests
go test ./...

# Run examples
go run examples/oauth2_example.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## ESI Resources

- [ESI Documentation](https://esi.evetech.net/ui/)
- [ESI OpenAPI Specification](https://esi.evetech.net/meta/openapi-3.0.json)
- [Third Party Developer Resources](https://developers.eveonline.com/)
- [ESI Community](https://github.com/esi/esi-issues)

## Requirements

- Go 1.25 or later
- Valid ESI application registration
- Network access to ESI endpoints

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This library is not affiliated with CCP Games. EVE Online is a trademark of CCP hf.

## Support

- Check the [examples/](examples/) directory for usage examples
- Report bugs via GitHub Issues  
- Discuss on EVE Online third-party developer communities
- Contact: See User-Agent header for contact information

---

**Built for the EVE Online development community**
