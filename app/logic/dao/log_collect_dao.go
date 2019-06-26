package dao

import (
	"encoding/json"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
)

type LogCollectDao struct {
	connect connect
}

func LogCollectDaoInstance() *LogCollectDao {
	return &LogCollectDao{
		connect: connectInit(),
	}
}

func (dao *LogCollectDao) Insert(logCollectModel models.LogCollect) error {
	logCollectJson, err := json.Marshal(logCollectModel)
	if err != nil {
		return err
	}
	queueName := enum.RedisLogCollectQueue(logCollectModel.AppName)
	return dao.connect.redisClient.LPush(queueName, string(logCollectJson)).Err()
}
