package mq

import (
	"encoding/json"
	"fmt"
	"gin-study/app/core/helper"
	"github.com/go-redis/redis"
	"time"
)

type job struct {
	topic               string
	defaultExpireIn     time.Duration
	readAfterExpireIn   time.Duration
	delayTickerIn       time.Duration
	retryTickerIn       time.Duration
	subscribeTickerIn   time.Duration
	subscribeQueueLimit int64
	delayTickerLimit    int64
	handel              *redisHandel
}

func NewJob(redisClient *redis.Client) *job {
	return &job{
		subscribeQueueLimit: 500,
		delayTickerLimit:    300,
		defaultExpireIn:     time.Second * 86400 * 30,
		readAfterExpireIn:   time.Second * 60,
		subscribeTickerIn:   time.Second * 1,
		delayTickerIn:       time.Second * 2,
		retryTickerIn:       time.Second * 2,
		handel:              newRedisHandle(redisClient),
	}
}

func (job *job) SetSubscribeQueueLimit(subscribeQueueLimit int64) *job {
	if subscribeQueueLimit < 1 {
		subscribeQueueLimit = 1
	}
	job.subscribeQueueLimit = subscribeQueueLimit
	return job
}

func (job *job) Publish(topic string, delay int64, body interface{}, tag string) (*JobStruct, error) {
	var bodyStr string
	paramType := fmt.Sprintf("%T", body)
	if paramType == "string" {
		bodyStr = body.(string)
	} else {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyStr = string(bodyBytes)
	}
	jobStruct := &JobStruct{
		JobId:    createUUID(),
		Topic:    topic,
		JobAt:    time.Now().Add(time.Duration(delay)*time.Second).UnixNano() / 1e6,
		Body:     bodyStr,
		Tag:      tag,
		ExpireAt: time.Now().Add(job.defaultExpireIn).Unix(),
		Delay:    delay,
		CreateAt: getDateByFormat(),
	}
	err := job.handel.Insert(jobStruct)
	if err != nil {
		return nil, err
	}
	return jobStruct, nil
}

func (job *job) Subscribe(topic string, callback func(JobStruct) (isAck bool)) {

	go job.delayTicker(topic)
	go job.retryTicker(topic)

	for {
		startT := time.Now().Unix()
		var i int64
		i = 0

		for i = 0; i < job.subscribeQueueLimit; i++ {

			jobStruct, err := job.handel.Read(topic, job.readAfterExpireIn)
			if jobStruct == nil && err == nil {
				break
			}

			if err != nil {
				continue
			}

			if jobStruct == nil {
				continue
			}

			if callback(*jobStruct) {
				job.handel.Ack(topic, jobStruct.JobId)
			}
		}
		endT := time.Now().Unix()
		if endT-startT < 1 {
			if i > 0 {
				fmt.Println("subscribe sleep")
			}
			time.Sleep(job.subscribeTickerIn)
		}
	}
}

func (job *job) delayTicker(topic string) {
	for {
		startT := time.Now().Unix()
		err, ok := job.handel.DelayTicker(job.delayTickerLimit, topic)
		if err != nil {
			fmt.Println(err)
		}
		endT := time.Now().Unix()
		if endT-startT < 2 {
			if !ok {
				time.Sleep(job.delayTickerIn)
			}
		}
	}
}

func (job *job) retryTicker(topic string) {
	for {
		startT := time.Now().Unix()
		err, ok := job.handel.Retry(topic)
		if err != nil {
			helper.CheckErr(err)
		}
		endT := time.Now().Unix()
		if endT-startT < 2 {
			if !ok {
				time.Sleep(job.retryTickerIn)
			}
		}
	}
}
