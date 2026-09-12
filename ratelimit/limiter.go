package ratelimit

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TokenBucket implements a thread-safe token bucket algorithm
// Inspired by golang.org/x/time/rate
type TokenBucket struct {
	mu          sync.Mutex
	capacity    float64
	remaining   float64
	rate        float64 // tokens per second
	lastUpdated time.Time
}

// NewTokenBucket creates a new token bucket
func NewTokenBucket(limit int, window time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:    float64(limit),
		remaining:   float64(limit),
		rate:        float64(limit) / window.Seconds(),
		lastUpdated: time.Now(),
	}
}

// Wait blocks until enough tokens are available or the context is canceled.
func (b *TokenBucket) Wait(ctx context.Context, cost int) error {
	waitDuration, err := b.reserve(ctx, cost)
	if err != nil {
		return err
	}

	if waitDuration > 0 {
		select {
		case <-time.After(waitDuration):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

func (b *TokenBucket) reserve(ctx context.Context, cost int) (time.Duration, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	b.advance(now)

	needed := float64(cost)
	b.remaining -= needed

	var waitDuration time.Duration
	if b.remaining < 0 {
		// Calculate time to refill the deficit
		deficit := -b.remaining
		waitDuration = time.Duration((deficit / b.rate) * float64(time.Second))
	}

	if deadline, ok := ctx.Deadline(); ok {
		if now.Add(waitDuration).After(deadline) {
			b.remaining += needed
			return 0, fmt.Errorf("rate: wait time exceeds context deadline")
		}
	}

	return waitDuration, nil
}

func (b *TokenBucket) advance(now time.Time) {
	elapsed := now.Sub(b.lastUpdated).Seconds()
	b.remaining += elapsed * b.rate
	if b.remaining > b.capacity {
		b.remaining = b.capacity
	}
	b.lastUpdated = now
}

// Sync updates the bucket state based on server headers
func (b *TokenBucket) Sync(remaining int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.remaining = float64(remaining)
	b.lastUpdated = time.Now()
}
