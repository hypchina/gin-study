package conf

import "gin-study/app/core/env"

type mysql struct {
	User     string
	Password string
	DSN      string
}

type common struct {
	Env    string
	Lang   string
	Listen string
	Addr   string
}

type config struct {
	Common common
	Mysql  mysql
}

func Config() *config {
	return &config{
		Common: common{
			Env:    env.Get("env", "local"),
			Listen: env.Get("listen", "8080"),
			Addr:   ":" + env.Get("listen", "8080"),
			Lang:   env.Get("lang", "cn"),
		},
		Mysql: mysql{
			User:     env.Get("mysql_user", ""),
			Password: env.Get("mysql_password", ""),
			DSN:      env.Get("mysql_dsn", ""),
		},
	}
}
