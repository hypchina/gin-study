package enum

//常量码规则
// 万位 0为成功码 1作保留 2参数级 3业务级 4系统级 | 注:千位1作保留
// 2-业务级 按千位分类 2为用户,3为订单 以此分类

const (
	StatusOk              = 0
	StatusParamIsError    = 20000 //参数级
	StatusDataIsNotExists = 30000 //业务级
	StatusDataOpError     = 39999 //业务级
	StatusAuthForbidden   = 40001 //权限
	StatusUnknownError    = 50000 //系统级
)

var statusText = map[int]string{
	StatusOk:              "ok",
	StatusParamIsError:    "参数错误",
	StatusDataIsNotExists: "数据不存在",
	StatusDataOpError:     "操作失败！请稍候重试",
	StatusAuthForbidden:   "未授权！",
	StatusUnknownError:    "未知错误！请稍候重试",
}

func StatusText(code int) string {
	text, ok := statusText[code]
	if ok {
		return text
	}
	return statusText[StatusUnknownError]
}
