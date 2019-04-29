package main

import (
	"fmt"
	"gin-study/app/core/extend/env"
	"gin-study/app/core/utils"
	"gin-study/app/logic/enum"
	"gin-study/library/mq"
)

func main() {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := mq.NewJob(utils.RedisClient())
	defer utils.RedisClient().Close()
	defer utils.ORM().Close()
	job.DelayTicker()
	job.Subscribe(enum.TagJobTopicBroadcast, func(jobStruct *mq.JobStruct, e error) {
		if e != nil {
			fmt.Println("subscribe", e.Error())
			return
		}
		fmt.Println(jobStruct.Body)
	})
}
