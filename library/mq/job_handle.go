package mq

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

const (
	detailName = "mq:detail:"
	delayName  = "mq:delay"
	listName   = "mq:list:"
)

type redisHandel struct {
	client *redis.Client
}

func newRedisHandle(client *redis.Client) *redisHandel {
	return &redisHandel{
		client: client,
	}
}

func (handel *redisHandel) Insert(jobStruct *JobStruct) error {

	jobStructJson, err := json.Marshal(jobStruct)
	if err != nil {
		return err
	}

	jobStructKeyName := handel.getJobStructName(jobStruct.Topic, jobStruct.JobId)
	err = handel.client.SetNX(jobStructKeyName, jobStructJson, time.Duration(jobStruct.ExpireAt)*time.Second).Err()
	if err != nil {
		return err
	}

	if jobStruct.Delay < 1 {
		err = handel.client.LPush(handel.getListNameByTopic(jobStruct.Topic), jobStruct.JobId).Err()
	} else {

		jobStructJson, _ := json.Marshal(jobDelayStruct{
			JobId: jobStruct.JobId,
			Topic: jobStruct.Topic,
		})

		err = handel.client.ZAdd(handel.getDelayName(), redis.Z{
			Score:  float64(jobStruct.JobAt),
			Member: jobStructJson,
		}).Err()
	}

	if err != nil {
		_ = handel.client.Del(jobStructKeyName).Err()
	}

	return err
}

func (handel *redisHandel) Read(topic string, readAfterExpireIn time.Duration) (*JobStruct, error) {

	jobId, err := handel.client.LPop(handel.getListNameByTopic(topic)).Result()

	if jobId == "" && err != nil {
		return nil, nil
	}

	jobStructKey := handel.getJobStructName(topic, jobId)
	jobStructJson, err2 := handel.client.Get(jobStructKey).Result()
	if jobStructJson == "" && err2 != nil {
		return nil, nil
	}

	var jobStruct JobStruct
	err3 := json.Unmarshal([]byte(jobStructJson), &jobStruct)

	if err3 != nil {
		return nil, err3
	}

	_ = handel.client.Expire(jobStructKey, readAfterExpireIn).Err()
	return &jobStruct, nil
}

func (handel *redisHandel) ReadByMulti(topic string, limit int, readAfterExpireIn time.Duration) ([]*JobStruct, error) {

	var JobStructSet []*JobStruct
	var jobStruct *JobStruct
	var err error
	var jobId string
	for i := 0; i < limit; i++ {

		jobStruct = nil
		jobId, err = handel.client.RPop(handel.getListNameByTopic(topic)).Result()
		if err != nil {
			continue
		}

		if jobId == "" && err != nil {
			continue
		}

		jobStructKey := handel.getJobStructName(topic, jobId)
		jobStructJson, err2 := handel.client.Get(jobStructKey).Result()
		if jobStructJson == "" && err2 != nil {
			continue
		}

		err3 := json.Unmarshal([]byte(jobStructJson), &jobStruct)
		if err3 != nil {
			continue
		}

		_ = handel.client.Expire(jobStructKey, readAfterExpireIn).Err()
		JobStructSet = append(JobStructSet, jobStruct)
	}

	return JobStructSet, err
}

func (handel *redisHandel) DelayTicker() error {
	delayNameKey := handel.getDelayName()
	Z, err := handel.client.ZRangeByScoreWithScores(delayNameKey, redis.ZRangeBy{
		Min:    strconv.Itoa(0),
		Max:    strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		Offset: 0,
		Count:  20,
	}).Result()
	if err != nil {
		return err
	}
	log.Println("delayTicker-length:", len(Z))
	for _, zItem := range Z {
		log.Println("delayTicker-range:", zItem.Member)
		handel.client.ZRem(delayNameKey, zItem.Member)
		if int64(zItem.Score) < (time.Now().Unix() - 300) {
			log.Println("delayTicker-expire:", err.Error())
			continue
		}
		if "string" == fmt.Sprintf("%T", zItem.Member) {
			var jobDelayStruct jobDelayStruct
			err = json.Unmarshal([]byte(zItem.Member.(string)), &jobDelayStruct)
			if err != nil {
				log.Println("delayTicker-err:", err.Error())
				continue
			}
			_ = handel.client.LPush(handel.getListNameByTopic(jobDelayStruct.Topic), jobDelayStruct.JobId).Err()
		}
	}
	return nil
}

func (handel *redisHandel) getJobStructName(topic string, jobId string) string {
	return detailName + topic + ":" + jobId
}

func (handel *redisHandel) getDelayName() string {
	return delayName + ":list"
}

func (handel *redisHandel) getListNameByTopic(topic string) string {
	return listName + topic
}
