package dao

import (
	"gin-study/app/core/helper"
	"gin-study/app/logic/models"
)

type LogSysRequestDao struct {
	connect connect
}

func LogSysRequestDaoInstance() *LogSysRequestDao {
	return &LogSysRequestDao{
		connect: connectInit(),
	}
}

func (dao *LogSysRequestDao) Insert(request models.LogSysRequest) (bool, int64) {
	request.CreatedAt = helper.GetDateByFormatWithMs()
	id, err := dao.connect.orm.InsertOne(request)
	return helper.CheckErr(err), id
}
