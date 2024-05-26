package xerr

var codeText = map[int]string{
	SERVER_COMMON_ERROR: "服务器异常",
	REQUEST_PARAM_ERROR: "请求参数异常",
	DB_ERROR:            "数据库异常",
}

func ErrMsg(errCode int) string {
	str, ok := codeText[errCode]
	if ok {
		return str
	}

	return codeText[SERVER_COMMON_ERROR]
}
