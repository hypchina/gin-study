package middleware

import (
	"gin-study/app/core/utils"
	"gin-study/app/http/controllers"
	"gin-study/app/http/filters"
	"gin-study/app/logic/bean"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var controller = &controllers.Controller{}
		utils.Try(func() {
			ctx.Next()
		}).Catch(&bean.ResponseBean{}, func(i interface{}) {
			ResponseBean, ok := (i).(*bean.ResponseBean)
			if ok {
				controller.Response(ctx, ResponseBean.Code, ResponseBean.Msg)
			}
		}).Finally(func() {
			ok, ResponseBean := controller.ResolveResponse(ctx)
			if ok {
				ctx.JSON(http.StatusOK, ResponseBean)
			} else {
				ctx.JSON(http.StatusInternalServerError, bean.ResponseBeanInstance().Response(enum.StatusUnknownResponse))
			}
		})
	}
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var controller = &controllers.Controller{}
		var filter filters.AuthToken
		err := ctx.ShouldBind(&filter)
		if err != nil {
			controller.Response(ctx, enum.StatusParamIsError)
			return
		}
		UserService := service.UserInstance()
		ok, authBean := UserService.GetAuth(filter.ClientId)
		if !ok {
			//controller.Response(ctx, enum.StatusAuthForbidden)
			return
		}
		ctx.Set(enum.TagUserBean, authBean)
		ctx.Next()
	}
}
