package enum

const (
	cacheNamespace    = "cache:"
	queueNamespace    = "queue:"
	tokenNamespace    = "token:"
	RedisCacheTestStr = "test:str"
)

func RedisCacheKey(key string) string {
	return cacheNamespace + key
}

func RedisTokenKey(key string) string {
	return tokenNamespace + key
}

func RedisQueueKey(key string) string {
	return queueNamespace + key
}
