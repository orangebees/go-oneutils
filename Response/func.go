package Response

import (
	"encoding/json"
	"github.com/orangebees/go-oneutils/ExtErr"
	"github.com/orangebees/go-oneutils/Fetch"
	"strings"
)

type EasyJsonResponse struct {
	Code int
	Msg  string
	Data Fetch.EasyJsonSerialization
}
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

// 0-9 为保留状态码
const (
	// SuccessCode 成功 相当于200
	SuccessCode = iota
	// FailCode 失败 相当于500
	FailCode
	// ParameterErrorCode 参数错误 相当于400
	ParameterErrorCode
	// AccessAbortedCode 访问中止 相当于403
	AccessAbortedCode
	// AccessAbortedCallbackCode 访问中止回调 相当于403+302 跳转到黑名单,登陆失败/失效的提示,工作页面 登录时不建议使用
	AccessAbortedCallbackCode
)

const (
	DefaultSuccessResponse        = `{"code":0,"msg":"ok"}`
	DefaultFailResponse           = `{"code":1,"msg":"Unknown error"}`
	DefaultParameterErrorResponse = `{"code":2,"msg":"Error or missing parameter"}`
	DefaultAccessAbortedResponse  = `{"code":3,"msg":"AccessFail"}`
)

// IsSuccess 快速判断是否成功，仅用于此包内生成的响应字符串
func IsSuccess(string2 string) bool {
	if len(string2) <= 9 {
		return false
	}
	if strings.HasPrefix(string2, DefaultSuccessResponse[:9]) {
		return true
	}
	return false
}

// EasyJsonSuccess  成功响应,data建议传指针类型的值
func EasyJsonSuccess(data Fetch.EasyJsonSerialization) string {
	if data == nil {
		return DefaultSuccessResponse
	}

	marshal, err := data.MarshalJSON()
	if err != nil {
		return ""
	}
	return `{"code":0,"msg":"ok","data":` + string(marshal) + "}"
}

// Success 成功响应
func Success(data any) string {
	if data == nil {
		return DefaultSuccessResponse
	}
	marshal, err := json.Marshal(Response{
		Data: data,
	})
	if err != nil {
		return ""
	}
	return string(marshal)
}

// Fail 失败响应
func Fail(err ExtErr.Err) string {
	if err == nil {
		return DefaultFailResponse
	}
	estr := err.Error()
	estrlen := len(estr)
	//删除特殊字符
	tmpbytes := make([]byte, estrlen)
	for i := 0; i < estrlen; i++ {
		t := estr[i]
		//跳过不可见非法字符
		if t == '\n' || t == '\t' {
			continue
		}
		//可见非法字符转换
		if t == '"' {
			tmpbytes = append(tmpbytes, '\\')
		}
		tmpbytes = append(tmpbytes, t)
	}
	return `{"code":` + err.ErrorCodeString() + `,"msg":"` + string(tmpbytes) + `"}`
}
