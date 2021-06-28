package string_util

import (
	"encoding/base64"
)

func GetMd5ByStr(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func GetMd5ByBytes(data []byte)string  {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
