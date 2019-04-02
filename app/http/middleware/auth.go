package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Query("uid") == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请登录",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
