package main

import (
	"encoding/json"
	"fmt"
	"gin-study/app/core/extend/env"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"gin-study/app/logic/service"
	"gin-study/library/mq"
)

func main() {

	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := mq.NewJob(utils.RedisClient())
	defer utils.RedisClient().Close()
	defer utils.ORM().Close()
	go job.DelayTicker()

	LogSysRequestServiceInstance := service.LogSysRequestServiceInstance()
	job.Subscribe(enum.TagJobTopicRequestLog, func(jobStruct mq.JobStruct, e error) {
		if !helper.CheckErr(e) {
			return
		}
		var LogSysRequest models.LogSysRequest
		err := json.Unmarshal([]byte(jobStruct.Body), &LogSysRequest)
		if !helper.CheckErr(err) {
			return
		}
		fmt.Println(LogSysRequest.RequestId)
		LogSysRequestServiceInstance.Insert(LogSysRequest)
	})
}
