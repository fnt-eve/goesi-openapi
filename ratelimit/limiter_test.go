package ratelimit

import (
	"context"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tests := []struct {
		name        string
		limit       int
		window      time.Duration
		initialWait int
		secondWait  int
		timeout     time.Duration
		expectError bool
		minWaitTime time.Duration
	}{
		{
			name:        "Basic Wait",
			limit:       10,
			window:      time.Second,
			initialWait: 10,
			secondWait:  0,
			timeout:     time.Second,
			expectError: false,
		},
		{
			name:        "Wait with Delay",
			limit:       10,
			window:      time.Second,
			initialWait: 10,
			secondWait:  1, // Requires waiting
			timeout:     time.Second,
			expectError: false,
			minWaitTime: 90 * time.Millisecond,
		},
		{
			name:        "Context Cancel",
			limit:       1,
			window:      time.Minute,
			initialWait: 1,
			secondWait:  1, // Should wait 1 minute, but timeout is short
			timeout:     10 * time.Millisecond,
			expectError: true,
		},
		{
			name:        "Zero Cost",
			limit:       10,
			window:      time.Second,
			initialWait: 0,
			secondWait:  0,
			timeout:     time.Second,
			expectError: false,
		},
		{
			name:        "Exceed Capacity",
			limit:       5,
			window:      time.Second,
			initialWait: 0,
			secondWait:  10, // Cost > Capacity. Should wait until accumulated.
			// Rate = 5/s. Need 10 tokens.
			// Start with 5. Remaining = -5. Deficit = 5.
			// Wait time = 5 / 5 = 1s.
			timeout:     2 * time.Second,
			expectError: false,
			minWaitTime: 950 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bucket := NewTokenBucket(tt.limit, tt.window)

			ctx := context.Background()
			if tt.timeout > 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tt.timeout)
				defer cancel()
			}

			// Initial consume
			if err := bucket.Wait(ctx, tt.initialWait); err != nil {
				if !tt.expectError {
					t.Errorf("Initial Wait() error = %v", err)
				}
				return
			}

			if tt.secondWait > 0 {
				start := time.Now()
				err := bucket.Wait(ctx, tt.secondWait)
				elapsed := time.Since(start)

				if (err != nil) != tt.expectError {
					t.Errorf("Wait() error = %v, expectError %v", err, tt.expectError)
				}

				if !tt.expectError && elapsed < tt.minWaitTime {
					t.Errorf("Wait() took %v, want > %v", elapsed, tt.minWaitTime)
				}
			}
		})
	}
}
