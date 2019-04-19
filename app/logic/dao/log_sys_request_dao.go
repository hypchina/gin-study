package dao

import (
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/models"
)

type LogSysRequestDao struct {
	base Base
}

func LogSysRequestDaoInstance() *LogSysRequestDao {
	return &LogSysRequestDao{
		Base{
			Orm: utils.ORM(),
		},
	}
}

func (dao *LogSysRequestDao) Insert(request models.LogSysRequest) (bool, int64) {
	id, err := dao.base.Orm.InsertOne(request)
	return helper.CheckErr(err), id
}
