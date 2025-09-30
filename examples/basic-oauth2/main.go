package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fnt-eve/goesi-openapi"
)

// This example demonstrates how to set up ESI OAuth2 authentication
// and make authenticated API calls to ESI endpoints.
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

	fmt.Printf("Successfully authenticated!\nRaw JWT Token: %s\n", token.AccessToken)

	// Step 4: Create authenticated ESI client
	esiClient := goesi.NewAuthenticatedESIClient(ctx, config, token, userAgent)

	// Step 5: Display character information from JWT claims
	fmt.Println("\nCharacter Information from JWT Claims:")

	// Get character ID from the JWT claims
	characterID, err := claims.CharacterID()
	if err != nil {
		log.Fatalf("Failed to get character ID from claims: %v", err)
	}

	fmt.Printf("Character ID: %d\n", characterID)
	fmt.Printf("Character Name: %s\n", claims.CharacterName())
	fmt.Printf("Token Scopes: %v\n", claims.TokenScopes())

	// Step 6: Make authenticated API calls
	fmt.Println("\nMaking authenticated API calls...")

	// Make API call using the character ID from the claims
	locationInfo, _, err := esiClient.LocationAPI.GetCharactersCharacterIdLocation(ctx, int64(characterID)).Execute()

	if err != nil {
		log.Printf("Failed to get character location: %v", err)
	} else {
		fmt.Printf("Location from API: %d\n", locationInfo.GetSolarSystemId())
	}

	fmt.Println("Authentication example completed successfully!")
}
