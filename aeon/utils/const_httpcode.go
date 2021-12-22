package utils

import "fmt"

// ResponseCode 定义错误代码别名类型
type ResponseCode int

// GetResponseMsg 通过消息码获取Map中存储的对应的消息
func GetResponseMsg(code ResponseCode) (responseMsg string) {
	if tmpMsg, ok := responseMsgMap[code]; ok {
		responseMsg = tmpMsg
	} else {
		responseMsg = fmt.Sprintf("该错误码:%d 无对应的错误消息", code)
	}
	return
}

// HTTP状态码
// 添加消息码之后必须在下面Map添加消息码对应的消息
const (
	// 权限错误

	// ErrPermissionDenied 权限错误
	ErrPermissionDenied ResponseCode = 1100 + iota
	ErrTokenInvalid
	ErrTokenOverTime

	// 参数错误

	// ErrCodeParamError 参数错误
	ErrCodeParamError ResponseCode = 1000 + iota
	ErrCodeParamShouldBindError
	ErrCodeMissingParamError

	// 处理错误

	// ErrServerError 处理错误
	ErrServerError ResponseCode = 1200 + iota
	ErrUserNotFound
	ErrUserNotFoundOrErr

	// 数据错误

	// ErrCodeMySQLError 数据错误
	ErrCodeMySQLError ResponseCode = 1300 + iota
)

// 消息码对应的消息
var responseMsgMap = map[ResponseCode]string{
	ErrPermissionDenied: "权限不足",
	ErrTokenInvalid:     "无效token",
	ErrTokenOverTime:    "token超时",

	ErrCodeParamError:           "参数错误",
	ErrCodeParamShouldBindError: "参数解析错误",
	ErrCodeMissingParamError:    "缺少参数",

	ErrUserNotFound:      "用户不存在",
	ErrUserNotFoundOrErr: "用户不存在或密码错误",

	ErrCodeMySQLError: "数据库错误",
}
