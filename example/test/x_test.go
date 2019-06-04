package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	StartDate := "2006-01-02 15:04:05.000"
	timeVars := time.Now()
	x := timeVars.Format(StartDate)
	fmt.Println(x)
}
