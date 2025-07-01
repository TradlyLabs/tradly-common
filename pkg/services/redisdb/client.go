package redisdb

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

type Client interface {
	XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd
}
