package api

import (
	"gin-study/app/http/controllers"
	"gin-study/app/http/filters"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/service"
	"github.com/gin-gonic/gin"
)

type Broadcast struct {
	*controllers.Controller
}

func (ctrl *Broadcast) Notify(ctx *gin.Context) {
	var filter filters.BroadcastNotifyFilter
	if err := ctx.ShouldBind(&filter); err != nil {
		ctrl.Response(ctx, enum.StatusParamIsError, err.Error())
		return
	}
	broadcast := service.BroadcastServiceInstance().CreateAndNotify(&filter)
	if broadcast != nil {
		ctrl.Response(ctx, enum.StatusOk, broadcast)
		return
	}
	ctrl.Response(ctx, enum.StatusDataOpError)
}

func (ctrl *Broadcast) Read(ctx *gin.Context) {

	msgId := ctx.Query("msg_id")
	if msgId == "" {
		ctrl.Response(ctx, enum.StatusParamIsError, "miss msg_id")
		return
	}

	if broadcast, ok := service.BroadcastServiceInstance().ReadById(msgId, ctrl.AuthBean(ctx).UserBean().UserEntity.OutUid); ok {
		ctrl.Response(ctx, enum.StatusOk, broadcast)
		return
	}

	ctrl.Response(ctx, enum.StatusNotFound)
}
