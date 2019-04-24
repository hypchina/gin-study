package enum

import "strconv"

const (
	nsBroadcast = "broadcast:"
	nsToken     = "token:"
	nsUser      = "user:"
)

func RedisUserKey(id int64) string {
	return nsUser + strconv.FormatInt(id, 12)
}

func RedisBroadcast(key string) string {
	return nsBroadcast + key
}

func RedisTokenKey(key string) string {
	return nsToken + key
}
