package middleware

import (
	"gin-study/app/core/utils"
	"gin-study/app/http/filters"
	"gin-study/app/logic/enum"
	"gin-study/app/logic/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var filter filters.AuthToken
		err := ctx.ShouldBind(&filter)
		if err != nil {
			ctx.JSON(http.StatusOK, utils.ResponseInstance().Response(enum.StatusParamIsError))
			ctx.Abort()
			return
		}

		UserService := service.UserInstance()
		ok, authBean := UserService.GetAuth(filter.ClientId)
		if !ok {
			ctx.JSON(http.StatusOK, utils.ResponseInstance().Response(enum.StatusAuthForbidden))
			ctx.Abort()
			return
		}

		ctx.Set(enum.TagUserBean, authBean)
		ctx.Next()
	}
}
