package service

import (
	"gin-study/app/core/helper"
	"gin-study/app/http/filters"
	"gin-study/app/logic/dao"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"gin-study/resource"
)

type UserService struct {
	userDao *dao.UserDao
}

func UserInstance() *UserService {
	return &UserService{
		userDao: dao.UserInstance(),
	}
}

func (service *UserService) Create(filter filters.UserFilter) (userModel *models.User, err error) {
	if service.userDao.IsExistsByEmail(filter.Email) {
		return nil, helper.CreateErr(resource.Trans("email_exists"))
	}
	user := &models.User{
		UserName: filter.UserName,
		Password: filter.Password,
		Email:    filter.Email,
		State:    1,
	}
	boolean, _ := service.userDao.CreateUser(user)
	if boolean {
		return user, nil
	}
	return nil, helper.CreateErr(enum.StatusDataOpError)
}
