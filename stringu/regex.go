package stringu

import "regexp"

// 正则表达式工具

// 通过正则表达式或者子字符串.
//
// 1.第二个正则字符串使用 ` 飘号包裹可以解决转义问题
//
// 2. 小括号包裹子字符串
//
// 3. ? 问号代表非贪婪模式,默认贪婪模式
//
// Example:
//
//	subs := stringu.GetSubStringByRegex("tao123shi5han567", `(.*?)(\d+)(.*?)\d(.*)\d`)
//	for _, v := range subs {
//		fmt.Printf("%s\n", v)
//
//	}
//	//输出
//	// tao
//	// 123
//	// shi
//	// han56
func GetSubStringByRegex(str, regx string) []string {
	reg := regexp.MustCompile(regx)
	subs := reg.FindStringSubmatch(str)
	return subs[1:]
}
