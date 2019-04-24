package main

import (
	"gin-study/app/core/extend/env"
	"gin-study/app/core/utils"
	"gin-study/conf"
	"gin-study/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	gin.SetMode(conf.Config().Common.Mode)
	r := gin.Default()
	routes.Dispatch(r)
	_ = r.Run(conf.Config().Common.Addr)
}
