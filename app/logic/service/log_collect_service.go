package service

import (
	"encoding/json"
	"errors"
	"gin-study/app/http/filters"
	"gin-study/app/logic/dao"
	"gin-study/app/logic/models"
	"gin-study/conf"
	"log"
)

type logCollectService struct {
	LogCollectDao *dao.LogCollectDao
}

func LogCollectServiceInstance() *logCollectService {
	return &logCollectService{
		LogCollectDao: dao.LogCollectDaoInstance(),
	}
}

//异步发送消息失败之后转为同步
func (service *logCollectService) CreateWithAsyncAndFailSync(logCollectFilter *filters.LogCollectFilter) error {

	logCollectConfig := conf.Config().LogCollect
	if !logCollectConfig.AppNameIsExists(logCollectFilter.AppName) {
		return errors.New("无效的app_name")
	}

	if !logCollectConfig.LogTypeIsExists(logCollectFilter.LogType) {
		return errors.New("无效的log_type")
	}

	go func() {
		logCollectModel := models.LogCollect{
			AppName:  logCollectFilter.AppName,
			TraceId:  logCollectFilter.TraceId,
			LogType:  logCollectFilter.LogType,
			LogLevel: logCollectFilter.LogLevel,
			LogBody:  logCollectFilter.LogBody,
			EventAt:  logCollectFilter.EventAt,
		}
		err := service.LogCollectDao.Insert(logCollectModel)
		if err != nil {
			logCollectJson, _ := json.Marshal(logCollectModel)
			log.Println(logCollectJson)
		}
	}()

	return nil
}
