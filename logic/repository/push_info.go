package repository

import (
	"context"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/faiz/llm-code-review/dal/model"
)

type PushInfoRepositoryV1 struct {
	Query *dao.Query
}

func (repo *PushInfoRepositoryV1) Create(ctx context.Context, info model.PushInfo) (model.PushInfo, error) {
	qpushInfo := repo.Query.PushInfo
	err := qpushInfo.WithContext(ctx).Create(&info)
	return info, err
}

func (repo *PushInfoRepositoryV1) GetByRepoID(ctx context.Context, repoID int64) (model.PushInfo, error) {
	q := repo.Query.PushInfo
	info, err := q.WithContext(ctx).Where(q.ID.Eq(repoID)).First()
	if err != nil {
		return model.PushInfo{}, err
	}
	return *info, nil
}

func (repo *PushInfoRepositoryV1) Update(ctx context.Context, info model.PushInfo) (model.PushInfo, error) {
	q := repo.Query.PushInfo.WithContext(ctx)
	err := q.Save(&info)
	if err != nil {
		return model.PushInfo{}, err
	}
	return info, nil
}

func (repo *PushInfoRepositoryV1) Delete(ctx context.Context, id int64) error {
	q := repo.Query.PushInfo
	_, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Delete()
	return err
}

func (repo *PushInfoRepositoryV1) GetByUsername(ctx context.Context, username string) (model.PushInfo, error) {
	q := repo.Query.PushInfo
	info, err := q.WithContext(ctx).Where(q.Username.Eq(username)).First()
	if err != nil {
		return model.PushInfo{}, err
	}
	return *info, nil
}

func NewPushInfoRepositoryV1(q *dao.Query) PushInfoRepository {
	return &PushInfoRepositoryV1{
		Query: q,
	}
}
