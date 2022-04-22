// Package result
// @Time : 2022/4/11 10:18
// @Author : lee
// @File : code
// @Software: GoLand
package result

const (
	Ok = 0

	StatusBadRequest = 20400
	NotFound         = 20404

	Error                 = 20500
	ErrorRequestParameter = 20501
	ErrToken              = 50008
	ErrSign               = 50009
	NoOperationPermission = 50010
)

var MsgFlags = map[int]string{
	Ok: "ok",

	StatusBadRequest: "请求失败",
	NotFound:         "无响应",

	Error:                 "请求错误",
	ErrorRequestParameter: "请求参数错误",
	ErrToken:              "token错误",
	ErrSign:               "签名错误",
	NoOperationPermission: "没有操作权限",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return ""
}
