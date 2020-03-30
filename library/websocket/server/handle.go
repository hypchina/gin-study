package server

import (
	"fmt"
	"gin-study/library/websocket/utils"
	"log"
)

type handle struct {
	broadcastCounter int64
	clientCounter    int64
	broadcast        chan broadcast
	register         chan *client
	unregister       chan *client
	clientMap        map[*client]string
	clientIdMap      map[string]map[*client]int
}

func newHandle() *handle {
	return &handle{
		clientCounter:    0,
		broadcastCounter: 0,
		broadcast:        make(chan broadcast),
		register:         make(chan *client),
		unregister:       make(chan *client),
		clientMap:        make(map[*client]string),
		clientIdMap:      make(map[string]map[*client]int),
	}
}

func (handle *handle) listener() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("listener is error:", err)
		}
	}()
	for {
		select {
		case client := <-handle.register:
			handle.registerEvent(client)
		case client := <-handle.unregister:
			handle.unregisterEvent(client)
		case broadcast := <-handle.broadcast:
			handle.broadcastEvent(broadcast)
		}
	}
}

func (handle *handle) onlineCount() int64 {
	return handle.clientCounter
}

func (handle *handle) getClientInfo(clientId string) []utils.ClientInfo {
	var clientInfoSet []utils.ClientInfo
	if clientSet, ok := handle.clientIdMap[clientId]; ok {
		for client := range clientSet {
			clientInfoSet = append(clientInfoSet, utils.ClientInfo{
				ClientId:         clientId,
				Addr:             client.conn.RemoteAddr().String(),
				ConnectAt:        client.connectAt,
				PongAt:           client.pongAt,
				BroadcastCounter: client.broadcastCounter,
			})
		}
	}
	return clientInfoSet
}

func (handle *handle) registerEvent(clientX *client) {
	log.Println("register:" + clientX.clientId)
	handle.clientMap[clientX] = clientX.clientId
	clientSet := map[*client]int{}
	if _, ok := handle.clientIdMap[clientX.clientId]; ok {
		clientSet = handle.clientIdMap[clientX.clientId]
	}
	clientSet[clientX] = 1
	handle.clientIdMap[clientX.clientId] = clientSet
	handle.clientCounter++
}

func (handle *handle) unregisterEvent(clientX *client) {
	log.Println("unregister:" + clientX.clientId)
	if _, ok := handle.clientMap[clientX]; ok {
		handle.clientCounter--
		delete(handle.clientIdMap[clientX.clientId], clientX)
		delete(handle.clientMap, clientX)
		close(clientX.broadcast)
		_ = clientX.conn.Close()
	}
}

func (handle *handle) broadcastEvent(broadcast broadcast) {
	if _, ok := handle.clientIdMap[broadcast.toId]; ok {
		clientSet := handle.clientIdMap[broadcast.toId]
		for client := range clientSet {
			client.broadcast <- broadcast
		}
	}
}
