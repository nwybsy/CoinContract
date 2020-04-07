package validator

import "encoding/json"

type ResponseBase struct {
	ErrNo  int         `json:"code"`
	ErrMsg string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

type SuccessMsgResponseBase struct {
	ErrNo int         `json:"code"`
	Data  interface{} `json:"data"`
}

type SuccessResponseBase struct {
	ErrNo int `json:"code"`
}

func Response(error int, data interface{}) string {
	r := &SuccessMsgResponseBase{
		ErrNo: error,
		Data:  data,
	}
	s, _ := json.Marshal(r)
	return string(s)
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
