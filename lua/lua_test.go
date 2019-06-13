package lua

import (
	"fmt"
	"gin-study/app/core/extend/env"
	"gin-study/app/core/utils"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestLua(t *testing.T) {
	env.Init()
	utils.RedisInit()
	IncrByXX := redis.NewScript(`
		local exists
		local jobInsertFlag
		local jobDetailKey = tostring(KEYS[1])
		local jobDetailVal = tostring(ARGV[1])
		local jobDetailTTL = tonumber(ARGV[2])
		local jobListKey   = tostring(KEYS[2])
		local jobListVal   = tostring(ARGV[3])
		local jobDelayKey    = tostring(KEYS[3])
		local jobDelayVal    = tostring(ARGV[4])
		local jobDelayTime   = tonumber(ARGV[5])
		local jobAt			 = tonumber(ARGV[6])
		exists = redis.call("SET",jobDetailKey,jobDetailVal,"NX","EX",jobDetailTTL)
		if  type(exists) == 'table' then
			if exists["ok"]=="OK" then
				if tonumber(jobDelayTime) < 1 then
					jobInsertFlag=redis.call("LPUSH",jobListKey,jobListVal)
				else
					jobInsertFlag=redis.call("ZADD",jobDelayKey,jobAt,jobDelayVal)
				end
				if jobInsertFlag ==1 then
					return "ok"
				end
				redis.call("DEL",jobDetailKey)
				return "error:jobInsertFlag is error"
			end
		end
			return "error:jobId is exists"
	`)
	n, err := IncrByXX.Run(utils.RedisClient(), []string{"job:detail:yy", "job:list:yy", "job:delay:yy"}, "detail", 30, "list", "delay", 0, time.Now().Unix()).Result()
	fmt.Println(n, err)
}
