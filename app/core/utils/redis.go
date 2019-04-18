package utils

import (
	"gin-study/app/core/helper"
	"gin-study/conf"
	"github.com/go-redis/redis"
	"log"
)

var client *redis.Client

func RedisInit() {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Config().Redis.Addr,
		Password: conf.Config().Redis.Password,
		DB:       conf.Config().Redis.Db,
	})
	pong, err := client.Ping().Result()
	helper.CheckErr(err, true)
	log.Println("redis-server connect pong:" + pong)
}

func RedisClient() *redis.Client {
	if client != nil {
		return client
	}
	panic(helper.CreateErr("please connect redis server"))
}
