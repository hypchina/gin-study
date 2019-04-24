package dao

import (
	"gin-study/app/core/utils"
	"github.com/go-redis/redis"
	"github.com/xormplus/xorm"
)

type connect struct {
	redisClient *redis.Client
	orm         *xorm.Engine
}

func connectInit() connect {
	return connect{
		redisClient: utils.RedisClient(),
		orm:         utils.ORM(),
	}
}
