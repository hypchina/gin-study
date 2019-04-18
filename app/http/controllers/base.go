package controllers

import (
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/bean"
	"gin-study/app/logic/enum"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
}

func (ctrl *Controller) Error404() {
	log.Println("404")
}

func (ctrl *Controller) AuthBean(ctx *gin.Context) *bean.AuthBean {
	authBeanInterface, _ := ctx.Get(enum.TagUserBean)
	authBean, ok := (authBeanInterface).(*bean.AuthBean)
	if !ok {
		panic(helper.CreateErr(enum.StatusAuthForbidden))
	}
	return authBean
}

func (ctrl *Controller) Response(ctx *gin.Context, code int, dataOrMsgParams ...interface{}) {
	ctx.JSON(http.StatusOK, utils.ResponseInstance().Response(code, dataOrMsgParams...))
	ctx.Abort()
}
