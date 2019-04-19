package dao

import (
	"encoding/json"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/entity"
	"gin-study/app/logic/enum"
	"github.com/go-redis/redis"
	"time"
)

type tokenDao struct {
	redisClient *redis.Client
}

func TokenDaoInstance() *tokenDao {
	return &tokenDao{
		redisClient: utils.RedisClient(),
	}
}

func (dao *tokenDao) Create(uid int64) *entity.TokenEntity {
	var token = entity.TokenEntity{
		Uid:         uid,
		ClientId:    helper.CreateUUID(),
		ClientToken: helper.CreateUUID(),
		ExpireAt:    time.Now().Unix() + (86400 * 60),
	}
	return &token
}

func (dao *tokenDao) CreateAndSet(uid int64) (bool, *entity.TokenEntity) {
	var token = dao.Create(uid)
	return dao.Set(token), token
}

func (dao *tokenDao) Get(clientId string) (bool, *entity.TokenEntity) {
	redisKey := enum.RedisTokenKey(clientId)
	tokenStr, err := dao.redisClient.Get(redisKey).Result()
	var token entity.TokenEntity
	if err != nil {
		return false, nil
	}
	err = json.Unmarshal([]byte(tokenStr), &token)
	return helper.CheckErr(err), &token
}

func (dao *tokenDao) Set(token *entity.TokenEntity) bool {
	redisKey := enum.RedisTokenKey(token.ClientId)
	jsonStr, _ := json.Marshal(token)
	ok, err := dao.redisClient.SetNX(redisKey, jsonStr, time.Duration(token.ExpireAt)*time.Second).Result()
	return ok && helper.CheckErr(err)
}

func (dao *tokenDao) Del(clientId string) bool {
	redisKey := enum.RedisTokenKey(clientId)
	err := dao.redisClient.Del(redisKey).Err()
	if err != nil {
		return false
	}
	return true
}
