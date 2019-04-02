package main

import (
	"gin-study/app/core/env"
	"gin-study/app/core/utils"
	"gin-study/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	env.Load()
	utils.DbInit()
	r := gin.Default()
	r.Use(gin.Recovery())
	routes.Dispatch(r)
	_ = r.Run(":" + env.Get("listen", "8080"))
}
