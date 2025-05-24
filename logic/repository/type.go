package repository

import (
	"context"
	"github.com/faiz/llm-code-review/dal/model"
	"github.com/faiz/llm-code-review/logic/domain"
)

type UsrUserRepository interface {
	// GetUsers 获取所有用户信息
	GetUsers(context.Context) ([]domain.UsrUser, error)
	// GetUserByID 根据用户ID获取用户信息
	GetUserByID(context.Context, int64) (domain.UsrUser, error)
	// GetUserByUsername 根据用户名获取
	GetUserByUsername(ctx context.Context, username string) (model.UsrUser, error)
	// CreateUser 创建新用户
	CreateUser(context.Context, domain.UsrUser) (domain.UsrUser, error)
	// UpdateUser 更新用户信息
	UpdateUser(context.Context, domain.UsrUser) (domain.UsrUser, error)
	// DeleteUser 删除用户
	DeleteUser(context.Context, int64) error
}
