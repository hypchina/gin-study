package utils

import (
	"fmt"
	"gin-study/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
	"time"
)

var ORM *xorm.Engine

func DbInit() { // 使用init来自动连接数据库，并创建ORM实例
	var err error
	source := fmt.Sprintf("%s:%s@%s", conf.Config().Mysql.User, conf.Config().Mysql.Password, conf.Config().Mysql.DSN)
	ORM, err = xorm.NewEngine("mysql", source)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = ORM.Ping() // 测试能操作数据库
	if err != nil {
		log.Fatalln(err)
		return
	}
	ORM.SetMaxIdleConns(5) //设置连接池的空闲数大小
	ORM.SetMaxOpenConns(10) //设置最大打开连接数
	ORM.ShowSQL(true)      // 测试环境，显示每次执行的sql语句长什么样子
	timer := time.NewTicker(time.Minute * 1)
	go func(x *xorm.Engine) {
		for t := range timer.C {
			log.Println("mysql:ticker:" + t.String())
			err = x.Ping()
			if err != nil {
				log.Fatalf("db connect error: %#v\n", err.Error())
			}
		}
	}(ORM)
}
