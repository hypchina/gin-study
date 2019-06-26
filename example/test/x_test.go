package test

import (
	"fmt"
	"gin-study/app/core/helper"
	"testing"
)

const (
	StartDate       = "2006-01-02 15:04:05"
	StartDateWithMs = "2006-01-02 15:04:05.000"
)

func TestGetDate(t *testing.T) {
	/*StartDate := "2006-01-02 15:04:05.000"
	timeVars := time.Now()
	x := timeVars.Format(StartDate)
	fmt.Println(x)
	*/
	x := helper.DateWithMs2TimestampWithMs("2019-06-25 13:00:51.484")
	fmt.Println(x)
}
