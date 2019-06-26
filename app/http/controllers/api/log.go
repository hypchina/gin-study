package api

import (
	"gin-study/app/core/helper"
	"gin-study/app/http/controllers"
	"gin-study/app/http/filters"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/service"
	"github.com/gin-gonic/gin"
)

type LogController struct {
	*controllers.Controller
}

func (ctrl *LogController) Create(ctx *gin.Context) {
	var filter filters.LogCollectFilter
	if err := ctx.ShouldBind(&filter); err != nil {
		ctrl.Response(ctx, enum.StatusParamIsError, err.Error())
		return
	}
	err := service.LogCollectServiceInstance().CreateWithAsyncAndFailSync(&filter)

	if err != nil {
		helper.CheckErr(err)
		ctrl.Response(ctx, enum.StatusDataOpError, err.Error())
		return
	}

	ctrl.Response(ctx, enum.StatusOk)
}
