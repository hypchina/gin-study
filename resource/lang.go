package resource

import (
	"gin-study/conf"
)

var langMap = map[string]map[string]string{
	"cn": {
		"user_not_exists": "用户不存在",
		"user_exists":     "用户已注册",
		"email_exists":    "该邮箱已注册",
	},
}

func Trans(key string) string {
	item, ok := langMap[conf.Config().Common.Lang]
	if ok {
		val, ok2 := item[key]
		if ok2 {
			return val
		}
	}
	return key
}
