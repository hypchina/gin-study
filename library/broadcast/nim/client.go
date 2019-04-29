package nim

import (
	"bytes"
	"gin-study/library/broadcast/nim/request"
	"io/ioutil"
	"net/http"
)

const (
	bashPath = "https://api.netease.im/nimserver/"
)

type config struct {
	appKey    string
	appSecret string
}

type client struct {
	config *config
}

func NewClient(appKey string, appSecret string) *client {
	return &client{
		config: &config{
			appKey:    appKey,
			appSecret: appSecret,
		},
	}
}

func (client *client) Do(request request.RequestInterface) (respStr string, err error) {

	values, err := request.GetQuery()
	if err != nil {
		return "", err
	}

	body := bytes.NewBuffer([]byte(values.Encode()))
	httpRequest, err := http.NewRequest(request.GetMethod(), request.GetUrl(bashPath), body)
	if err != nil {
		return "", err
	}

	headers := request.GetHeader(client.config.appKey, client.config.appSecret)
	for headerKey, headerVal := range headers {
		httpRequest.Header.Set(headerKey, headerVal)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return string(content), nil
}
