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

type PushInfoRepository interface {
	// Create 创建数据
	Create(context.Context, model.PushInfo) (model.PushInfo, error)
	// GetByRepoID 根据仓库ID获取推送信息
	GetByRepoID(context.Context, int64) (model.PushInfo, error)
	// Update 更新推送信息
	Update(context.Context, model.PushInfo) (model.PushInfo, error)
	// Delete 删除推送信息
	Delete(context.Context, int64) error
	// GetByUsername 通过 username 获取推送信息
	GetByUsername(context.Context, string) (model.PushInfo, error)
}
