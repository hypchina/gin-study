package service

import (
	"fmt"
	"gin-study/app/core/utils"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"gin-study/conf"
	"gin-study/library/broadcast/nim"
	"gin-study/library/broadcast/nim/request"
	"gin-study/library/mq"
)

type imService struct {
}

func NewImService() *imService {
	return &imService{}
}

func (service *imService) Push(broadcast *models.LogUserBroadcast) {
	_, err := mq.NewJob(utils.RedisClient()).Publish(enum.TagJobTopicBroadcast, 0, broadcast, enum.TagJobTagDefault)
	if err != nil {
		client := nim.NewClient(conf.Config().Nim.AppKey, conf.Config().Nim.AppSecret)
		requestX := request.NewSendMsgRequest(broadcast.FromUid, broadcast.ToUid, broadcast.MsgBody)
		_, err = client.Do(requestX)
	}
	fmt.Println(err)
}
