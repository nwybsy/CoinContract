package validator

var MsgFlags = map[int]string{
	Succeed:         "操作成功",
	Unknown:         "未知错误",
	ServerErrDescDB: "网络繁忙，请重新打开",
	RedisErr:        "redis操作错误",
	DataBaseErr:     "数据库错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}
