package mq

type JobStruct struct {
	JobId    string `json:"job_id"`
	Topic    string `json:"topic"`
	JobAt    int64  `json:"job_at"`
	Body     string `json:"body"`
	Tag      string `json:"tag"`
	ExpireAt int64  `json:"-"`
	Delay    int64  `json:"delay"`
	CreateAt string `json:"create_at"`
}

type jobDelayStruct struct {
	JobId string `json:"job_id"`
	Topic string `json:"topic"`
}
