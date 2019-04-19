package enum

import "strconv"

const (
	cacheNamespace = "cache:"
	queueNamespace = "queue:"
	tokenNamespace = "token:"
	userNamespace  = "user:"
)

func RedisUserKey(id int64) string {
	return userNamespace + strconv.FormatInt(id, 12)
}

func RedisCacheKey(key string) string {
	return cacheNamespace + key
}

func RedisTokenKey(key string) string {
	return tokenNamespace + key
}

func RedisQueueKey(key string) string {
	return queueNamespace + key
}
