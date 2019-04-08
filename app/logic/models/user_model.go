package models

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	State    int    `json:"state"`
}
