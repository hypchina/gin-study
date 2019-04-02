package controllers

import (
	"gin-study/app/core/utils"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (ctrl *Controller) Error404() {
	log.Info("404")
}

func (ctrl *Controller) Response(ctx *gin.Context, code int, dataOrMsgParams ...interface{}) {
	ctx.JSON(http.StatusOK, utils.ResponseInstance().Response(code, dataOrMsgParams...))
	ctx.Abort()
}
