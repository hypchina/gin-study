package conf

import (
	"gin-study/app/core/extend/env"
	"strconv"
	"strings"
)

type common struct {
	Mode   string
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

type nim struct {
	AppKey    string
	AppSecret string
}

type logCollect struct {
	AppName []string
	LogType []string
}

type config struct {
	Common     common
	Mysql      mysql
	Redis      redis
	Nim        nim
	LogCollect logCollect
}

func Config() *config {
	return &config{
		Common: common{
			Mode:   env.Get("mode", "debug"),
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
			Db: func() int {
				dbIndex, _ := strconv.Atoi(env.Get("redis_db", "0"))
				return dbIndex
			}(),
		},
		Nim: nim{
			AppKey:    env.Get("nim_app_key"),
			AppSecret: env.Get("nim_app_secret"),
		},
		LogCollect: logCollect{
			AppName: func() []string {
				appNameStr := env.Get("log_collect_app_name", "test")
				return strings.Split(appNameStr, ",")
			}(),
			LogType: func() []string {
				logTypeStr := env.Get("log_collect_log_type", "api")
				return strings.Split(logTypeStr, ",")
			}(),
		},
	}
}

func (logCollect *logCollect) LogTypeIsExists(key string) bool {
	for _, item := range logCollect.LogType {
		if item == key {
			return true
		}
	}
	return false
}

func (logCollect *logCollect) AppNameIsExists(key string) bool {
	for _, item := range logCollect.AppName {
		if item == key {
			return true
		}
	}
	return false
}
