package models

type LogUserBroadcast struct {
	MsgId     string `json:"msg_id"`
	FromUid   string `json:"from_uid"`
	ToUid     string `json:"to_uid"`
	MsgLevel  int64  `json:"msg_level"`
	MsgType   int64  `json:"msg_type"`
	MsgBody   string `json:"msg_body"`
	ReadState int64  `json:"read_state"`
	ExpireAt  string `json:"expire_at"`
	NotifyAt  string `json:"notify_at"`
	ReadAt    string `json:"read_at"`
	CreatedAt string `json:"created_at"`
}
