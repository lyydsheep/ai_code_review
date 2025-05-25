package cache

import (
	"context"
	"time"
)

// 可以做一个 option 模式
// 过期时间，是否续期

type option struct {
	Expiration time.Duration
}

type Option interface {
	apply(opts *option)
}

type OptionFunc func(opts *option)

func (f OptionFunc) apply(opts *option) {
	f(opts)
}

func WithExpiration(expiration time.Duration) Option {
	return OptionFunc(func(opts *option) {
		opts.Expiration = expiration
	})
}

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, val string, opts ...Option) error
}
