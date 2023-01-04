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

	// 文章管理错误
	ERROR_ARTICLE_FALSE_PAGE_PARAM      = 3000
	ERROR_ARTICLE_PARSE_PAGE_SIZE_PARAM = 3001
	ERROR_ARTICLE_WRITING               = 3002
	ERROR_ARTICLE_FAILED_TO_ADD_ARTICLE = 3003
)

var errMsg = map[int]string{
	ERROR_LOGIN_FAILED:                   "登录失败",
	ERROR_USERNAME_OR_PASSWORD_NOT_EXIST: "用户名或密码错误",
	ERROR_AUTH_NO_TOKEN:                  "没有token",
	ERROR_AUTH_INVALID_TOKEN:             "无效的token",

	ERROR_ARTICLE_FALSE_PAGE_PARAM:      "page参数错误",
	ERROR_ARTICLE_PARSE_PAGE_SIZE_PARAM: "page_size参数错误",
	ERROR_ARTICLE_WRITING:               "上传文章错误",
	ERROR_ARTICLE_FAILED_TO_ADD_ARTICLE: "添加文章失败",
}

func GetErrMsg(code int) string {
	return errMsg[code]
}
