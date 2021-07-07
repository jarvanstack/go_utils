package throwu

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"runtime"
	"strings"
)
// Trace 返回错误的 trace 自己打印即可
// Usage:
//	defer func() {
//		if err := recover(); err != nil {
//			log.Printf("%s\n\n", throwu.Trace(err))
//			ctx.Json(restfulu.ServerError("SERVER_ERROR"))
//		}
//	}()
func Trace(err interface{}) string {
	message := fmt.Sprintf("%s", err)
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller
	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
func catch()  {
	err := recover()
	if err != nil {
		err.(errorx.Error).PrintStackTrace()
	}
}
func Throw(err error) {
	defer catch()
	if err != nil {
		panic(err)
	}
}
