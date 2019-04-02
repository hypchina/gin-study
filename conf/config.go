package conf

import "gin-study/app/core/env"

type mysql struct {
	User     string
	Password string
	DSN      string
}

type configContext struct {
	Mysql mysql
}

func Config() *configContext {
	return &configContext{
		Mysql: mysql{
			User:     env.Get("mysql_user", ""),
			Password: env.Get("mysql_password", ""),
			DSN:      env.Get("mysql_dsn", ""),
		},
	}
}
