package bean

import (
	"fmt"
	"gin-study/app/logic/enum"
)

type ResponseBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseBeanInstance() *ResponseBean {
	return &ResponseBean{
		Code: enum.StatusUnknownError,
		Msg:  enum.StatusText(enum.StatusUnknownError),
		Data: map[string]interface{}{},
	}
}

func (response *ResponseBean) Response(code int, params ...interface{}) *ResponseBean {
	response.Code = code
	i := 0
	isSetMsg := false
	if paramSize := len(params); paramSize > 0 {
		for index := range params {
			param := params[index]
			paramType := fmt.Sprintf("%T", param)
			if paramType == "string" {
				response.Msg = param.(string)
				isSetMsg = true
			} else {
				response.Data = param
			}
			if i > 1 {
				break
			}
			i++
		}
	}
	if !isSetMsg {
		response.Msg = enum.StatusText(code)
	}
	return response
}
