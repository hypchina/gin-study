package server

import (
	"bytes"
	_ "bytes"
	"encoding/json"
	"fmt"
	"gin-study/app/core/helper"
	"gin-study/library/websocket/utils"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type client struct {
	clientId         string
	handle           *handle
	broadcast        chan broadcast
	conn             *websocket.Conn
	connectAt        string
	pongAt           string
	broadcastCounter int64
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func newClient(clientId string, handle *handle, conn *websocket.Conn) *client {
	return &client{
		clientId:         clientId,
		handle:           handle,
		broadcast:        make(chan broadcast),
		conn:             conn,
		connectAt:        helper.GetDateByFormatWithMs(),
		broadcastCounter: 0,
	}
}

func (c *client) read() {

	defer func() {
		c.handle.unregister <- c
	}()

	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println(fmt.Sprintf("客户端退出 ip:%v", c.conn.RemoteAddr()))
				return
			}
			log.Println(fmt.Sprintf("意外错误,error:%v", err.Error()))
			break
		}

		var BroadcastRequest utils.BroadcastRequest
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		err = json.Unmarshal(message, &BroadcastRequest)
		if err != nil {
			log.Println(fmt.Sprintf("无效的请求格式,error:%v", err.Error()))
			continue
		}

		c.handle.broadcast <- broadcast{
			fromId:    c.clientId,
			fromIp:    c.conn.RemoteAddr().String(),
			toId:      BroadcastRequest.To,
			msgType:   BroadcastRequest.Type,
			body:      BroadcastRequest.Body,
			receiveAt: helper.GetDateByFormatWithMs(),
		}
	}
}

func (c *client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case broadcast := <-c.broadcast:
			if len(broadcast.body) == 0 {
				return
			}
			c.broadcastCounter++
			c.handle.broadcastCounter++
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			jsonStrByte, err := json.Marshal(utils.BroadcastResponse{
				RequestId:   utils.CreateUUID(),
				From:        broadcast.fromId,
				To:          broadcast.toId,
				Type:        broadcast.msgType,
				Body:        broadcast.body,
				BroadcastAt: helper.GetDateByFormatWithMs(),
				Counter:     c.handle.broadcastCounter,
			})
			err = c.conn.WriteMessage(websocket.TextMessage, []byte(jsonStrByte))
			if err != nil {
				log.Println("socketWriteErr: " + err.Error())
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
			c.pongAt = helper.GetDateByFormatWithMs()
		}
	}
}
