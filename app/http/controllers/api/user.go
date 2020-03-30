package api

import (
	"C"
	"gin-study/app/http/controllers"
	"gin-study/app/http/filters"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	*controllers.Controller
}

//用户个人信息
func (ctrl *UserController) Index(ctx *gin.Context) {
	ctrl.Response(ctx, enum.StatusOk, ctrl.AuthBean(ctx).UserBean())
}

//用户注册
func (ctrl *UserController) Register(ctx *gin.Context) {

	var filter filters.UserRegister
	err := ctx.ShouldBind(&filter)
	if err != nil {
		ctrl.Response(ctx, enum.StatusParamIsError, err.Error())
		return
	}

	UserService := service.UserInstance()
	err = UserService.Create(filter)
	if err != nil {
		ctrl.Response(ctx, enum.StatusDataOpError, err.Error())
		return
	}

	ctrl.Response(ctx, enum.StatusOk)
}

func (ctrl *UserController) Login(ctx *gin.Context) {

	var filter filters.UserLogin
	err := ctrl.Filter(ctx, &filter)
	//err := ctx.ShouldBind(&filter)
	if err != nil {
		ctrl.Response(ctx, enum.StatusParamIsError, err.Error())
		return
	}

	UserService := service.UserInstance()
	authBean, err := UserService.CreateAuth(filter)
	if err != nil {
		ctrl.Response(ctx, enum.StatusDataOpError, err.Error())
		return
	}

	ctrl.Response(ctx, enum.StatusOk, authBean)
}