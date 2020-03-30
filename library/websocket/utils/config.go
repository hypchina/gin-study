package utils

type WebSocketServerConfig struct {
	Addr      string
	AppId     string
	AppSecret string
	ExpireIn  int
}

type WebSocketClientConfig struct {
	Addr      string
	AppId     string
	AppSecret string
	ClientId  string
	ExpireIn  int
}
