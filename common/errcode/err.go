package errcode

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type AppError struct {
	// 唯一标识
	code int
	// 直观描述
	msg      string
	cause    error
	occurred string
}

func newError(code int, msg string) *AppError {
	if _, duplicated := codes[code]; duplicated {
		panic(fmt.Sprintf("the code %s already existed", code))
	}
	return &AppError{
		code: code,
		msg:  msg,
	}
}

// Error 实现了 error 接口
// 直接序列化，更快
func (a *AppError) Error() string {
	// 避免访问空指针成员
	if a == nil {
		return ""
	}
	bytes, err := json.Marshal(a.toStructuredError())
	if err != nil {
		return fmt.Sprintf("Error() is error: json marshal error: %v", err)
	}
	return string(bytes)
}

func (a *AppError) String() string {
	return a.Error()
}

func (a *AppError) Code() int {
	return a.code
}

func (a *AppError) Msg() string {
	return a.msg
}

func (a *AppError) HttpStatusCode() int {
	switch a.code {
	case Success.Code():
		return http.StatusOK
	case ErrServer.Code():
		return http.StatusInternalServerError
	case ErrParams.Code():
		return http.StatusBadRequest
	case ErrNotFound.Code():
		return http.StatusNotFound
	case ErrTooManyRequests.Code():
		return http.StatusTooManyRequests
	case ErrToken.Code():
		return http.StatusUnauthorized
	case ErrForbidden.Code():
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// Wrap 函数主要用户包装错误信息，用于日志记录
// 包装底层任务和 WithCause 一样是为了记录错误链条
func Wrap(msg string, err error) *AppError {
	return &AppError{
		msg:      msg,
		cause:    err,
		code:     -1,
		occurred: getErrorInfo(),
	}
}

// WithCause 复用预定好的错误信息
// 使用于错误码定义地比较详细的项目
func (a *AppError) WithCause(err error) *AppError {
	// 创建新的err，否则有并发问题
	newErr := a.Clone()
	newErr.cause = err
	newErr.occurred = getErrorInfo()
	return newErr
}

func getErrorInfo() string {
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("funName: %s file: %s line: %d", funcName, file, line)
}

func (e *AppError) Clone() *AppError {
	n := new(AppError)
	n.code = e.code
	n.msg = e.msg
	n.cause = e.cause
	n.occurred = e.occurred
	return n
}

// AppendMsg 在Code不变的情况下, 在预定义Msg的基础上追加错误信息
func (e *AppError) AppendMsg(msg string) *AppError {
	n := e.Clone()
	n.msg = fmt.Sprintf("%s, %s", e.msg, msg)
	return n
}

// SetMsg 在Code不变的情况下, 重新设置错误信息, 覆盖预定义的Msg
func (e *AppError) SetMsg(msg string) *AppError {
	n := e.Clone()
	n.msg = msg
	return n
}

type formattedErr struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Cause    interface{} `json:"cause"`
	Occurred string      `json:"occurred"`
}

// toStructuredError 在JSON Encode 前把Error进行格式化
func (e *AppError) toStructuredError() *formattedErr {
	fe := new(formattedErr)
	fe.Code = e.Code()
	fe.Msg = e.Msg()
	fe.Occurred = e.occurred
	if e.cause != nil {
		if appErr, ok := e.cause.(*AppError); ok {
			fe.Cause = appErr.toStructuredError()
		} else {
			fe.Cause = e.cause.Error()
		}
	}
	return fe
}
