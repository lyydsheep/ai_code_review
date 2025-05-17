package cache

import (
	"context"
	"github.com/faiz/llm-code-review/config"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	_RedisClient *redis.Client
)

func RedisInit() {
	_RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.Redis.Addr,
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		PoolSize:     config.Redis.PoolSize,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolTimeout:  30 * time.Second,
	})
	if err := _RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func Redis() *redis.Client {
	return _RedisClient
}
