package dao

import (
	"context"
	"github.com/faiz/llm-code-review/dal/model"
	"github.com/faiz/llm-code-review/logic/domain"
)

type DemoDAO interface {
	FindAllDemo(c context.Context) ([]model.DemoOrder, error)
	CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*model.DemoOrder, error)
}
