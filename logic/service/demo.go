package appService

import (
	"context"
	"github.com/faiz/llm-code-review/api/reply"
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/common/errcode"
	"github.com/faiz/llm-code-review/common/util"
	"github.com/faiz/llm-code-review/logic/domain"
	"github.com/faiz/llm-code-review/logic/repository"
)

type DemoAppServiceV1 struct {
	ds domainService.DemoDomainService
}

func NewDemoAppServiceV1(domainSvc domainService.DemoDomainService) *DemoAppServiceV1 {
	return &DemoAppServiceV1{
		ds: domainSvc,
	}
}

func (as *DemoAppServiceV1) GetAllIdentities(c context.Context) ([]int64, error) {
	domains, err := as.ds.GetDemos(c)
	if err != nil {
		return nil, err
	}
	res := make([]int64, 0, len(domains))
	for i := range domains {
		res = append(res, domains[i].Id)
	}
	return res, nil
}

func (as *DemoAppServiceV1) CreateDemoOrder(c context.Context, order *request.DemoOrderReq) (*reply.DemoOrder, error) {
	domainOrder := new(domain.DemoOrder)
	err := util.Convert(domainOrder, order)
	if err != nil {
		return nil, errcode.Wrap("fail to convert req.OrderReq to domain.Order", err)
	}
	domainOrder, err = as.ds.CreateDemoOrder(c, domainOrder)
	if err != nil {
		return nil, errcode.Wrap("fail to create order", err)
	}
	rep := new(reply.DemoOrder)
	if err = util.Convert(rep, domainOrder); err != nil {
		return nil, errcode.Wrap("fail to convert domain.Order to reply.Order", err)
	}
	return rep, nil
}
