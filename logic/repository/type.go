package domainService

import (
	"context"
	"github.com/faiz/llm-code-review/logic/domain"
)

// DemoDomainService 保持依赖注入风格
type DemoDomainService interface {
	GetDemos(context.Context) ([]domain.DemoOrder, error)
	CreateDemoOrder(context.Context, *domain.DemoOrder) (*domain.DemoOrder, error)
}
