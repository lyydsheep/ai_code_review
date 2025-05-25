package cache

import (
	"context"
	"github.com/faiz/llm-code-review/common/errcode"
	"sync"
)

type MemCache struct {
	cache sync.Map
}

func (r *MemCache) Get(ctx context.Context, key string) (string, error) {
	val, ok := r.cache.Load(key)
	if !ok {
		return "", errcode.ErrNotFound
	}
	return val.(string), nil
}

func (r *MemCache) Set(ctx context.Context, key string, val string, opts ...Option) error {
	option := option{
		Expiration: defaultExpiration,
	}
	for _, opt := range opts {
		opt.apply(&option)
	}
	r.cache.Store(key, val)
	if option.Expiration > 0 {
		// TODO
		// 定时删除
	}
	return nil
}

func NewMemCache() Cache {
	return &MemCache{}
}
