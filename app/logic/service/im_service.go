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

//异步发送消息失败之后转为同步
func (service *imService) SendWithAsyncAndFailSync(broadcast *models.LogUserBroadcast) {
	var respStr string
	_, err := mq.NewJob(utils.RedisClient()).Publish(enum.TagJobTopicBroadcast, 0, broadcast, enum.TagJobTagDefault)
	if err != nil {
		respStr, err = service.SendWithSync(broadcast)
	}
	fmt.Println(respStr, err)
}

//同步发送
func (service *imService) SendWithSync(broadcast *models.LogUserBroadcast) (respStr string, err error) {
	client := nim.NewClient(conf.Config().Nim.AppKey, conf.Config().Nim.AppSecret)
	requestX := request.NewSendMsgRequest(broadcast.FromUid, broadcast.ToUid, broadcast.MsgBody)
	respStr, err = client.Do(requestX)
	return
}
