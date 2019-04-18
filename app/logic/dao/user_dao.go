package dao

import (
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/models"
)

type UserDao struct {
	Base
}

func UserInstance() *UserDao {
	return &UserDao{
		Base{
			Orm: utils.ORM(),
		},
	}
}

func (dao *UserDao) GetById(id int64) (boolean bool, userModel *models.UcUser) {
	user := &models.UcUser{}
	exists, err := dao.Orm.Id(id).Get(user)
	helper.CheckErr(err)
	return exists, user
}

func (dao *UserDao) GetByEmail(email string) (boolean bool, userModel *models.UcUser) {
	user := &models.UcUser{}
	exists, err := dao.Orm.Where("email=? and status=?", email, 1).Get(user)
	helper.CheckErr(err)
	return exists, user
}

func (dao *UserDao) IsExistsByEmail(email string) bool {
	user := &models.UcUser{}
	countNum, err := dao.Orm.Where("email=? and status=?", email, 1).Count(user)
	helper.CheckErr(err)
	if countNum > 0 {
		return true
	}
	return false
}

func (dao *UserDao) CreateUser(user *models.UcUser) (boolean bool, id int64) {
	id, err := dao.Orm.InsertOne(user)
	return helper.CheckErr(err), id
}
