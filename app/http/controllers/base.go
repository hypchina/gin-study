package controllers

import (
	"errors"
	"fmt"
	"gin-study/app/core/helper"
	"gin-study/app/core/utils"
	"gin-study/app/logic/bean"
	"gin-study/app/logic/enum"
	"gin-study/vendor/gopkg.in/go-playground/validator.v8"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type Controller struct {
}

func (ctrl *Controller) Filter(ctx *gin.Context, filterStruct interface{}) error {

	err := ctx.ShouldBind(filterStruct)
	if err != nil {
		errorMap, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		for _, v := range errorMap {
			return errors.New(fmt.Sprintf("参数%s 数据类型%s 验证规则%s", v.Field(), v.Kind(), v.Tag()+":"+v.Param()))
		}
	}
	return nil
}

func (ctrl *Controller) Error404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, bean.ResponseBeanInstance().Response(enum.StatusNotFound))
	ctx.Abort()
}

func (ctrl *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, bean.ResponseBeanInstance().Response(enum.StatusOk, "pong"))
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
