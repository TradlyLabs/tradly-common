package rate

import (
	"context"
	"time"
)

type Limiter interface {
	Allow() bool
	AllowN(t time.Time, n int) bool
	Wait(ctx context.Context) (err error)
	WaitN(ctx context.Context, n int) error
	Tokens() float64
	TokensAt(t time.Time) float64
	DelayFrom(t time.Time) time.Duration
	Delay() time.Duration
}
