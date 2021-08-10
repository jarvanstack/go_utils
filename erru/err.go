package erru

import "fmt"

type Err struct {
	//例如 RequestError or ResponseError
	Code string `json:"code"`
	//错误的内容,原因,建议的解决
	Msg string `json:"msg"`
	//用来链路纠错等,可选
	RequestId string `json:"requestId"`
}

func (e *Err) Error() string {
	if e.RequestId == "" {
		return fmt.Sprintf("[Error]|Code=%s|Msg=%s", e.Code, e.Msg)
	}
	return fmt.Sprintf("[Error]|Code=%s|Msg=%s|RequestId=%s", e.Code, e.Msg, e.RequestId)
}
func NewError(code, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}
