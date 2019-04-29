package request

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

type requestBody struct {
	From string `url:"from"`
	To   string `url:"to"`
	Body string `url:"body"`
	Type int    `url:"type"`
	Ope  int    `url:"ope"`
}

type sendMsgRequest struct {
	*nimRequest
	requestBody requestBody
}

func NewSendMsgRequest(fromUser string, toUser string, body string) *sendMsgRequest {
	return &sendMsgRequest{
		nimRequest: &nimRequest{
			api: "sendMsg.action",
		},
		requestBody: requestBody{
			From: fromUser,
			To:   toUser,
			Body: body,
			Type: 0,
			Ope:  0,
		},
	}
}

func (request *sendMsgRequest) GetQuery() (url.Values, error) {
	return query.Values(request.requestBody)
}

func (request *sendMsgRequest) GetUrl(basePath string) string {
	return "http://local.mini.com/test"
}
