package service

import (
	"gin-study/app/core/helper"
	"gin-study/app/http/filters"
	"gin-study/app/logic/dao"
	"gin-study/app/logic/models"
	"log"
	"time"
)

type BroadcastService struct {
	broadcastDao *dao.BroadcastDao
}

func BroadcastServiceInstance() *BroadcastService {
	return &BroadcastService{
		broadcastDao: dao.BroadcastDaoInstance(),
	}
}

func (service *BroadcastService) CreateAndNotify(filter *filters.BroadcastNotifyFilter) *models.LogUserBroadcast {
	LogUserBroadcast := &models.LogUserBroadcast{
		MsgId:     helper.CreateUUID(),
		FromUid:   "0",
		ToUid:     filter.ToUid,
		MsgType:   filter.MsgType,
		MsgBody:   filter.MsgBody,
		MsgLevel:  1,
		ReadState: 0,
		ExpireAt:  helper.GetDateByFormat(time.Now().Add(time.Second * 1800)),
		NotifyAt:  helper.GetDateByFormat(),
		ReadAt:    helper.GetDateByFormat(),
		CreatedAt: helper.GetDateByFormat(),
	}
	if ok := service.broadcastDao.Insert(LogUserBroadcast); ok {
		return LogUserBroadcast
	}
	return nil
}

func (service *BroadcastService) NotifyById(msgId string) {
	LogUserBroadcast, err := service.broadcastDao.GetByMsgId(msgId)
	log.Println(LogUserBroadcast, err)
}

func (service *BroadcastService) ReadById(msgId string, outUid string) (*models.LogUserBroadcast, bool) {
	if LogUserBroadcast, ok := service.broadcastDao.GetByMsgId(msgId); ok && LogUserBroadcast.ToUid == outUid {
		return LogUserBroadcast, true
	}
	return nil, false
}
