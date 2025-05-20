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

type UsrUserRepository interface {
// GetUsers 获取所有用户信息
	GetUsers(context.Context) ([]domain.UsrUser, error)
	// GetUserByID 根据用户ID获取用户信息
	GetUserByID(context.Context, int64) (*domain.UsrUser, error)
	// CreateUser 创建新用户
	CreateUser(context.Context, *domain.UsrUser) (*domain.UsrUser, error)
	// UpdateUser 更新用户信息
	UpdateUser(context.Context, *domain.UsrUser) (*domain.UsrUser, error)
	// DeleteUser 删除用户
	DeleteUser(context.Context, int64) error
}