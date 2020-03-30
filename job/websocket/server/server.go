package main

import (
	"gin-study/app/core/extend/env"
	"gin-study/library/websocket/server"
	"gin-study/library/websocket/utils"
)

func main() {
	env.Init()
	config := &utils.WebSocketServerConfig{
		AppId:     env.Get("websocket_app_id"),
		AppSecret: env.Get("websocket_app_secret"),
		Addr:      env.Get("websocket_server_addr"),
	}
	websocketServer := server.NewWebSocketServer(config)
	websocketServer.Run()
}
