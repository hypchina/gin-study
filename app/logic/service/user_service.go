package service

import (
	"gin-study/app/core/helper"
	"gin-study/app/http/filters"
	"gin-study/app/logic/bean"
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

func (service *UserService) Create(filter filters.UserRegister) (err error) {
	if service.userDao.IsExistsByEmail(filter.Email) {
		return helper.CreateErr(resource.Trans("email_exists"))
	}
	user := &models.UcUser{
		UserName:  filter.UserName,
		Password:  helper.CreatePwd(filter.Password),
		TokenSlat: helper.CreateUUID(),
		Email:     filter.Email,
		Status:    1,
		UpdatedAt: helper.GetDateByFormat(),
		CreatedAt: helper.GetDateByFormat(),
	}
	boolean, _ := service.userDao.CreateUser(user)
	if boolean {
		return nil
	}
	return helper.CreateErr(enum.StatusDataOpError)
}

func (service *UserService) CreateAuth(filter filters.UserLogin) (authBean *bean.AuthBean, err error) {

	exists, UserModel := service.userDao.GetByEmail(filter.Email)
	if !exists {
		return nil, helper.CreateErr(resource.Trans("user_not_exists"))
	}

	if UserModel.Password != helper.CreatePwd(filter.Password) {
		return nil, helper.CreateErr(resource.Trans("pwd_error"))
	}

	TokenDao := dao.TokenDaoInstance()
	ok, TokenEntity := TokenDao.CreateAndStore(UserModel.Id)
	if !ok {
		return nil, helper.CreateErr(resource.Trans("system_error"))
	}

	return &bean.AuthBean{
		UserEntity:  UserModel,
		TokenEntity: TokenEntity,
	}, nil

}

func (service *UserService) GetAuth(clientId string) (bool, *bean.AuthBean) {
	ok, TokenEntity := dao.TokenDaoInstance().GetToken(clientId)
	if !ok {
		return false, nil
	}
	_, UserModel := dao.UserInstance().GetById(TokenEntity.Uid)
	return true, &bean.AuthBean{
		UserEntity:  UserModel,
		TokenEntity: TokenEntity,
	}
}
