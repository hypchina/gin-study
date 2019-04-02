package utils

import (
	"fmt"
	"gin-study/app/logic/enum"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseInstance() *Response {
	return &Response{
		Code: enum.StatusUnknownError,
		Msg:  enum.StatusText(enum.StatusUnknownError),
		Data: map[string]interface{}{},
	}
}

func (response *Response) Response(code int, params ...interface{}) *Response {
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
