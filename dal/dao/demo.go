package dao

import (
	"context"
	"github.com/faiz/llm-code-review/common/errcode"
	"github.com/faiz/llm-code-review/common/util"
	"github.com/faiz/llm-code-review/dal/model"
	"github.com/faiz/llm-code-review/logic/domain"
)

type DemoDAOV1 struct {
}

func NewDemoDAO() *DemoDAOV1 {
	return &DemoDAOV1{}
}

func (dao *DemoDAOV1) FindAllDemo(c context.Context) ([]model.DemoOrder, error) {
	var res []model.DemoOrder
	err := DB().WithContext(c).Model(&model.DemoOrder{}).Find(&res).Error
	if err != nil {
		err = errcode.Wrap("query entity error", err)
	}
	return res, err
}

func (dao *DemoDAOV1) CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*model.DemoOrder, error) {
	mod := new(model.DemoOrder)
	err := util.Convert(mod, order)
	if err != nil {
		return nil, err
	}
	return mod, DB().WithContext(c).Model(&model.DemoOrder{}).Create(mod).Error
}
