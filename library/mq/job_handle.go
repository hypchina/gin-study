package mq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const (
	detailName = "mq:detail:"
	delayName  = "mq:delay:"
	listName   = "mq:list:"
	ackName    = "mq:ack:"
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

	jobDelayStructJson, _ := json.Marshal(jobDelayStruct{
		JobId: jobStruct.JobId,
		Topic: jobStruct.Topic,
	})

	script := redis.NewScript(publishLuaScript)
	resp, err := script.Run(handel.client, []string{
		handel.getJobStructName(jobStruct.Topic, jobStruct.JobId),
		handel.getListNameByTopic(jobStruct.Topic),
		handel.getDelayName(jobStruct.Topic),
	}, jobStructJson, jobStruct.ExpireAt-time.Now().Unix(), jobStruct.JobId, jobDelayStructJson, jobStruct.Delay, jobStruct.JobAt).Result()

	if err == nil {
		respStr, ok := resp.(string)
		if ok && respStr != "ok" {
			return errors.New(respStr)
		}
	}

	return err
}

func (handel *redisHandel) Read(topic string, readAfterExpireIn time.Duration) (*JobStruct, error) {

	jobId, err := handel.client.RPopLPush(handel.getListNameByTopic(topic), handel.getWaitAckNameByTopic(topic)).Result()
	if jobId == "" && err != nil && err.Error() == "redis: nil" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	jobStructKey := handel.getJobStructName(topic, jobId)
	jobStructJson, err := handel.client.Get(jobStructKey).Result()

	if jobStructJson == "" && (err != nil && err.Error() == "redis: nil") {
		handel.Ack(topic, jobId)
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var jobStruct JobStruct
	err = json.Unmarshal([]byte(jobStructJson), &jobStruct)

	if err != nil {
		return nil, err
	}

	return &jobStruct, nil
}

func (handel *redisHandel) Ack(topic string, jobId string) {
	_ = handel.client.LRem(handel.getWaitAckNameByTopic(topic), 1, jobId).Err()
	_ = handel.client.Del(handel.getJobStructName(topic, jobId)).Err()
}

func (handel *redisHandel) Retry(topic string) (error, bool) {

	retryLockName := handel.getRetryLockNameByTopic(topic)
	listName := handel.getListNameByTopic(topic)
	waitAckName := handel.getWaitAckNameByTopic(topic)

	lock, _ := handel.client.SetNX(retryLockName, 1, time.Second*10).Result()
	if !lock {
		fmt.Println(fmt.Sprintf("retry %s is lock", topic))
		return nil, false
	}

	for i := 0; i < 50; i++ {

		jobId, err := handel.client.RPopLPush(waitAckName, listName).Result()
		if err != nil && err.Error() == "redis: nil" {
			handel.client.Del(retryLockName)
			return nil, false
		}

		if err != nil {
			fmt.Println("retry err xxx", jobId, err)
			handel.client.Del(retryLockName)
			return err, false
		}

		fmt.Printf("rollback jobId %s \n", jobId)
	}

	handel.client.Del(retryLockName)
	return nil, true
}

func (handel *redisHandel) DelayTicker(maxQuerySize int64, topic string) (error, bool) {
	lock, _ := handel.client.SetNX(handel.getDelayLockName(topic), 1, time.Second*10).Result()
	if !lock {
		return errors.New(fmt.Sprintf("topic %s is lock", topic)), false
	}
	delayNameKey := handel.getDelayName(topic)
	Z, err := handel.client.ZRangeByScoreWithScores(delayNameKey, redis.ZRangeBy{
		Min:    strconv.Itoa(0),
		Max:    strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		Offset: 0,
		Count:  maxQuerySize,
	}).Result()

	if err != nil && err.Error() == "redis: nil" {
		_ = handel.client.Del(handel.getDelayLockName(topic)).Err()
		return nil, false
	}

	if err != nil {
		return err, false
	}

	index := 0
	for _, zItem := range Z {

		index++
		if int64(zItem.Score) < (time.Now().Unix() - 600) {
			handel.client.ZRem(delayNameKey, zItem.Member)
			continue
		}

		var jobDelayStruct jobDelayStruct
		err = json.Unmarshal([]byte(zItem.Member.(string)), &jobDelayStruct)
		if err != nil {
			handel.client.ZRem(delayNameKey, zItem.Member)
			continue
		}

		popExists, err := handel.client.SetNX(handel.getDelayPopName(jobDelayStruct.JobId), 1, time.Second*120).Result()
		if err != nil {
			continue
		}

		if !popExists {
			continue
		}

		err = handel.client.LPush(handel.getListNameByTopic(jobDelayStruct.Topic), jobDelayStruct.JobId).Err()
		if err != nil {
			continue
		}

		handel.client.ZRem(delayNameKey, zItem.Member)
	}

	_ = handel.client.Del(handel.getDelayLockName(topic)).Err()
	if index > 0 {
		return nil, true
	}

	return nil, false
}

func (handel *redisHandel) getJobStructName(topic string, jobId string) string {
	return detailName + topic + ":" + jobId
}

func (handel *redisHandel) getDelayName(topic string) string {
	return delayName + "list:" + topic
}

func (handel *redisHandel) getDelayPopName(jobId string) string {
	return delayName + "pop:" + jobId
}

func (handel *redisHandel) getDelayLockName(topic string) string {
	return delayName + "lock:" + topic
}

func (handel *redisHandel) getListNameByTopic(topic string) string {
	return listName + topic
}

func (handel *redisHandel) getWaitAckNameByTopic(topic string) string {
	return ackName + "wait:" + topic
}

func (handel *redisHandel) getRetryLockNameByTopic(topic string) string {
	return ackName + "retry_lock:" + topic
}
