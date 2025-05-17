package appService

import (
	"context"
	"github.com/faiz/llm-code-review/api/reply"
	"github.com/faiz/llm-code-review/api/request"
)

type DemoAppService interface {
	GetAllIdentities(c context.Context) ([]int64, error)
	CreateDemoOrder(c context.Context, order *request.DemoOrderReq) (*reply.DemoOrder, error)
}
