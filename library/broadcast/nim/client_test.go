package nim

import (
	"fmt"
	"gin-study/app/core/extend/env"
	"gin-study/library/broadcast/nim/request"
	"testing"
)

func TestSendMsg(t *testing.T) {
	client := NewClient(env.Get("nim_app_key"), env.Get("nim_app_secret"))
	requestX := request.NewSendMsgRequest("2", "2", "Hello")
	respStr, err := client.Do(requestX)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(respStr)
}
