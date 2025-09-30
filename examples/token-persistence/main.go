package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fnt-eve/goesi-openapi"
)

// This example demonstrates how to persist OAuth2 tokens to storage
// and restore them later using TokenSource for automatic token refresh.
func main() {
	// Step 1: Set up OAuth2 configuration
	clientID := os.Getenv("ESI_CLIENT_ID")
	if clientID == "" {
		log.Fatal("ESI_CLIENT_ID environment variable is required")
	}

	redirectURL := "http://localhost:8080/esi/callback"
	userAgent := "MyEVEApp/1.0 (contact@example.com)"

	ctx := context.Background()
	keyFunc, err := goesi.ESIDefaultKeyfunc(ctx)
	if err != nil {
		log.Fatalf("Failed to create ESI keyfunc: %v", err)
	}

	config, err := goesi.NewConfig(clientID, redirectURL, &keyFunc)
	if err != nil {
		log.Fatalf("Failed to create OAuth2 config: %v", err)
	}

	// Step 2: Get authorization with requested scopes
	state := "random-state-string"
	scopes := []string{
		goesi.ScopeLocationReadLocationV1,
		goesi.ScopeSkillsReadSkillsV1,
	}
	authURL := config.AuthURL(state, scopes)

	fmt.Printf("Please visit this URL to authorize the application:\n%s\n", authURL)
	fmt.Print("Enter the authorization code from the callback: ")

	var code string
	fmt.Scanln(&code)

	// Step 3: Exchange code for token
	token, claims, err := config.Exchange(ctx, code, state, state)
	if err != nil {
		log.Fatalf("Failed to exchange code for token: %v", err)
	}

	characterID, err := claims.CharacterID()
	if err != nil {
		log.Fatalf("Failed to get character ID: %v", err)
	}

	fmt.Printf("\nAuthenticated as: %s (ID: %d)\n", claims.Name, characterID)

	// Step 4: Serialize token for storage
	tokenJSON, err := goesi.TokenToJSON(token)
	if err != nil {
		log.Fatalf("Failed to serialize token: %v", err)
	}

	fmt.Println("\nToken serialized to JSON (ready to store in database/file)")
	fmt.Printf("JSON length: %d bytes\n", len(tokenJSON))

	var storedTokenJSON string = tokenJSON // Simulating storage

	// Step 5: Deserialize token from storage
	restoredToken, err := goesi.TokenFromJSON(storedTokenJSON)
	if err != nil {
		log.Fatalf("Failed to deserialize token: %v", err)
	}

	fmt.Println("Token restored from storage")

	// Step 6: Create TokenSource for automatic token refresh
	// IMPORTANT: Create TokenSource before parsing claims, as it will
	// automatically refresh expired tokens
	tokenSource := config.TokenSource(ctx, restoredToken)
	fmt.Println("Created TokenSource for automatic token refresh")

	// Step 7: Get current token from TokenSource
	// This will refresh the token if it's expired
	currentToken, err := tokenSource.Token()
	if err != nil {
		log.Fatalf("Failed to get current token: %v", err)
	}

	if currentToken.AccessToken != restoredToken.AccessToken {
		fmt.Println("Token was automatically refreshed by TokenSource!")

		// Be sure to serialize and store the token as the refresh token may be rotated
		//refreshedJSON, _ := goesi.TokenToJSON(currentToken) //nolint:ineffassign
		//storedTokenJSON = refreshedJSON                     //nolint:ineffassign
		fmt.Println("Refreshed token saved back to storage")
	} else {
		fmt.Println("Token is still valid, no refresh needed")
	}

	// Step 8: Now parse claims from the current (possibly refreshed) token
	// ParseClaims validates the token, so we need a valid token first
	restoredClaims, err := config.ParseClaims(currentToken)
	if err != nil {
		log.Fatalf("Failed to parse claims: %v", err)
	}

	restoredCharID, _ := restoredClaims.CharacterID()
	fmt.Printf("\nRestored session for: %s (ID: %d)\n", restoredClaims.Name, restoredCharID)

	// Step 9: Create ESI client with current token
	esiClient := goesi.NewAuthenticatedESIClient(ctx, config, currentToken, userAgent)

	// Step 10: Make API calls
	fmt.Println("\nMaking authenticated API calls...")

	locationInfo, _, err := esiClient.LocationAPI.GetCharactersCharacterIdLocation(ctx, int64(restoredCharID)).Execute()
	if err != nil {
		log.Printf("Failed to get character location: %v", err)
	} else {
		fmt.Printf("Current location: System ID %d\n", locationInfo.GetSolarSystemId())
	}

	fmt.Println("\nToken persistence example completed successfully!")
}
