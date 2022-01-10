package erru

import "fmt"

type Err2 struct {
	//例如 RequestError or ResponseError
	Code int `json:"code"`
	//错误的内容,原因,建议的解决
	Msg string `json:"msg"`
	//用来链路纠错等,可选
	Notice string `json:"notice"`
}

func (e *Err2) Error() string {
	return fmt.Sprintf("[Error]|Code=%d|Msg=%s|Notice=%s", e.Code, e.Msg, e.Notice)
}
func NewError2(code int, msg string) *Err2 {
	return &Err2{
		Code:   code,
		Msg:    msg,
		Notice: msg,
	}
}

//输出带提示  err
func NewErrorNotice2(code int, msg, notice string) *Err2 {
	return &Err2{
		Code:   code,
		Msg:    msg,
		Notice: notice,
	}
}
