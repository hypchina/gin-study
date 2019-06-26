package filters

type baseFilter struct {
	MapData map[string]interface{}
}

type AuthToken struct {
	ClientId  string `form:"client_id" uri:"client_id" binding:"required"`
	Timestamp int64  `form:"timestamp" uri:"timestamp" binding:"required"`
}

type UserRegister struct {
	UserName string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
	baseFilter
}

type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
	baseFilter
}

type BroadcastNotifyFilter struct {
	ToUid   string `form:"to_uid" uri:"to_uid"  binding:"required"`
	MsgType int64  `form:"msg_type" uri:"msg_type"  binding:"required"`
	MsgBody string `form:"msg_body" uri:"msg_body"  binding:"required"`
}

type LogCollectFilter struct {
	AppName  string `form:"app_name"  form:"app_name"  binding:"required"`
	TraceId  string `form:"trace_id"  form:"trace_id"`
	LogType  string `form:"log_type"  form:"log_type"  binding:"required"`
	LogLevel string `form:"log_level" form:"log_level"  binding:"required"`
	LogBody  string `form:"log_body"  form:"log_body"  binding:"required"`
	EventAt  string `form:"event_at"  form:"event_at"  binding:"required"`
}
