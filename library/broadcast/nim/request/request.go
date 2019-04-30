package request

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/google/go-querystring/query"
	"math/big"
	"net/url"
	"strconv"
	"time"
)

type IRequest interface {
	GetQuery() (url.Values, error)
	GetHeader(string, string) map[string]string
	GetMethod() string
	GetUrl(string) string
}

type nimRequest struct {
	api       string
	queryData map[string]interface{}
}

func (request *nimRequest) GetQuery() (url.Values, error) {
	return query.Values(request.queryData)
}

func (request *nimRequest) GetMethod() string {
	return "POST"
}

func (request *nimRequest) GetUrl(basePath string) string {
	return basePath + request.api
}

func (request *nimRequest) GetHeader(appKey string, appSecret string) map[string]string {
	nonce := createNonce(32)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded;param=value",
		"AppKey":       appKey,
		"Nonce":        nonce,
		"CurTime":      curTime,
		"CheckSum": func(appSecret string, nonce string, curTime string) string {
			h := sha1.New()
			h.Write([]byte(appSecret + nonce + curTime))
			hash := h.Sum(nil)
			return fmt.Sprintf("%x", hash)
		}(appSecret, nonce, curTime),
	}
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
