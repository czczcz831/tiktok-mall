package packer

import (
	"runtime"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func logErrorWithStack(code int, err error) {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	hlog.Errorf("error code: %d, error: %v, stack: %s", code, err, buf)
}

// NewResponse 封装响应，根据 code 自动获取 msg
func NewSuccessResponse(data interface{}) Response {
	return Response{
		Code: DEFAULT_SUCCESS_CODE,
		Msg:  ErrorMessages[DEFAULT_SUCCESS_CODE],
		Data: data,
	}
}

func NewErrorResponse(err error) Response {
	// error 转为 MyError
	myErr, ok := err.(*MyError)
	if !ok {
		myErr = NewMyError(DEFAULT_ERROR_CODE, err)
	}
	errMsg, ok := ErrorMessages[myErr.Code]
	if !ok {
		logErrorWithStack(myErr.Code, myErr)
		myErr.Code = DEFAULT_ERROR_CODE
		errMsg = ErrorMessages[DEFAULT_ERROR_CODE]
	}
	return Response{
		Code: myErr.Code,
		Msg:  errMsg,
	}
}
