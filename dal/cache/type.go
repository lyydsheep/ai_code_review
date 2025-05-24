package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/faiz/llm-code-review/common/enum"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/logic/domain"
)

type Cache interface {
	Get(ctx context.Context, orderId string) (*domain.DemoOrder, error)
	Set(ctx context.Context, order *domain.DemoOrder) error
}

type DemoCacheV1 struct {
}

func (r *DemoCacheV1) Get(ctx context.Context, orderId string) (*domain.DemoOrder, error) {
	key := fmt.Sprintf(enum.REDIS_KEY_DEMO_ORDER_DETAIL, orderId)
	str, err := Redis().Get(ctx, key).Result()
	if err != nil {
		log.New(ctx).Warn("获取 order 缓存失败", "orderId", orderId)
		return nil, errcode.Wrap("获取order缓存失败", err)
	}
	order := new(domain.DemoOrder)
	if err = json.Unmarshal([]byte(str), order); err != nil {
		log.New(ctx).Error("unmarshal 出错", "orderId", orderId, "err", err)
		return nil, errcode.Wrap("解析缓存失败", err)
	}
	return order, nil
}

func (r *DemoCacheV1) Set(ctx context.Context, order *domain.DemoOrder) error {
	key := fmt.Sprintf(enum.REDIS_KEY_DEMO_ORDER_DETAIL, order.OrderId)
	str, err := json.Marshal(order)
	if err != nil {
		log.New(ctx).Error("marshal 出错", "orderId", order.OrderId, "err", err)
		return errcode.Wrap("解析缓存失败", err)
	}
	if err = Redis().Set(ctx, key, string(str), 0).Err(); err != nil {
		log.New(ctx).Error("设置缓存失败", "orderId", order.OrderId, "err", err)
		return errcode.Wrap("设置缓存失败", err)
	}
	return nil
}

func NewCacheV1() *DemoCacheV1 {
	return &DemoCacheV1{}
}
