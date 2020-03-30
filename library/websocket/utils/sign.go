package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net/url"
	"strconv"
	"time"
)

type SignFilter struct {
	AppId     string `json:"app_id" query:"app_id" form:"app_id" validate:"len=32"`
	ClientId  string `json:"client_id" query:"client_id"  form:"client_id" validate:"min=4,max=64"`
	Body      string `json:"body" query:"body" form:"body" validate:"min=0,max=5000"`
	Once      string `json:"once" query:"once" form:"once" validate:"len=32"`
	Timestamp string `json:"timestamp" query:"timestamp" form:"timestamp" validate:"min=10,max=13"`
	Sign      string `json:"sign" query:"sign" form:"sign" validate:"len=32"`
}

type sign struct {
	appId     string
	appSecret string
	expireIn  int
}

func NewSign(appId string, appSecret string) *sign {
	return &sign{
		appId:     appId,
		appSecret: appSecret,
		expireIn:  300,
	}
}

func (_this sign) Verify(param url.Values) (*SignFilter, error) {
	SignFilter := &SignFilter{}
	if err := mapForm(SignFilter, param, "query"); err != nil {
		return nil, err
	}
	if err := GetValidator().Struct(SignFilter); err != nil {
		return nil, FormatValidatorError(err)
	}
	if signStr := _this.Create(*SignFilter); signStr != SignFilter.Sign {
		return SignFilter, errors.New("invalid sign")
	}

	return SignFilter, nil
}

func (_this sign) Create(signFilterX SignFilter) string {
	params := struct2MapX(signFilterX)
	paramStr := httpBuildQuery(params, true)
	paramStr += _this.appSecret
	return md5X(paramStr)
}

func (_this sign) CreateAuthLink(clientId string) string {
	signFilterX := SignFilter{
		AppId:     _this.appId,
		ClientId:  clientId,
		Body:      "ping",
		Once:      createNonce(32),
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		Sign:      "",
	}
	signFilterX.Sign = _this.Create(signFilterX)
	params := struct2MapX(signFilterX)
	paramStr := httpBuildQuery(params, false)
	return paramStr
}

func createNonce(strYLen int) string {
	strX := "12345567890abcdefghijklmnopqrstuvwxyz"
	strY := ""
	if strYLen < 1 {
		return ""
	}
	strXLen := int64(len(strX))
	for i := 0; i < strYLen; i++ {
		randX, _ := rand.Int(rand.Reader, big.NewInt(strXLen))
		start := randX.Int64()
		end := start + 1
		strY += strX[start:end]
	}
	return strY
}
