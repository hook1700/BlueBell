package controller

type Recode int64

const (
	CodeSuccess Recode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInValidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeCommunityIsNull
	CodeGetCommunityFail
)

var codeMsgMap = map[Recode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "请求参数错误",
	CodeUserExist:        "用户不存在",
	CodeUserNotExist:     "用户名不存在",
	CodeInValidPassword:  "用户名或密码错误",
	CodeServerBusy:       "服务器繁忙",
	CodeNeedLogin:        "用户未登录",
	CodeInvalidToken:     "token参数错误",
	CodeCommunityIsNull:  "社区列表为空",
	CodeGetCommunityFail: "获取社区失败",
}

func (c Recode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
