package main

import (
	"encoding/json"
	"gin-study/app/core/extend/env"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"gin-study/app/logic/service"
	"gin-study/library/mq"
	"strings"
)

func main() {

	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := mq.NewJob(utils.RedisClient())
	defer utils.RedisClient().Close()
	defer utils.ORM().Close()

	LogSysRequestServiceInstance := service.LogSysRequestServiceInstance()

	job.Subscribe(enum.TagJobTopicDispatchOrderRealTime, func(jobStruct mq.JobStruct) (isAck bool) {

		var LogSysRequest models.LogSysRequest
		err := json.Unmarshal([]byte(jobStruct.Body), &LogSysRequest)
		if !helper.CheckErr(err) {
			return true
		}

		_, err = LogSysRequestServiceInstance.Insert(LogSysRequest)
		if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
			return true
		}

		return helper.CheckErr(err)
	})
}
