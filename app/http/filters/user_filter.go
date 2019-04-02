package filters

type baseFilter struct {
	MapData map[string]interface{}
}

type UserFilter struct {
	UserName string `form:"user_name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string
	baseFilter
}
