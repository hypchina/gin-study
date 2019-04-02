package helper

import (
	"errors"
	"fmt"
	"gin-study/app/logic/enum"
	"github.com/apex/log"
	"strconv"
	"strings"
)

func CreateErr(param interface{}) error {
	paramType := fmt.Sprintf("%T", param)
	var err string
	if paramType == "int" {
		err = enum.StatusText(param.(int))
	} else {
		err = param.(string)
	}
	return errors.New(err)
}

func CheckErr(err error, isThrows ...bool) bool {
	isThrow := false
	if len(isThrows) > 0 {
		isThrow = isThrows[0]
	}
	if err != nil {
		log.Info("checkErr:" + err.Error())
		if isThrow {
			panic(err)
		}
		return false
	}
	return true
}

func Byte2Str(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}
