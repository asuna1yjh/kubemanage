package common

import "errors"

type ResCode int64

var (
	ErrInvalidParams     = errors.New("请求参数错误")
	ErrLoginRequired     = errors.New("需要登录")
	ErrorInvalidToken    = errors.New("无效的token")
	ErrorUserNotLogin    = errors.New("用户未登录")
	ErrorServerBusy      = errors.New("服务繁忙")
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
)

var CodeMsg = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMsg[c]
	if !ok {
		return CodeMsg[CodeServerBusy]
	}
	return msg
}
