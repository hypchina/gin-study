package models

type LogCollect struct {
	AppName  string `json:"app_name"`
	TraceId  string `json:"trace_id"`
	EventAt  string `json:"event_at"`
	LogType  string `json:"log_type"`
	LogLevel string `json:"log_level"`
	LogBody  string `json:"log_body"`
}
