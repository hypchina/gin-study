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
