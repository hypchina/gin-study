package entity

type TokenEntity struct {
	Uid         int64  `json:"uid"`
	ClientId    string `json:"client_id"`
	ClientToken string `json:"client_token"`
	ExpireAt    int64  `json:"expire_at"`
}
