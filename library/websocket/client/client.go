package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-study/library/websocket/utils"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

type websocketSocketClient struct {
	done        chan os.Signal
	socket      *Socket
	config      *utils.WebSocketClientConfig
	OnConnected func(clientId string)
	OnBroadcast func(response utils.BroadcastResponse)
	OnClose     func()
}

func NewWebsocketClient(config *utils.WebSocketClientConfig) *websocketSocketClient {
	log.Println(config)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	client := &websocketSocketClient{
		config: config,
		done:   done,
		socket: New(getAuthUrl(config)),
	}
	return client
}

func (client *websocketSocketClient) Run() {

	client.socket.OnConnected = func(socket Socket) {
		if client.OnConnected != nil {
			client.socket = &socket
			client.OnConnected(fmt.Sprintf("client_id %s connect success", client.config.ClientId))
		}
	}
	client.socket.OnTextMessage = func(message string, socket Socket) {
		if client.OnBroadcast != nil {
			var BroadcastResponse utils.BroadcastResponse
			err := json.Unmarshal([]byte(message), &BroadcastResponse)
			if err != nil {
				fmt.Println("response format is err", err.Error())
				return
			}
			client.OnBroadcast(BroadcastResponse)
		}
	}
	client.socket.OnDisconnected = func(err error, socket Socket) {
		if err.Error() == "quit" {
			log.Println("客户端退出")
			return
		}
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			e, _ := err.(*websocket.CloseError)
			log.Println(fmt.Sprintf("服务端退出,error:%v", e.Text))
			return
		}
		log.Println(fmt.Sprintf("意外错误5秒后启动重试机制,error:%v", err.Error()))
		time.AfterFunc(time.Second*5, func() {
			socket.Url = getAuthUrl(client.config)
			log.Println(fmt.Sprintf("启动重试机制"))
			socket.Connect()
		})
	}

	client.socket.OnConnectError = func(err error, socket Socket) {
		socket.Close(err)
	}

	client.socket.OnReadError = func(err error, socket Socket) {
		socket.Close(err)
	}

	client.socket.Connect()
	for {
		select {
		case <-client.done:
			client.socket.Close(errors.New("quit"))
			return
		}
	}
}

func (client *websocketSocketClient) Broadcast(to string, msgType int, msgBody string) {
	broadcastRequest := utils.BroadcastRequest{
		To:   to,
		Type: msgType,
		Body: msgBody,
	}
	client.socket.SendText(broadcastRequest.ToJson())
}

func getAuthUrl(config *utils.WebSocketClientConfig) string {
	signTool := utils.NewSign(config.AppId, config.AppSecret)
	link := config.Addr + "?" + signTool.CreateAuthLink(config.ClientId)
	log.Println("auth:" + link)
	return link
}
