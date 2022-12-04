package constant

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

const (
	//登录错误
	ERROR_LOGIN_FAILED                   = 1001
	ERROR_USERNAME_OR_PASSWORD_NOT_EXIST = 1002

	// 认证错误
	ERROR_AUTH_NO_TOKEN      = 2001
	ERROR_AUTH_INVALID_TOKEN = 2002
)

var errMsg = map[int]string{
	ERROR_LOGIN_FAILED:                   "登录失败",
	ERROR_USERNAME_OR_PASSWORD_NOT_EXIST: "用户名或密码错误",
	ERROR_AUTH_NO_TOKEN:                  "没有token",
	ERROR_AUTH_INVALID_TOKEN:             "无效的token",
}

func GetErrMsg(code int) string {
	return errMsg[code]
}
