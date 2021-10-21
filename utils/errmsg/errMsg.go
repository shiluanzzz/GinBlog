package errmsg

import "fmt"

const (
	SUCCESS = 200
	ERROR   = 500

	// error code for user
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WORON     = 1002
	ERROR_USER_NOT_EXIST     = 1003
	ERROR_TOKEN_RUNTIME      = 1004
	ERROR_USER_ROOT_NOT      = 1005
	ERROR_CATENAME_USED      = 2001
	ERROR_CATENAME_NOT_EXIST = 2002

	ERROR_ARTICLE_NOT_EXIST = 3001

	// token related
	ERROR_TOKEN_NOT_EXIST   = 4001
	ERROR_TOKEN_TIME_OUT    = 4002
	ERROR_TOKEN_WOKEN_WRONG = 4003
	ERROR_TOKEN_TYPE_WRONG  = 4004
)

var errMsg = map[int]string{
	SUCCESS: "成功",
	ERROR:   "系统错误,请联系管理员",
	// user
	ERROR_USERNAME_USED:  "用户名已经使用",
	ERROR_PASSWORD_WORON: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_USER_ROOT_NOT:  "用户权限不足",
	ERROR_TOKEN_RUNTIME:  "token超时",
	// 分类错误信息
	ERROR_CATENAME_USED:      "分类名称已经使用",
	ERROR_CATENAME_NOT_EXIST: "不存在",
	// 文章错误信息
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",

	//token
	ERROR_TOKEN_NOT_EXIST:   "TOKEN not exist",
	ERROR_TOKEN_TIME_OUT:    "token超时",
	ERROR_TOKEN_WOKEN_WRONG: "token 错误",
	ERROR_TOKEN_TYPE_WRONG:  "token格式错误",
}

func GetErrMsg(code int) string {
	if msg, ok := errMsg[code]; ok {
		return msg
	} else {
		return fmt.Sprintf("状态码%v未定义", code)
	}
}
