package controllers

import (
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/bean"
	"gin-study/app/logic/enum"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (ctrl *Controller) Error404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, bean.ResponseBeanInstance().Response(enum.StatusNotFound))
	ctx.Abort()
}

func (ctrl *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, bean.ResponseBeanInstance().Response(enum.StatusOk, "ping"))
	ctx.Abort()
}

func (ctrl *Controller) AuthBean(ctx *gin.Context) *bean.AuthBean {
	authBeanInterface, _ := ctx.Get(enum.TagUserBean)
	authBean, ok := (authBeanInterface).(*bean.AuthBean)
	if ok && authBean != nil {
		return authBean
	}
	utils.NewException(enum.StatusAuthForbidden, helper.CreateMsg(enum.StatusAuthForbidden))
	return nil
}

func (ctrl *Controller) ResolveResponse(ctx *gin.Context) (bool, *bean.ResponseBean) {
	BeanInterface, _ := ctx.Get(enum.TagResponseBean)
	Bean, ok := (BeanInterface).(*bean.ResponseBean)
	return ok, Bean
}

func (ctrl *Controller) Response(ctx *gin.Context, code int, dataOrMsgParams ...interface{}) {
	ctx.Set(enum.TagResponseBean, bean.ResponseBeanInstance().Response(code, dataOrMsgParams...))
	ctx.Abort()
}
