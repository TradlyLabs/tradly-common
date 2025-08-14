package rate

import (
	"context"
	"fmt"
	"time"

	"github.com/TradlyLabs/tradly-common/pkg/services/redisdb"
)

// RedisLimiter implements the Limiter interface using Redis
// It uses a token bucket algorithm for rate limiting

type RedisLimiter struct {
	client   redisdb.Client
	key      string  // Redis key for this limiter
	rate     float64 // Rate of token generation per second
	capacity float64 // Maximum tokens in the bucket
}

// NewRedisLimiter creates a new RedisLimiter instance
func NewRedisLimiter(key string, rate float64, capacity float64) *RedisLimiter {
	return &RedisLimiter{
		client:   redisdb.Get(),
		key:      fmt.Sprintf("rate_limit:%s", key),
		rate:     rate,
		capacity: capacity,
	}
}

// Allow checks if a single request is allowed
func (l *RedisLimiter) Allow() bool {
	return l.AllowN(time.Now(), 1)
}

// AllowN checks if n requests are allowed
func (l *RedisLimiter) AllowN(t time.Time, n int) bool {
	ctx := context.Background()

	// Use Lua script for atomic operation
	script := `
	local key = KEYS[1]
	local rate = tonumber(ARGV[1])
	local capacity = tonumber(ARGV[2])
	local now = tonumber(ARGV[3])
	local n = tonumber(ARGV[4])

	-- Get current state
	local lastRefill = redis.call('hget', key, 'lastRefill')
	local tokens = redis.call('hget', key, 'tokens')

	-- Initialize if not exists
	if not lastRefill then
		lastRefill = now
		redis.call('hset', key, 'lastRefill', lastRefill)
	end

	if not tokens then
		tokens = capacity
		redis.call('hset', key, 'tokens', tokens)
	else
		tokens = tonumber(tokens)
	end

	lastRefill = tonumber(lastRefill)

	-- Calculate new tokens to add based on time elapsed
	local elapsed = now - lastRefill
	local newTokens = math.min(capacity, tokens + (elapsed * rate))

	-- Check if we can allow n requests
	if newTokens >= n then
		newTokens = newTokens - n
		-- Update state
		redis.call('hset', key, 'tokens', newTokens)
		redis.call('hset', key, 'lastRefill', now)
		return true
	end

	return false
	`

	result, err := l.client.Eval(ctx, script, []string{l.key}, l.rate, l.capacity, float64(t.UnixNano())/1e9, n).Bool()
	if err != nil {
		// Log error but allow request to prevent service failure
		fmt.Printf("Redis rate limiter error: %v\n", err)
		return true
	}

	return result
}

// Wait waits until a single request is allowed
func (l *RedisLimiter) Wait(ctx context.Context) error {
	return l.WaitN(ctx, 1)
}

// WaitN waits until n requests are allowed
func (l *RedisLimiter) WaitN(ctx context.Context, n int) error {
	for {
		if l.AllowN(time.Now(), n) {
			return nil
		}

		// Calculate wait time
		delay := l.Delay()
		if delay <= 0 {
			delay = time.Millisecond * 10
		}

		// Wait with context
		timer := time.NewTimer(delay)
		select {
		case <-timer.C:
			// Continue waiting
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		}
	}
}

// Tokens returns the current number of available tokens
func (l *RedisLimiter) Tokens() float64 {
	return l.TokensAt(time.Now())
}

// TokensAt returns the number of available tokens at a specific time
func (l *RedisLimiter) TokensAt(t time.Time) float64 {
	ctx := context.Background()

	script := `
	local key = KEYS[1]
	local rate = tonumber(ARGV[1])
	local capacity = tonumber(ARGV[2])
	local now = tonumber(ARGV[3])

	-- Get current state
	local lastRefill = redis.call('hget', key, 'lastRefill')
	local tokens = redis.call('hget', key, 'tokens')

	-- Initialize if not exists
	if not lastRefill then
		lastRefill = now
	end

	if not tokens then
		tokens = capacity
	else
		tokens = tonumber(tokens)
	end

	lastRefill = tonumber(lastRefill)

	-- Calculate new tokens to add based on time elapsed
	local elapsed = now - lastRefill
	local newTokens = math.min(capacity, tokens + (elapsed * rate))

	return newTokens
	`

	result, err := l.client.Eval(ctx, script, []string{l.key}, l.rate, l.capacity, float64(t.UnixNano())/1e9).Float64()
	if err != nil {
		fmt.Printf("Redis rate limiter error: %v\n", err)
		return l.capacity // Return maximum capacity on error
	}

	return result
}

// DelayFrom calculates the delay needed until the next token is available
func (l *RedisLimiter) DelayFrom(t time.Time) time.Duration {
	tokens := l.TokensAt(t)
	if tokens >= 1 {
		return 0
	}

	// Calculate how much time until we have one token
	return time.Duration((1 - tokens) / l.rate * float64(time.Second))
}

// Delay calculates the delay needed until the next token is available
func (l *RedisLimiter) Delay() time.Duration {
	return l.DelayFrom(time.Now())
}
