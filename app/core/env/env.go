package env

import (
	"gin-study/app/core/helper"
	"io/ioutil"
	"strings"
)

var (
	flag       = false
	envContext = make(map[string]string)
)

const (
	envFn   = ".env"
	envLf   = "\n"
	envSq   = "="
	envNull = "null"
)

func Load() {
	if !flag {
		byteSet, err := ioutil.ReadFile(envFn)
		helper.CheckErr(err)
		contentSet := strings.Split(string(byteSet), envLf)
		for index := range contentSet {
			line := contentSet[index]
			if line == "" {
				continue
			}
			lineSet := strings.Split(line, envSq)
			if len(lineSet) < 2 {
				continue
			}
			envContext[lineSet[0]] = strings.Join(lineSet[1:], envSq)
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
