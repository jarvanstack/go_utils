package stringu

import (
	"fmt"
	"regexp"
)

// 通过正则表达式或者子字符串.
//
// 1.字符串使用 ` 飘号包裹可以解决转义问题
//
// 2. 小括号包裹子字符串
//
// 3. ? 问号代表非贪婪模式,默认贪婪模式
//
// Example:
//
// 	str := `Content-Disposition: form-data; name="file"; filename="Snipaste.png"`
// 	regx := `Content-Disposition: (.*?); name="(.*?)"; filename="(.*?)"`
// 	subs := stringu.GetSubStringByRegex(str, regx)
// 	for _, s := range subs {
// 		fmt.Printf("%s\n", s)
// 	}
// 	// form-data
// 	// file
// 	// Snipaste.png
func GetSubStringByRegex(str, regx string) ([]string, error) {
	reg := regexp.MustCompile(regx)
	subs := reg.FindStringSubmatch(str)
	if len(subs) >= 1 {
		return subs[1:], nil
	}
	return nil, fmt.Errorf("[error]:reg result is nil")
}
