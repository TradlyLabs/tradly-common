package rate

import (
	"context"
	"testing"
	"time"
)

func TestRedisLimiter_Allow(t *testing.T) {
	// Create a limiter with rate 10 per second and capacity 10
	limiter := NewRedisLimiter("test-allow", 10, 10)

	// Test allowing 10 requests
	for i := 0; i < 10; i++ {
		if !limiter.Allow() {
			t.Errorf("Expected request %d to be allowed", i)
		}
	}

	// The 11th request should be denied
	if limiter.Allow() {
		t.Error("Expected 11th request to be denied")
	}

	// Wait for tokens to refill
	time.Sleep(1 * time.Second)

	// Now should allow 10 more requests
	for i := 0; i < 10; i++ {
		if !limiter.Allow() {
			t.Errorf("Expected request %d to be allowed after refill", i)
		}
	}
}

func TestRedisLimiter_AllowN(t *testing.T) {
	// Create a limiter with rate 5 per second and capacity 10
	limiter := NewRedisLimiter("test-allow-n", 5, 10)

	// Test allowing 5 requests
	if !limiter.AllowN(time.Now(), 5) {
		t.Error("Expected 5 requests to be allowed")
	}

	// Test allowing 5 more requests
	if !limiter.AllowN(time.Now(), 5) {
		t.Error("Expected 5 more requests to be allowed")
	}

	// Test allowing 1 more request (should be denied)
	if limiter.AllowN(time.Now(), 1) {
		t.Error("Expected 1 more request to be denied")
	}

	// Wait for tokens to refill
	time.Sleep(1 * time.Second)

	// Should allow 5 more requests
	if !limiter.AllowN(time.Now(), 5) {
		t.Error("Expected 5 requests to be allowed after refill")
	}
}

func TestRedisLimiter_Tokens(t *testing.T) {
	// Create a limiter with rate 10 per second and capacity 10
	limiter := NewRedisLimiter("test-tokens", 10, 10)

	// Initial tokens should be 10
	tokens := limiter.Tokens()
	if tokens != 10 {
		t.Errorf("Expected 10 tokens, got %.2f", tokens)
	}

	// Allow 3 requests
	limiter.AllowN(time.Now(), 3)

	// Tokens should be 7
	tokens = limiter.Tokens()
	if tokens != 7 {
		t.Errorf("Expected 7 tokens, got %.2f", tokens)
	}

	// Wait for 0.5 seconds, should gain 5 tokens
	time.Sleep(500 * time.Millisecond)
	tokens = limiter.Tokens()
	if tokens != 10 {
		t.Errorf("Expected 10 tokens after refill, got %.2f", tokens)
	}
}

func TestRedisLimiter_Delay(t *testing.T) {
	// Create a limiter with rate 1 per second and capacity 1
	limiter := NewRedisLimiter("test-delay", 1, 1)

	// Initial delay should be 0
	delay := limiter.Delay()
	if delay != 0 {
		t.Errorf("Expected 0 delay, got %v", delay)
	}

	// Allow 1 request
	limiter.Allow()

	// Delay should be around 1 second
	delay = limiter.Delay()
	if delay < 900*time.Millisecond || delay > 1100*time.Millisecond {
		t.Errorf("Expected delay around 1 second, got %v", delay)
	}
}

func TestRedisLimiter_Wait(t *testing.T) {
	// Create a limiter with rate 1 per second and capacity 1
	limiter := NewRedisLimiter("test-wait", 1, 1)

	// Allow 1 request to exhaust the bucket
	limiter.Allow()

	// Start waiting
	start := time.Now()
	ctx := context.Background()
	err := limiter.Wait(ctx)
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Elapsed time should be around 1 second
	if elapsed < 900*time.Millisecond || elapsed > 1100*time.Millisecond {
		t.Errorf("Expected wait time around 1 second, got %v", elapsed)
	}
}
