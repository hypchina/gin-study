package test

import (
	"gin-study/app/core/extend/env"
	"regexp"
	"strings"
	"testing"
)

func TestEnv(t *testing.T) {
	env.Init()
	println(env.All())
	//ok, x := isGroup("[fuck]")
	//fmt.Println(ok, x)
	/*x:=formatKey("fuck", "fuck_you","_")
	fmt.Println(x)*/
}

func isGroup(str string) (isGroup bool, groupName string) {
	regexp2 := regexp.MustCompile("^\\[([a-zA-Z]+[\\w_]+)]$")
	x := regexp2.FindStringSubmatch(str)
	if len(x) == 2 {
		return true, x[1]
	}
	return false, ""
}

func formatKey(groupName string, key string, dash string) string {
	n := strings.Index(key, groupName+dash)
	if n == 0 {
		return key
	}
	return groupName + dash + key
}
