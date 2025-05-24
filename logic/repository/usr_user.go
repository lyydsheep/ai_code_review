package repository

import (
	"context"
	"github.com/faiz/llm-code-review/common/errcode"
	"github.com/faiz/llm-code-review/dal/cache"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/faiz/llm-code-review/dal/model"
	"github.com/faiz/llm-code-review/logic/domain"
)

type UsrUserRepositoryV1 struct {
	Query *dao.Query
	Cache cache.Cache
}

func (repo *UsrUserRepositoryV1) GetUserByUsername(ctx context.Context, username string) (model.UsrUser, error) {
	qUser := repo.Query.UsrUser
	entity, err := qUser.WithContext(ctx).Where(qUser.Username.Eq(username)).First()
	if err != nil {
		return model.UsrUser{}, errcode.Wrap("获取用户失败", err)
	}
	return *entity, nil
}

func (repo *UsrUserRepositoryV1) GetUsers(ctx context.Context) ([]domain.UsrUser, error) {
	qUser := repo.Query.UsrUser
	entities, err := qUser.WithContext(ctx).Find()
	if err != nil {
		return []domain.UsrUser{}, errcode.Wrap("获取用户列表失败", err)
	}
	domains := make([]domain.UsrUser, len(entities))
	for i, entity := range entities {
		domains[i] = domain.UsrUserEntityToDomain(*entity)
	}
	return domains, nil
}

func (repo *UsrUserRepositoryV1) GetUserByID(ctx context.Context, id int64) (domain.UsrUser, error) {
	qUser := repo.Query.UsrUser
	entity, err := qUser.WithContext(ctx).Where(qUser.ID.Eq(id)).First()
	if err != nil {
		return domain.UsrUser{}, errcode.Wrap("获取用户失败", err)
	}
	return domain.UsrUserEntityToDomain(*entity), nil
}

func (repo *UsrUserRepositoryV1) CreateUser(ctx context.Context, user domain.UsrUser) (domain.UsrUser, error) {
	entity := domain.UsrUserDomainToEntity(user)
	qUser := repo.Query.UsrUser
	err := qUser.WithContext(ctx).Create(&entity)
	if err != nil {
		return domain.UsrUser{}, errcode.Wrap("创建用户失败", err)
	}
	user = domain.UsrUserEntityToDomain(entity)
	return user, nil
}

func (repo *UsrUserRepositoryV1) UpdateUser(ctx context.Context, user domain.UsrUser) (domain.UsrUser, error) {
	entity := domain.UsrUserDomainToEntity(user)
	qUser := repo.Query.UsrUser
	_, err := qUser.WithContext(ctx).Where(qUser.ID.Eq(entity.ID)).Updates(&entity)
	if err != nil {
		return domain.UsrUser{}, errcode.Wrap("更新用户失败", err)
	}
	return user, nil
}

func (repo *UsrUserRepositoryV1) DeleteUser(ctx context.Context, id int64) error {
	qUser := repo.Query.UsrUser
	_, err := qUser.WithContext(ctx).Where(qUser.ID.Eq(id)).Delete()
	if err != nil {
		return errcode.Wrap("删除用户失败", err)
	}
	return nil
}

func NewUsrUserRepositoryV1(q *dao.Query, c cache.Cache) UsrUserRepository {
	return &UsrUserRepositoryV1{
		Query: q,
		Cache: c,
	}
}
