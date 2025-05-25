package cache

import (
	"context"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
)

type RedisCache struct {
}

const (
	defaultExpiration = 0
)

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	str, err := Redis().Get(ctx, key).Result()
	if err != nil {
		log.New(ctx).Warn("获取 order 缓存失败", "key", key)
		return "", errcode.ErrServer.WithCause(err).AppendMsg("获取缓存失败")
	}
	return str, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, val string, opts ...Option) error {
	option := &option{
		Expiration: 0,
	}
	for _, opt := range opts {
		opt.apply(option)
	}

	if err := Redis().Set(ctx, key, val, 0).Err(); err != nil {
		log.New(ctx).Error("设置缓存失败", "key", key, "val", val, "expiration", option.Expiration, "err", err)
		return errcode.ErrServer.WithCause(err).AppendMsg("设置缓存失败")
	}
	return nil
}

func NewRedisCache() Cache {
	return &RedisCache{}
}
