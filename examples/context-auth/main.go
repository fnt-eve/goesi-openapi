package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fnt-eve/goesi-openapi"
	"golang.org/x/oauth2"
)

// This example demonstrates context-based ESI authentication where
// authentication is passed per-request via context to a public ESI client.
func main() {
	// Step 1: Set up OAuth2 configuration
	clientID := os.Getenv("ESI_CLIENT_ID")
	if clientID == "" {
		log.Fatal("ESI_CLIENT_ID environment variable is required")
	}

	redirectURL := "http://localhost:8080/esi/callback"
	userAgent := "MyEVEApp/1.0 (contact@example.com)" // Replace with your contact info

	// Request scopes for character information
	scopes := []string{
		goesi.ScopeLocationReadLocationV1,
		goesi.ScopeAssetsReadAssetsV1,
		goesi.ScopeSkillsReadSkillsV1,
	}

	// Create keyfunc for JWT token validation
	ctx := context.Background()
	keyFunc, err := goesi.ESIDefaultKeyfunc(ctx)
	if err != nil {
		log.Fatalf("Failed to create ESI keyfunc: %v", err)
	}

	config, err := goesi.NewConfig(clientID, redirectURL, scopes, &keyFunc)
	if err != nil {
		log.Fatalf("Failed to create OAuth2 config: %v", err)
	}

	// Step 2: Generate authorization URL
	state := "random-state-string" // In production, use a secure random string
	authURL := config.AuthURL(state)

	fmt.Printf("Please visit this URL to authorize the application:\n%s\n", authURL)
	fmt.Print("Enter the authorization code from the callback: ")

	var code string
	fmt.Scanln(&code)

	// Step 3: Exchange authorization code for access token
	token, claims, err := config.Exchange(ctx, code, state, state)
	if err != nil {
		log.Fatalf("Failed to exchange code for token: %v", err)
	}

	fmt.Printf("Successfully authenticated!\n")
	fmt.Printf("Character Name: %s\n", claims.CharacterName())

	// Step 4: Create a public ESI client
	esiClient := goesi.NewPublicESIClient(userAgent)

	// Step 5: Create context with OAuth2 token source for authentication
	tokenSource := oauth2.StaticTokenSource(token)
	authCtx := context.WithValue(ctx, goesi.ContextOAuth2, tokenSource)

	// Step 6: Display character information from JWT claims
	characterID, err := claims.CharacterID()
	if err != nil {
		log.Fatalf("Failed to get character ID from claims: %v", err)
	}

	fmt.Printf("Character ID: %d\n", characterID)
	fmt.Printf("Token Scopes: %v\n", claims.TokenScopes())

	// Step 7: Make authenticated API calls using the authenticated context
	fmt.Println("\nMaking authenticated API calls...")

	locationInfo, _, err := esiClient.LocationAPI.GetCharactersCharacterIdLocation(authCtx, int64(characterID)).Execute()
	if err != nil {
		log.Printf("Failed to get character location: %v", err)
	} else {
		fmt.Printf("Location from API: %d\n", locationInfo.GetSolarSystemId())
	}

	fmt.Println("Context-based authentication example completed successfully!")
}
