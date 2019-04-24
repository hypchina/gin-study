package utils

import (
	"fmt"
	"gin-study/conf"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
	"time"
)

var orm *xorm.Engine
var ormInit bool

func DbInit() { // 使用init来自动连接数据库，并创建ORM实例
	var err error
	source := fmt.Sprintf("%s:%s@%s", conf.Config().Mysql.User, conf.Config().Mysql.Password, conf.Config().Mysql.DSN)
	orm, err = xorm.NewEngine("mysql", source)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = orm.Ping() // 测试能操作数据库
	if err != nil {
		log.Fatalln(err)
		return
	}
	orm.SetMaxIdleConns(10)  //设置连接池的空闲数大小
	orm.SetMaxOpenConns(200) //设置最大打开连接数
	if conf.Config().Common.Mode == gin.DebugMode {
		orm.ShowSQL(true) // 测试环境，显示每次执行的sql语句长什么样子
	}
	timer := time.NewTicker(time.Minute * 1)
	ormInit = true
	go func(x *xorm.Engine) {
		for range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("db connect error: %#v\n", err.Error())
			}
		}
	}(orm)
}

func ORM() *xorm.Engine {
	if !ormInit {
		panic("请初始化数据库")
	}
	return orm
}
