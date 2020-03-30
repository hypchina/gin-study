package server

import (
	"encoding/json"
	"errors"
	"gin-study/app/core/helper"
	"gin-study/library/websocket/utils"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

type webSocketServer struct {
	handle   *handle
	upGrader websocket.Upgrader
	config   *utils.WebSocketServerConfig
}

func NewWebSocketServer(config *utils.WebSocketServerConfig) *webSocketServer {
	log.Println(config)
	return &webSocketServer{
		handle: newHandle(),
		upGrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		config: config,
	}
}

func (_this *webSocketServer) Run() {
	go _this.handle.listener()
	http.HandleFunc("/ws", _this.WebSocketListener)
	http.HandleFunc("/api/connect", _this.GetAuthUrl)
	http.HandleFunc("/api/broadcast", _this.Http2WebSocketListener)
	http.HandleFunc("/api/onlineCount", _this.GetOnlineCount)
	http.HandleFunc("/api/clientInfo", _this.GetClientInfo)
	_ = http.ListenAndServe(_this.config.Addr, nil)
}

func (_this *webSocketServer) GetAuthUrl(writer http.ResponseWriter, request *http.Request) {
	var bytes []byte
	clientId := request.URL.Query().Get("client_id")
	if clientId == "" {
		bytes = utils.NewRequestResponse().Fail("param client_id is missing").ToJsonByte()
		_, _ = writer.Write(bytes)
		return
	}
	bytes = utils.NewRequestResponse().Ok(_this.createAuthUrl(clientId)).ToJsonByte()
	_, _ = writer.Write(bytes)
}

func (_this *webSocketServer) WebSocketListener(writer http.ResponseWriter, request *http.Request) {
	conn, err := _this.upGrader.Upgrade(writer, request, nil)
	var bytes []byte
	if err != nil {
		bytes = utils.NewRequestResponse().Fail(err.Error()).ToJsonByte()
		_, _ = writer.Write(bytes)
		return
	}
	signFilter, err := _this.checkSign(request)
	if err != nil {
		log.Println("checkSign", err.Error())
		bytes = utils.NewRequestResponse().Fail(err.Error()).ToJsonByte()
		log.Println(string(bytes))
		expectedErr := &websocket.CloseError{Code: websocket.CloseNormalClosure, Text: string(bytes)}
		_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(expectedErr.Code, expectedErr.Text))
		return
	}
	client := newClient(signFilter.ClientId, _this.handle, conn)
	_this.handle.register <- client
	go client.read()
	go client.write()
}

func (_this *webSocketServer) Http2WebSocketListener(writer http.ResponseWriter, request *http.Request) {

	var bytes []byte
	signFilter, err := _this.checkSign(request)
	if err != nil {
		bytes = utils.NewRequestResponse().Fail(err.Error()).ToJsonByte()
		_, _ = writer.Write(bytes)
		return
	}

	var requestFilter utils.BroadcastRequest
	broadcastRequest, err := _this.checkParam(requestFilter, signFilter.Body)
	if err != nil {
		bytes = utils.NewRequestResponse().Fail(err.Error()).ToJsonByte()
		_, _ = writer.Write(bytes)
		return
	}

	_this.handle.broadcast <- broadcast{
		fromId:    signFilter.ClientId,
		fromIp:    request.RemoteAddr,
		toId:      broadcastRequest.To,
		msgType:   broadcastRequest.Type,
		body:      broadcastRequest.Body,
		receiveAt: helper.GetDateByFormatWithMs(),
	}

	bytes = utils.NewRequestResponse().Ok().ToJsonByte()
	_, _ = writer.Write(bytes)
}

func (_this *webSocketServer) GetOnlineCount(writer http.ResponseWriter, request *http.Request) {
	var bytes []byte
	response := map[string]interface{}{}
	response["count"] = _this.handle.onlineCount()
	bytes = utils.NewRequestResponse().Ok(response).ToJsonByte()
	_, _ = writer.Write(bytes)
}

func (_this *webSocketServer) GetClientInfo(writer http.ResponseWriter, request *http.Request) {
	var bytes []byte
	clientId := request.URL.Query().Get("client_id")
	if clientId == "" {
		bytes = utils.NewRequestResponse().Fail("param client_id is missing").ToJsonByte()
		_, _ = writer.Write(bytes)
		return
	}
	response := map[string]interface{}{}
	response["clients"] = _this.handle.getClientInfo(clientId)
	bytes = utils.NewRequestResponse().Ok(response).ToJsonByte()
	_, _ = writer.Write(bytes)
}

func (_this *webSocketServer) checkSign(request *http.Request) (signFilter *utils.SignFilter, err error) {
	var params url.Values
	if request.Method == "GET" {
		params = request.URL.Query()
	} else if request.Method == "POST" {
		_ = request.ParseForm()
		params = request.Form
	} else {
		return nil, errors.New("method is not allow")
	}
	signTool := utils.NewSign(_this.config.AppId, _this.config.AppSecret)
	return signTool.Verify(params)
}

func (_this *webSocketServer) checkParam(filter utils.BroadcastRequest, body string) (*utils.BroadcastRequest, error) {
	err := json.Unmarshal([]byte(body), &filter)
	if err != nil {
		return nil, err
	}
	if err := utils.GetValidator().Struct(filter); err != nil {
		return nil, utils.FormatValidatorError(err)
	}
	return &filter, err
}

func (_this *webSocketServer) createAuthUrl(clientId string) (url string) {
	signTool := utils.NewSign(_this.config.AppId, _this.config.AppSecret)
	return signTool.CreateAuthLink(clientId)
}
