package ExtErr

import (
	"github.com/orangebees/go-oneutils/Convert"
	"strconv"
)

// Err  错误指针
type Err = *ErrorStruct

// ErrorStruct 错误结构
type ErrorStruct struct {
	code int
	msg  []byte
}

// 返回错误信息，标准错误接口实现的方法
func (e Err) Error() string {
	return Convert.B2S(e.msg)
}

// ErrorBytes 返回错误信息字节切片
func (e Err) ErrorBytes() []byte {
	return e.msg
}

// ErrorCode 返回错误码
func (e Err) ErrorCode() int {
	return e.code
}

// ErrorCodeString 返回错误码字符串
func (e Err) ErrorCodeString() string {
	return strconv.Itoa(e.code)
}

// SetErrorCode 设置错误码
func (e Err) SetErrorCode(code int) {
	e.code = code
}

// SetError 从标准错误设置错误信息
func (e Err) SetError(err error) {
	e.msg = append(e.msg[:0], err.Error()...)
}

// SetMsg  从标准错误设置错误信息
func (e Err) SetMsg(msg string) {
	e.msg = append(e.msg[:0], msg...)
}

// NewErr  新建错误
func NewErr(msg string) Err {
	t := ErrorStruct{
		code: 1,
	}
	t.SetMsg(msg)
	return &t
}
