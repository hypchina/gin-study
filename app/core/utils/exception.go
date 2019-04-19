package utils

import "gin-study/app/logic/bean"

func NewException(code int, msg string) {
	panic(&bean.ResponseBean{
		Code: code,
		Msg:  msg,
	})
}
