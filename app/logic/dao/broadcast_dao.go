package dao

import (
	"encoding/json"
	"gin-study/app/core/helper"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"time"
)

type BroadcastDao struct {
	connect connect
}

func BroadcastDaoInstance() *BroadcastDao {
	return &BroadcastDao{
		connect: connectInit(),
	}
}

func (dao *BroadcastDao) Insert(LogUserBroadcast *models.LogUserBroadcast) bool {
	_, err := dao.connect.orm.Insert(LogUserBroadcast)
	var redisKey string
	redisKey = enum.RedisBroadcastDetail(LogUserBroadcast.MsgId)
	redisVal, _ := json.Marshal(LogUserBroadcast)
	_time, err := helper.GetTimeByDate(LogUserBroadcast.ExpireAt)
	expireAt := time.Duration(_time.Unix()) * time.Second
	_ = dao.connect.redisClient.Set(redisKey, redisVal, expireAt).Err()
	return helper.CheckErr(err)
}

func (dao *BroadcastDao) GetByMsgId(msgId string) (*models.LogUserBroadcast, bool) {
	var redisKey string
	LogUserBroadcast := models.LogUserBroadcast{}
	redisKey = enum.RedisBroadcastDetail(msgId)
	redisVal, err := dao.connect.redisClient.Get(redisKey).Result()
	if err != nil {
		return nil, helper.CheckErr(err)
	}
	err = json.Unmarshal([]byte(redisVal), &LogUserBroadcast)
	return &LogUserBroadcast, helper.CheckErr(err)
}

func (dao *BroadcastDao) UpdateByMsgId() {

}
