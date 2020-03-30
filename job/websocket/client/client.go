package main

import (
	"gin-study/app/core/extend/env"
	"gin-study/library/websocket/client"
	"gin-study/library/websocket/utils"
	"log"
)

func main() {
	env.Init()
	config := &utils.WebSocketClientConfig{
		Addr:      env.Get("websocket_client_addr"),
		AppId:     env.Get("websocket_app_id"),
		AppSecret: env.Get("websocket_app_secret"),
		ClientId:  env.Get("websocket_client_id"),
	}
	websocketClient := client.NewWebsocketClient(config)
	websocketClient.OnConnected = func(ok string) {
		log.Println(ok)
		websocketClient.Broadcast(config.ClientId, 1, "Hello")
	}
	websocketClient.OnBroadcast = func(response utils.BroadcastResponse) {
		log.Println(response.Type,response.From,response.Body)
	}
	websocketClient.Run()
}
