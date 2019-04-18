package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gin-study/app/logic/enum"
	"github.com/apex/log"
	"github.com/google/uuid"
	"strings"
	"time"
)

func CreateUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func Md5(param string) string {
	data := []byte(param)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func CreatePwd(param string) string {
	hexStr := "RzjXn%gk6yxzAK-}7vNr.?TVs~yEHggP"
	return Md5(Md5(param) + hexStr)
}

func GetDateByFormat(timeParams ...time.Time) string {
	timeVars := time.Now()
	if len(timeParams) > 0 {
		timeVars = timeParams[0]
	}
	return timeVars.Format("2006-01-02 15:04:05")
}

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
		log.Error("checkErr:" + err.Error())
		if isThrow {
			panic(err)
		}
		return false
	}
	return true
}
