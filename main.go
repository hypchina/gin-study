package main

import (
	"gin-study/app/core/extend/env"
	"gin-study/app/core/utils"
	"gin-study/conf"
	"gin-study/library/v9_validator"
	"gin-study/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	env.Init()
	utils.DbInit()
	utils.RedisInit()
	binding.Validator = new(v9_validator.DefaultValidator)
	gin.SetMode(conf.Config().Common.Mode)
	r := gin.Default()
	routes.Dispatch(r)
	_ = r.Run(conf.Config().Common.Addr)
}
