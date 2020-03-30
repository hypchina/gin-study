package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"sort"
	"strings"
)

func httpBuildQuery(params map[string]interface{}, singExcept bool) string {
	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var pairs []string
	for _, k := range keys {
		if k == "sign" && singExcept {
			continue
		}
		val := fmt.Sprintf("%v", params[k])
		pairs = append(pairs, k+"="+val)
	}
	return strings.Join(pairs[:], "&")
}

func struct2MapX(obj interface{}) map[string]interface{} {
	x := map[string]interface{}{}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return x
	}
	err = json.Unmarshal(jsonBytes, &x)
	if err != nil {
		return x
	}
	return x
}

func md5X(param string) string {
	data := []byte(param)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func CreateUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
