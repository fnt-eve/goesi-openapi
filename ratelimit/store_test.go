package ratelimit

import (
	"testing"
	"time"
)

func TestBucketStore(t *testing.T) {
	store := NewBucketStore()
	key := BucketKey("test")

	b1 := store.GetBucket(key, 10, time.Second)
	b2 := store.GetBucket(key, 10, time.Second)

	if b1 != b2 {
		t.Errorf("GetBucket returned different instances for same key")
	}
}

func TestBucketStore_Concurrency(t *testing.T) {
	store := NewBucketStore()
	key := BucketKey("concurrent")

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			store.GetBucket(key, 10, time.Second)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify we have 1 bucket
	store.mu.RLock()
	count := len(store.buckets)
	store.mu.RUnlock()

	if count != 1 {
		t.Errorf("Expected 1 bucket, got %d", count)
	}
}
