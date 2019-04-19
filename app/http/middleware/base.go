package middleware

import (
	"gin-study/app/core/utils"
	"gin-study/app/http/controllers"
	"gin-study/app/http/filters"
	"gin-study/app/logic/bean"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/service"
	"github.com/gin-gonic/gin"
	"log"
)

func Request() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var controller = &controllers.Controller{}
		utils.Try(func() {
			ctx.Next()
		}).Catch(&bean.ResponseBean{}, func(i interface{}) {
			log.Println("Exception:ResponseBean")
			ResponseBean, _ := (i).(*bean.ResponseBean)
			ctx.Set(enum.TagResponseBean, bean.ResponseBeanInstance().Response(ResponseBean.Code, ResponseBean.Msg))
		}).Catch("", func(i interface{}) {
			log.Println("Exception:String")
			ResponseBean := bean.ResponseBeanInstance().Response(enum.StatusUnknownError)
			ctx.Set(enum.TagResponseBean, bean.ResponseBeanInstance().Response(ResponseBean.Code))
		}).Finally(func() {
			log.Println("Finally")
			ok, ResponseBean := controller.ResolveResponse(ctx)
			if !ok {
				ResponseBean = bean.ResponseBeanInstance().Response(enum.StatusUnknownResponse)
			}
			ctx.JSON(enum.StatusOk, ResponseBean)
			ctx.Abort()
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
			controller.Response(ctx, enum.StatusAuthForbidden)
			return
		}
		ctx.Set(enum.TagUserBean, authBean)
		ctx.Next()
	}
}
