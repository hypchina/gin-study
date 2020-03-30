package utils

import "encoding/json"

type RequestResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewRequestResponse() RequestResponse {
	return RequestResponse{
		Data: map[string]interface{}{},
	}
}

func (requestResponseX RequestResponse) Ok(data ...interface{}) RequestResponse {
	requestResponseX.Code = 0
	requestResponseX.Msg = "ok"
	if len(data) > 0 {
		requestResponseX.Data = data[0]
	}
	return requestResponseX
}

func (requestResponseX RequestResponse) Fail(msg string) RequestResponse {
	requestResponseX.Msg = msg
	requestResponseX.Code = 9999
	return requestResponseX
}

func (requestResponseX RequestResponse) ToJsonByte() []byte {
	bytes, _ := json.Marshal(requestResponseX)
	return bytes
}

type BroadcastResponse struct {
	RequestId   string `json:"request_id"`
	From        string `json:"from"`
	To          string `json:"to"`
	Type        int    `json:"type"`
	Body        string `json:"body"`
	BroadcastAt string `json:"broadcast_at"`
	Counter     int64  `json:"counter"`
}

type BroadcastRequest struct {
	To   string `json:"to" query:"to" binding:"required" validate:"min=4,max=32"`
	Type int    `json:"type" validate:"min=0,max=100000"`
	Body string `json:"body" query:"to" binding:"required" validate:"min=1,max=5000"`
}

func (BroadcastRequestX BroadcastRequest) ToJson() string {
	bytes, _ := json.Marshal(BroadcastRequestX)
	return string(bytes)
}
