package routes

import (
	"gin-study/app/http/controllers/api"
	"gin-study/app/http/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dispatch(r *gin.Engine) {
	webRoute(r)
	apiRoute(r)
}

func apiRoute(r *gin.Engine) {
	UserController := &api.UserController{}
	auth := r.Group("/")
	auth.Use(middleware.RequestLogger())
	{
		auth.POST("/user/register", UserController.Register)
		auth.POST("/user/login", UserController.Login)
		auth.Use(middleware.Auth())
		{
			auth.GET("/user", UserController.Index)
		}
	}
}

func webRoute(r *gin.Engine) {
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})
}
