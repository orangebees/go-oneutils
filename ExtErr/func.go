package ExtErr

// Err  错误指针
type Err = *ErrorStruct

// ErrorStruct 错误结构
type ErrorStruct struct {
	code int
	msg  string
}

// 返回错误信息，标准错误接口实现的方法
func (e Err) Error() string {
	return e.msg
}

// ErrorCode 返回错误码
func (e Err) ErrorCode() int {
	return e.code
}

// SetErrorCode 设置错误码
func (e Err) SetErrorCode(code int) {
	e.code = code
}

// SetError 从标准错误设置错误信息
func (e Err) SetError(err error) {
	e.msg = err.Error()
}

// SetMsg  从标准错误设置错误信息
func (e Err) SetMsg(msg string) {
	e.msg = msg
}

// NewErr  新建错误
func NewErr(msg string) Err {
	return &ErrorStruct{
		code: 0,
		msg:  msg,
	}
}
