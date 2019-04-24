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

const (
	StartDate = "2006-01-02 15:04:05"
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
	return timeVars.Format(StartDate)
}

func GetTimeByDate(dateString string) (time.Time, error) {
	_time, err := time.Parse(StartDate, dateString)
	return _time, err
}

func GetDefautlDate() string {
	return "0000-00-00 00:00:00"
}

func CreateMsg(param interface{}) string {
	paramType := fmt.Sprintf("%T", param)
	var msg string
	if paramType == "int" {
		msg = enum.StatusText(param.(int))
	} else {
		msg = param.(string)
	}
	return msg
}

func CreateErr(param interface{}) error {
	err := CreateMsg(param)
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
