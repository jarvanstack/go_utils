package throw_util

import (
	"github.com/fwhezfwhez/errorx"
)

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
