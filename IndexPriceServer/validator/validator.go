package validator

const (
	Succeed         = 0     // 请求成功
	Unknown         = 1000  // 未知错误
	ServerErrDescDB = 10001 // 网络繁忙，请重新打开；
	RedisErr        = 10002 // redis错误；
	DataBaseErr     = 10003 // 数据库错误
)
