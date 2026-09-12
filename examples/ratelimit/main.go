package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/fnt-eve/goesi-openapi"
	"github.com/fnt-eve/goesi-openapi/ratelimit"
)

// This example demonstrates how to use the rate limiter transport
// with the GoESI client to respect ESI's rate limits.
func main() {
	userAgent := "MyEVEApp/1.0 (contact@example.com)"

	// Step 1: Create the rate limiting transport
	// This wraps the default http.Transport (or any other transport)
	rlTransport := ratelimit.NewTransport(http.DefaultTransport)

	// Step 2: Create an http.Client using the rate limiting transport
	httpClient := &http.Client{
		Transport: rlTransport,
	}

	// Step 3: Create the ESI client using the rate limited http client
	client := goesi.NewESIClientWithOptions(httpClient, goesi.ClientOptions{
		UserAgent: userAgent,
	})

	// Step 4: Make requests
	// The rate limiter will automatically:
	// - Match the request to the correct ESI rate limit bucket
	// - Wait if the bucket is empty
	// - Update the bucket state based on response headers

	ctx := context.Background()
	fmt.Println("Making requests to /status/...")

	// Make a few requests to demonstrate rate limiting (though limits are high for /status)
	for i := range 3 {
		status, _, err := client.StatusAPI.GetStatus(ctx).Execute()
		if err != nil {
			log.Printf("Request %d failed: %v", i+1, err)
			continue
		}

		fmt.Printf("Request %d: Server %s, Players %d\n", i+1, status.ServerVersion, status.Players)
	}

	fmt.Println("\nRate limiter example completed successfully!")
}
