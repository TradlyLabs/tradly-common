package redisdb

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

type Client interface {
	XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	SPublish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	PSubscribe(ctx context.Context, channels ...string) *redis.PubSub
	SSubscribe(ctx context.Context, channels ...string) *redis.PubSub
}
