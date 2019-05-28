package main

import (
	"encoding/json"
	"gin-study/app/core/extend/env"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"gin-study/library/mq"
	"log"
	"time"
)

func main() {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := mq.NewJob(utils.RedisClient())
	defer utils.RedisClient().Close()
	defer utils.ORM().Close()
	go job.DelayTicker()
	job.Subscribe(enum.TagJobTopicBroadcast, func(jobStruct mq.JobStruct, e error) {
		if e != nil {
			log.Println(e)
			helper.CheckErr(e)
			return
		}
		var broadcast models.LogUserBroadcast
		err := json.Unmarshal([]byte(jobStruct.Body), &broadcast)
		if err != nil {
			log.Println("subscribe:parse-broadcast:error", err.Error())
			return
		}
		log.Println("subscribe:broadcast:", time.Now().Unix(), broadcast.MsgBody)
	})
}
