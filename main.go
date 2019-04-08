package main

import (
	"gin-study/app/core/env"
	"gin-study/app/core/utils"
	"gin-study/conf"
	"gin-study/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	env.Load()
	utils.DbInit()
	r := gin.Default()
	routes.Dispatch(r)
	_ = r.Run(conf.Config().Common.Addr)
}
