package main

import (
	"fmt"

	"github.com/dengjiawen8955/go_utils/stringu"
	"github.com/dengjiawen8955/go_utils/sysu"
	"github.com/dengjiawen8955/go_utils/throwu"
)

func main() {
	subs := stringu.GetSubStringByRegex("tao123shi5han567", `(.*?)(\d+)(.*?)\d(.*)\d`)
	for _, v := range subs {
		fmt.Printf("%s\n", v)

	}
	//输出
	// tao
	// 123
	// shi
	// han56
}
func t1() {
	fmt.Printf("%s\n", "start")
	var err error
	listen, err := sysu.Listen(9999)
	throwu.Throw(err)
	for {
		conn, err := listen.Accept()
		throwu.Throw(err)
		go func() {
			defer conn.Close()
			fmt.Printf("conn.ClientFd=%#v\n", conn.ClientFd)
			fmt.Printf("conn.ServerFd=%#v\n", conn.ServerFd)
			conn.Write([]byte("hi,sysu"))
		}()
	}
}
