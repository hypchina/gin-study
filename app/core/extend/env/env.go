package env

import (
	"gin-study/app/core/helper"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	flag       = false
	envContext = make(map[string]string)
)

const (
	envFn    = ".env"
	envLf    = "\n"
	envSq    = "="
	envNull  = "null"
	envGroup = "_"
)

func Init() {
	if !flag {
		byteSet, err := ioutil.ReadFile(envFn)
		helper.CheckErr(err)
		contentSet := strings.Split(string(byteSet), envLf)
		prefix := ""
		for index := range contentSet {
			line := contentSet[index]
			if line == "" {
				continue
			}

			if ok, groupName := isGroup(line); ok {
				prefix = groupName
			}

			lineSet := strings.Split(line, envSq)
			if len(lineSet) < 2 {
				continue
			}
			key := getKey(prefix, lineSet[0], envGroup)
			envContext[key] = strings.Join(lineSet[1:], envSq)
		}
		flag = true
	}
}

func All() map[string]string {
	return envContext
}

func Get(key string, defaults ...string) string {
	if val, ok := envContext[key]; ok {
		return val
	}
	defaultVal := envNull
	if len(defaults) > 0 {
		defaultVal = defaults[0]
	}
	return defaultVal
}

func Set(key string, val string) {
	envContext[key] = val
}

func isGroup(str string) (isGroup bool, groupName string) {
	regexp2 := regexp.MustCompile("^\\[([a-zA-Z]+[\\w_]+)]$")
	x := regexp2.FindStringSubmatch(str)
	if len(x) == 2 {
		return true, x[1]
	}
	return false, ""
}

func getKey(groupName string, key string, dash string) string {
	n := strings.Index(key, groupName+dash)
	if n == 0 {
		return key
	}
	return groupName + dash + key
}
