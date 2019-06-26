package models

type LogSysRequest struct {
	Id               int64  `json:"-"`
	Uid              int64  `json:"uid"`
	RequestId        string `json:"request_id"`
	RequestApi       string `json:"request_api"`
	RequestMethod    string `json:"request_method"`
	RequestHeader    string `json:"request_header"`
	RequestIp        string `json:"request_ip"`
	RequestUri       string `json:"request_uri"`
	RequestBody      string `json:"request_body"`
	RequestAt        string `json:"request_at"`
	ResponseCode     int    `json:"response_code"`
	ResponseMsg      string `json:"response_msg"`
	ResponseBody     string `json:"response_body"`
	ResponseAt       string `json:"response_at"`
	ResponseDuration int64  `json:"response_duration"`
	CreatedAt        string `json:"created_at"`
}
