package controller

type ResCode int64

const (
	CodeSuccess         ResCode = 1000 + iota // 成功
	CodeInvalidParam                          // 无效参数
	CodeUserExist                             // 用户已存在
	CodeUserNotExist                          // 用户不存在
	CodeInvalidPassword                       // 无效密码
	CodeServerBusy                            // 服务器繁忙

	CodeInvalidToken // 需要登录
	CodeNeedLogin    // 无效令牌
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidToken:    "无效的token",
	CodeNeedLogin:       "需要登录",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
