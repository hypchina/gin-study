package models

type LogSysRequest struct {
	Id            int64  `json:"-"`
	RequestId     string `json:"request_id"`
	RequestAt     string `json:"request_at"`
	RequestHeader string `json:"request_header"`
	RequestIp     string `json:"request_ip"`
	RequestUri    string `json:"request_uri"`
	RequestBody   string `json:"request_body"`
	ResponseAt    string `json:"response_at"`
	ResponseCode  string `json:"response_code"`
	ResponseMsg   string `json:"response_msg"`
	ResponseBody  string `json:"response_body"`
	CreatedAt     string `json:"created_at"`
}
