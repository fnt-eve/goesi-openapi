package ratelimit

import (
	"sync"
	"time"
)

// BucketKey identifies a specific rate limit bucket (e.g., "group:user_id")
type BucketKey string

// BucketStore manages a collection of token buckets
type BucketStore struct {
	mu      sync.RWMutex
	buckets map[BucketKey]*TokenBucket
}

// NewBucketStore creates a new store
func NewBucketStore() *BucketStore {
	return &BucketStore{
		buckets: make(map[BucketKey]*TokenBucket),
	}
}

// GetBucket retrieves or creates a bucket
func (s *BucketStore) GetBucket(key BucketKey, limit int, window time.Duration) *TokenBucket {
	s.mu.RLock()
	bucket, exists := s.buckets[key]
	s.mu.RUnlock()

	if exists {
		return bucket
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Double check
	if bucket, exists := s.buckets[key]; exists {
		return bucket
	}

	bucket = NewTokenBucket(limit, window)
	s.buckets[key] = bucket
	return bucket
}
