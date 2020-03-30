package service

import (
	"encoding/json"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/bean"
	"gin-study/app/logic/dao"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/models"
	"github.com/gin-gonic/gin"
)

type LogSysRequestService struct {
	logSysRequestDao *dao.LogSysRequestDao
}

func LogSysRequestServiceInstance() *LogSysRequestService {
	return &LogSysRequestService{
		logSysRequestDao: dao.LogSysRequestDaoInstance(),
	}
}

func (service *LogSysRequestService) SyncInsert(ctx *gin.Context, bean *bean.ResponseBean) {
	respBody, _ := json.Marshal(bean.Data)

	startAt := helper.DateWithMs2TimestampWithMs(ctx.Writer.Header().Get(enum.TagRequestAt))
	endAt := helper.DateWithMs2TimestampWithMs(ctx.Writer.Header().Get(enum.TagResponseAt))

	LogSysRequestModel := models.LogSysRequest{
		Uid:              ctx.GetInt64(enum.TagRequestUid),
		RequestId:        ctx.Writer.Header().Get(enum.TagRequestId),
		RequestApi:       ctx.Request.URL.Path,
		RequestMethod:    ctx.Request.Method,
		RequestHeader:    "",
		RequestIp:        ctx.ClientIP(),
		RequestUri:       ctx.Request.RequestURI,
		RequestBody:      "",
		RequestAt:        ctx.Writer.Header().Get(enum.TagRequestAt),
		ResponseCode:     bean.Code,
		ResponseMsg:      bean.Msg,
		ResponseBody:     string(respBody),
		ResponseAt:       ctx.Writer.Header().Get(enum.TagResponseAt),
		ResponseDuration: endAt - startAt,
		CreatedAt:        helper.GetDateByFormatWithMs(),
	}
	go func() {
		bodyBytes, _ := json.Marshal(LogSysRequestModel)
		bodyStr := string(bodyBytes)
		err := utils.RedisClient().LPush("log:sync:app", bodyStr).Err()
		helper.CheckErr(err)
		//_, err := mq.NewJob(utils.RedisClient()).Publish(enum.TagJobTopicRequestLog, 1, LogSysRequestModel, enum.TagJobTagDefault)
		//helper.CheckErr(err)
	}()
}

func (service *LogSysRequestService) Insert(LogSysRequestModel models.LogSysRequest) (int64, error) {
	return service.logSysRequestDao.Insert(LogSysRequestModel)
}
