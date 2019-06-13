package mq

import (
	"fmt"
	"gin-study/app/core/extend/env"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"strconv"
	"testing"
)

func TestNewJob(t *testing.T) {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := NewJob(utils.RedisClient())
	j := 30
	for i := 0; i < 30; i++ {
		broadcast := &models.LogUserBroadcast{
			MsgId:     helper.CreateUUID(),
			FromUid:   "007",
			ToUid:     "9527",
			MsgLevel:  1,
			MsgType:   2,
			MsgBody:   "你好第" + strconv.Itoa(i) + "号",
			ReadState: 0,
			ExpireAt:  helper.GetDateByFormat(),
			NotifyAt:  helper.GetDateByFormat(),
			ReadAt:    helper.GetDefaultDate(),
			CreatedAt: helper.GetDateByFormat(),
		}
		_, err := job.Publish(enum.TagJobTopicBroadcast, 3, broadcast, "tag:default")
		j--
		fmt.Println(broadcast.MsgBody, err)
	}
}

func TestJob_Publish(t *testing.T) {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := NewJob(utils.RedisClient())
	job.Subscribe(enum.TagJobTopicBroadcast, func(jobStruct JobStruct) (isAck bool) {
		return true
	})
}
