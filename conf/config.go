package conf

import (
	"gin-study/app/core/extend/env"
	"strconv"
)

type common struct {
	Env    string
	Lang   string
	Listen string
	Addr   string
}

type mysql struct {
	User     string
	Password string
	DSN      string
}

type redis struct {
	Addr     string
	Password string
	Db       int
}

type config struct {
	Common common
	Mysql  mysql
	Redis  redis
}

func Config() *config {
	db, _ := strconv.Atoi(env.Get("redis_db", "0"))
	return &config{
		Common: common{
			Env:    env.Get("env", "local"),
			Listen: env.Get("listen", "8080"),
			Addr:   ":" + env.Get("listen", "8080"),
			Lang:   env.Get("lang", "cn"),
		},
		Mysql: mysql{
			User:     env.Get("mysql_user"),
			Password: env.Get("mysql_password"),
			DSN:      env.Get("mysql_dsn"),
		},
		Redis: redis{
			Addr:     env.Get("redis_addr"),
			Password: env.Get("redis_password"),
			Db:       db,
		},
	}
}
