package dao

import (
	"encoding/json"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"time"
)

type UserDao struct {
	base Base
}

func UserInstance() *UserDao {
	return &UserDao{
		Base{
			Orm: utils.ORM(),
		},
	}
}

func (dao *UserDao) GetByIdAndCache(id int64) (bool, *models.UcUser) {
	redisClient := utils.RedisClient()
	redisKey := enum.RedisUserKey(id)
	jsonStr, err := redisClient.Get(redisKey).Result()
	UcUser := &models.UcUser{}
	if err == nil && jsonStr != "" {
		err = json.Unmarshal([]byte(jsonStr), UcUser)
		if err == nil {
			return true, UcUser
		}
	}
	exists, err := dao.base.Orm.Id(id).Get(UcUser)
	if exists {
		jsonStr, _ := json.Marshal(UcUser)
		_, err := redisClient.Set(redisKey, jsonStr, time.Hour*2).Result()
		helper.CheckErr(err)
	}
	return exists, UcUser
}

func (dao *UserDao) GetById(id int64) (boolean bool, userModel *models.UcUser) {
	user := &models.UcUser{}
	exists, err := dao.base.Orm.Id(id).Get(user)
	helper.CheckErr(err)
	return exists, user
}

func (dao *UserDao) GetByEmail(email string) (boolean bool, userModel *models.UcUser) {
	user := &models.UcUser{}
	exists, err := dao.base.Orm.Where("email=? and status=?", email, 1).Get(user)
	helper.CheckErr(err)
	return exists, user
}

func (dao *UserDao) IsExistsByEmail(email string) bool {
	user := &models.UcUser{}
	countNum, err := dao.base.Orm.Where("email=? and status=?", email, 1).Count(user)
	helper.CheckErr(err)
	if countNum > 0 {
		return true
	}
	return false
}

func (dao *UserDao) CreateUser(user *models.UcUser) (boolean bool, id int64) {
	id, err := dao.base.Orm.InsertOne(user)
	return helper.CheckErr(err), id
}
