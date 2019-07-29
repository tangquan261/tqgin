package errorcode

const (
	SUCCESS              = 0 //成功
	ERROR                = 1 //通用错误
	ERROR_INVALID_PARAMS = 2 //参数错误

	ERROR_AUTH_TOKEN_CHECK_FAIL = 20001 //token鉴权失败
	ERROR_AUTH_TOKEN_TIMEOUT    = 20002 //token超时
	ERROR_AUTH_TOKEN_CREATE     = 20003 //创建失败
	ERROR_AUTH_TOKEN            = 20004 //tocken错误

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001 //保存图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002 //检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003 //校验错误，格式或大小有误
)

var MsgFlags = map[int]string{
	SUCCESS:              "成功",
	ERROR:                "错误",
	ERROR_INVALID_PARAMS: "参数错误",

	ERROR_AUTH_TOKEN_CHECK_FAIL: "token鉴权失败",
	ERROR_AUTH_TOKEN_TIMEOUT:    "token超时",
	ERROR_AUTH_TOKEN_CREATE:     "创建失败",
	ERROR_AUTH_TOKEN:            "tocken错误",

	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验错误，格式或大小有误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
