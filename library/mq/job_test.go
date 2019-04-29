package mq

import (
	"fmt"
	"gin-study/app/core/extend/env"
	"gin-study/app/core/utils"
	"testing"
)

func TestNewJob(t *testing.T) {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := NewJob(utils.RedisClient())
	type XX struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	X := XX{Id: 1, Name: "xxx"}
	x, err := job.Publish("topic:x", 30, X, "tag:default")
	x, err = job.Publish("topic:x", 30, X, "tag:default")
	x, err = job.Publish("topic:x", 30, X, "tag:default")
	fmt.Println(x, err)
}

func TestJob_Publish(t *testing.T) {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := NewJob(utils.RedisClient())
	job.Subscribe("topic:x", func(jobStruct *JobStruct, e error) {
		fmt.Println("subscribe:", jobStruct, e)
	})
}

func TestJob_DelayTicker(t *testing.T) {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	job := NewJob(utils.RedisClient())
	job.DelayTicker()
}
