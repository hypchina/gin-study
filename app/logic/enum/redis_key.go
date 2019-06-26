package enum

import "strconv"

const (
	nsBroadcastDetail = "broadcast:detail"
	nsBroadcastQueue  = "broadcast:queue"
	nsLogCollectQueue = "log_collect:queue:"
	nsToken           = "token:"
	nsUser            = "user:"
)

func RedisUserKey(id int64) string {
	return nsUser + strconv.FormatInt(id, 12)
}

func RedisBroadcastDetail(key string) string {
	return nsBroadcastDetail + key
}

func RedisBroadcastQueue(key string) string {
	return nsBroadcastQueue + key
}

func RedisTokenKey(key string) string {
	return nsToken + key
}

func RedisLogCollectQueue(key string) string {
	return nsLogCollectQueue + key
}
