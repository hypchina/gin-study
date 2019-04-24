package models

type UcUser struct {
	Id        int64  `json:"-"`
	OutUid    string `json:"out_uid"`
	UserName  string `json:"username" xorm:"username"`
	Password  string `json:"-"`
	TokenSlat string `json:"-"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
