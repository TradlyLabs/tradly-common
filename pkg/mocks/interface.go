package mocks

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

type RedisClient interface {
	XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd
}
