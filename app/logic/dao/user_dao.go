package dao

import (
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/models"
)

type UserDao struct{}

func UserInstance() *UserDao {
	return &UserDao{}
}

func (dao *UserDao) GetById(id int64) (boolean bool, userModel *models.User) {
	user := &models.User{}
	exists, err := utils.ORM.Id(id).Get(user)
	helper.CheckErr(err)
	return exists, user
}

func (dao *UserDao) GetByEmail(email string) (boolean bool, userModel *models.User) {
	user := &models.User{}
	exists, err := utils.ORM.Where("email=? and state=?", email, 1).Get(user)
	helper.CheckErr(err)
	return exists, user
}

func (dao *UserDao) IsExistsByEmail(email string) bool {
	user := &models.User{}
	countNum, err := utils.ORM.Where("email=? and state=?", email, 1).Count(user)
	helper.CheckErr(err)
	if countNum > 0 {
		return true
	}
	return false
}

func (dao *UserDao) CreateUser(user *models.User) (boolean bool, id int64) {
	id, err := utils.ORM.InsertOne(user)
	return helper.CheckErr(err), id
}
