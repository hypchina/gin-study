package api

import (
	"C"
	"gin-study/app/http/controllers"
	"gin-study/app/http/filters"
	"gin-study/app/logic/dao"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	*controllers.Controller
}

func (ctrl *UserController) Index(ctx *gin.Context) {
	if exists, user := dao.UserInstance().GetByEmail(ctx.DefaultQuery("email", "-1")); exists {
		ctrl.Response(ctx, enum.StatusOk, user)
		return
	}
	ctrl.Response(ctx, enum.StatusDataIsNotExists)
	return
}

func (ctrl *UserController) Create(ctx *gin.Context) {

	var userFilter filters.UserFilter
	err := ctx.ShouldBind(&userFilter)

	if err != nil {
		ctrl.Response(ctx, enum.StatusParamIsError, err.Error())
		return
	}

	userService := service.UserInstance()
	user, err := userService.Create(userFilter)

	if err != nil {
		ctrl.Response(ctx, enum.StatusDataOpError, err.Error())
		return
	}

	ctrl.Response(ctx, enum.StatusOk, user)
}
