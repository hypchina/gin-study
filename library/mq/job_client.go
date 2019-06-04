package mq

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type job struct {
	defaultExpireIn     time.Duration
	readAfterExpireIn   time.Duration
	delayTickerIn       time.Duration
	subscribeTickerIn   time.Duration
	subscribeQueueLimit int
	handel              *redisHandel
}

func NewJob(redisClient *redis.Client) *job {
	return &job{
		subscribeQueueLimit: 200,
		defaultExpireIn:     time.Second * 86400 * 30,
		readAfterExpireIn:   time.Second * 60,
		subscribeTickerIn:   time.Second * 1,
		delayTickerIn:       time.Second * 2,
		handel:              newRedisHandle(redisClient),
	}
}

func (job *job) SetSubscribeQueueLimit(subscribeQueueLimit int) *job {
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

func (job *job) Subscribe(topic string, callback func(JobStruct, error)) {
	for {
		startT := time.Now().Unix()
		i := 0
		for i = 0; i < job.subscribeQueueLimit; i++ {
			jobStruct, err := job.handel.Read(topic, job.readAfterExpireIn)
			if err != nil || jobStruct == nil {
				break
			}
			callback(*jobStruct, err)
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

func (job *job) DelayTicker() {
	timerX := time.NewTicker(job.delayTickerIn)
	for range timerX.C {
		err := job.handel.DelayTicker()
		if err != nil {
			fmt.Println(err)
		}
	}
}
