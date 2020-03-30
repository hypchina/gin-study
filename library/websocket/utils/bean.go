package utils

type ClientInfo struct {
	ClientId         string `json:"client_id"`
	Addr             string `json:"addr"`
	ConnectAt        string `json:"connect_at"`
	PongAt           string `json:"pong_at"`
	BroadcastCounter int64  `json:"broadcast_counter"`
}
