package validator

import "encoding/json"

type ResponseBase struct {
	ErrNo  int         `json:"code"`
	ErrMsg string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

type SuccessResponseBase struct {
	ErrNo int `json:"code"`
}

// 返回成功信息
func SuccessResponse(error int) string {
	r := &SuccessResponseBase{
		ErrNo: error,
	}
	s, _ := json.Marshal(r)
	return string(s)
}

// 返回错误信息
func ErrResponse(error int) string {
	r := &ResponseBase{
		ErrNo:  error,
		ErrMsg: GetMsg(error),
	}
	s, _ := json.Marshal(r)
	return string(s)
}
