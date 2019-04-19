package routes

import (
	"gin-study/app/http/controllers"
	"gin-study/app/http/controllers/api"
	"gin-study/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func Dispatch(r *gin.Engine) {
	apiRoute(r)
}

func apiRoute(r *gin.Engine) {
	BaseController := &controllers.Controller{}
	UserController := &api.UserController{}
	r.NoRoute(BaseController.Error404)
	r.NoMethod(BaseController.Error404)
	auth := r.Group("/")
	auth.Use(middleware.RequestLogger())
	{
		auth.POST("/", BaseController.Ping)
		auth.POST("/user/register", UserController.Register)
		auth.POST("/user/login", UserController.Login)
		auth.Use(middleware.Auth())
		{
			auth.GET("/user", UserController.Index)
		}
	}
}
