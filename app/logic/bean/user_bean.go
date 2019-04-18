package bean

import (
	"gin-study/app/logic/entity"
	"gin-study/app/logic/models"
)

type AuthBean struct {
	UserEntity  *models.UcUser      `json:"user"`
	TokenEntity *entity.TokenEntity `json:"token"`
}

type _UserBean struct {
	UserEntity *models.UcUser `json:"user"`
}

func (AuthBean *AuthBean) UserBean() *_UserBean {
	return &_UserBean{
		UserEntity: AuthBean.UserEntity,
	}
}
