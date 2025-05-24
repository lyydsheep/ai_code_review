package domain

import (
	"github.com/faiz/llm-code-review/dal/model"
)

type UsrUser struct {
	Username string `gorm:"column:username;type:varchar(36);comment:用户名" json:"username"`    // 用户名
	Email    string `gorm:"column:email;type:varchar(128);not null;comment:邮箱" json:"email"` // 邮箱
}

func UsrUserDomainToEntity(domain UsrUser) model.UsrUser {
	return model.UsrUser{
		Username: domain.Username,
		Email:    domain.Email,
	}
}

func UsrUserEntityToDomain(entity model.UsrUser) UsrUser {
	return UsrUser{
		Username: entity.Username,
		Email:    entity.Email,
	}
}
