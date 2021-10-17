package errmsg

import "fmt"

const (
	SUCCESS = 200
	ERROR   = 500

	// error code for user
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WORON = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_RUNTIME  = 1004

	ERROR_CATENAME_USED = 2001
)

var errMsg = map[int]string{
	SUCCESS: "成功",
	ERROR:   "系统错误,请联系管理员",
	// user
	ERROR_USERNAME_USED:  "用户名已经使用",
	ERROR_PASSWORD_WORON: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_RUNTIME:  "token超时",
	// article
	ERROR_CATENAME_USED: "分类名称已经使用",
}

func GetErrMsg(code int) string {
	if msg, ok := errMsg[code]; ok {
		return msg
	} else {
		return fmt.Sprintf("状态码%v未定义", code)
	}
}
