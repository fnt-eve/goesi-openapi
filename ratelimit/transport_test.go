package ratelimit

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"
)

// mockRoundTripper implements http.RoundTripper
type mockRoundTripper struct {
	resp *http.Response
	err  error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	default:
	}
	return m.resp, m.err
}

func TestTransport_RoundTrip(t *testing.T) {
	// Setup dummy route in mapping for testing
	// Example: /status/ (GET, group status, 600, 15m)
	// Example: /alliances/{alliance_id}/contacts/ (GET, alliance-social, 300, 15m)

	tests := []struct {
		name            string
		reqMethod       string
		reqURL          string
		reqHeader       http.Header
		respHeader      map[string]string
		expectedSubject string
		expectedGroup   string
		expectWait      bool
		expectSync      int // -1 if no sync expected
		expectError     bool
		mockErr         error
		cancelContext   bool
		setup           func(t *Transport)
	}{
		{
			name:       "No Match PassThrough",
			reqMethod:  "GET",
			reqURL:     "https://esi.evetech.net/v1/unknown",
			reqHeader:  make(http.Header),
			expectSync: -1,
		},
		{
			name:            "Public Rate Limit",
			reqMethod:       "GET",
			reqURL:          "https://esi.evetech.net/status/",
			reqHeader:       make(http.Header),
			expectedSubject: "public",
			expectedGroup:   "status",
			expectSync:      50,
			respHeader:      map[string]string{"X-RateLimit-Remaining": "50"},
		},
		{
			name:      "Public Rate Limit with JWT",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/status/",
			reqHeader: http.Header{
				"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaGFyYWN0ZXI6RVZFOjEyMyIsImF6cCI6ImNsaWVudCJ9.signature"},
			},
			expectedSubject: "public:client",
			expectedGroup:   "status",
			expectSync:      -1,
		},
		{
			name:      "Auth Rate Limit (Standard JWT)",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/characters/123/location/",
			reqHeader: http.Header{
				// {"sub":"123", "azp":"client1"}
				"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjMiLCJhenAiOiJjbGllbnQxIn0.signature"},
			},
			expectedSubject: "client1:123",
			expectedGroup:   "char-location",
			expectSync:      -1,
		},
		{
			name:      "Auth Rate Limit (ESI Format)",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/characters/123456/location/",
			reqHeader: http.Header{
				// {"sub":"CHARACTER:EVE:123456", "azp":"myclient"}
				"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaGFyYWN0ZXI6RVZFOjEyMzQ1NiIsImF6cCI6Im15Y2xpZW50In0.signature"},
			},
			expectedSubject: "myclient:Character:EVE:123456",
			expectedGroup:   "char-location",
			expectSync:      -1,
		},
		{
			name:            "Regex Test - Character Fleet",
			reqMethod:       "GET",
			reqURL:          "https://esi.evetech.net/characters/123/fleet/",
			reqHeader:       make(http.Header),
			expectedSubject: "public",
			expectedGroup:   "fleet",
			expectSync:      -1,
		},
		{
			name:            "Regex Test - Corp Killmails",
			reqMethod:       "GET",
			reqURL:          "https://esi.evetech.net/corporations/456/killmails/recent/",
			reqHeader:       make(http.Header),
			expectedSubject: "public",
			expectedGroup:   "corp-killmail",
			expectSync:      -1,
		},
		{
			name:          "Wait Error",
			reqMethod:     "GET",
			reqURL:        "https://esi.evetech.net/status/",
			reqHeader:     make(http.Header),
			cancelContext: true,
			expectWait:    false, // Should fail during wait
			expectError:   true,
			setup: func(tr *Transport) {
				// Exhaust bucket for status:public to force wait
				key := BucketKey("status:public")
				// status limit 600, window 15m.
				bucket := tr.Store.GetBucket(key, 600, 15*time.Minute)
				bucket.Sync(0) // Set remaining to 0
			},
		},
		{
			name:        "RoundTrip Error",
			reqMethod:   "GET",
			reqURL:      "https://esi.evetech.net/status/",
			reqHeader:   make(http.Header),
			mockErr:     errors.New("network error"),
			expectError: true,
		},
		{
			name:      "Auth Fallback - Sub Only",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/characters/123/location/",
			reqHeader: http.Header{
				// {"sub":"123"}
				"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjMifQ.sig"},
			},
			expectedSubject: "123",
			expectedGroup:   "char-location",
			expectSync:      -1,
		},
		{
			name:      "Auth Malformed - No Bearer",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/characters/123/location/",
			reqHeader: http.Header{
				"Authorization": []string{"Basic 123"},
			},
			expectedSubject: "public",
			expectedGroup:   "char-location",
			expectSync:      -1,
		},
		{
			name:      "Auth Malformed - Bearer Empty",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/characters/123/location/",
			reqHeader: http.Header{
				"Authorization": []string{"Bearer"},
			},
			expectedSubject: "public",
			expectedGroup:   "char-location",
			expectSync:      -1,
		},
		{
			name:      "Auth Invalid JWT",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/characters/123/location/",
			reqHeader: http.Header{
				"Authorization": []string{"Bearer invalid.token"},
			},
			expectedSubject: "public",
			expectedGroup:   "char-location",
			expectSync:      -1,
		},
		{
			name:      "Auth Empty Claims",
			reqMethod: "GET",
			reqURL:    "https://esi.evetech.net/characters/123/location/",
			reqHeader: http.Header{
				// {}
				"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.sig"},
			},
			expectedSubject: "public",
			expectedGroup:   "char-location",
			expectSync:      -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := make(http.Header)
			if tt.respHeader != nil {
				for k, v := range tt.respHeader {
					header.Set(k, v)
				}
			}

			mockResp := &http.Response{
				StatusCode: 200,
				Header:     header,
			}

			base := &mockRoundTripper{resp: mockResp, err: tt.mockErr}
			transport := NewTransport(base)

			if tt.setup != nil {
				tt.setup(transport)
			}

			u, _ := url.Parse(tt.reqURL)
			req := &http.Request{
				Method: tt.reqMethod,
				URL:    u,
				Header: tt.reqHeader,
			}

			if tt.cancelContext {
				ctx, cancel := context.WithCancel(req.Context())
				cancel() // cancel immediately
				req = req.WithContext(ctx)
			}

			_, err := transport.RoundTrip(req)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("RoundTrip failed: %v", err)
			}

			rc := transport.matchRoute(req)
			if tt.expectedGroup != "" {
				if rc == nil {
					t.Errorf("Expected group %s, got no match", tt.expectedGroup)
				} else if rc.Group != tt.expectedGroup {
					t.Errorf("Expected group %s, got %s", tt.expectedGroup, rc.Group)
				}
			}

			if tt.expectSync != -1 {
				if rc == nil {
					t.Errorf("Expected match for sync, got nil")
				} else {
					key := transport.getBucketKey(req, rc)
					bucket := transport.Store.GetBucket(key, rc.Limit, rc.Window)

					bucket.mu.Lock()
					rem := bucket.remaining
					bucket.mu.Unlock()

					if int(rem) != tt.expectSync {
						t.Errorf("Expected remaining %d, got %f", tt.expectSync, rem)
					}
				}
			}

			if tt.expectedSubject != "" && rc != nil {
				key := transport.getBucketKey(req, rc)
				expectedKey := BucketKey(tt.expectedGroup + ":" + tt.expectedSubject)
				if key != expectedKey {
					t.Errorf("getBucketKey() = %v, want %v", key, expectedKey)
				}
			}
		})
	}
}
