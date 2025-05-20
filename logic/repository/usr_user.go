package domainService

import (
	"context"
	"github.com/faiz/llm-code-review/common/errcode"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/common/util"
	"github.com/faiz/llm-code-review/dal/cache"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/faiz/llm-code-review/logic/domain"
)

type DemoDomainServiceV1 struct {
	Dao   dao.DemoDAO
	Cache cache.DemoCache
}

func NewDemoDomainServiceV1(d dao.DemoDAO, cache cache.DemoCache) *DemoDomainServiceV1 {
	return &DemoDomainServiceV1{
		Dao:   d,
		Cache: cache,
	}
}

func (ds *DemoDomainServiceV1) GetDemos(c context.Context) ([]domain.DemoOrder, error) {
	demos, err := ds.Dao.FindAllDemo(c)
	if err != nil {
		err = errcode.Wrap("query entity error", err)
	}
	res := make([]domain.DemoOrder, 0, len(demos))
	for i := range demos {
		res = append(res, domain.DemoOrder{
			Id:           demos[i].Id,
			OrderId:      demos[i].OrderId,
			UserId:       demos[i].UserId,
			OrderGoodsId: demos[i].OrderGoodsId,
			BillMoney:    demos[i].BillMoney,
			State:        demos[i].State,
			PaidAt:       demos[i].PaidAt,
		})
	}
	return res, err
}

// CreateDemoOrder 创建订单
// 核心业务逻辑
func (ds *DemoDomainServiceV1) CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*domain.DemoOrder, error) {
	order.OrderId = "this is random Id"
	demo, err := ds.Dao.CreateDemoOrder(c, order)
	if err != nil {
		return nil, errcode.Wrap("create entity error", err)
	}
	// do something...
	if err = util.Convert(order, demo); err != nil {
		return nil, errcode.Wrap("copy entity error", err)
	}
	if err = ds.Cache.Set(c, order); err != nil {
		return nil, errcode.Wrap("cache entity error", err)
	}
	log.New(c).Info("create order success", "orderId", order.OrderId)
	data, err := ds.Cache.Get(c, order.OrderId)
	if err != nil {
		return nil, errcode.Wrap("get entity error", err)
	}
	log.New(c).Info("get order from cache", "orderId", order.OrderId, "data", data)
	return order, nil
}
