package main

import (
	"fmt"

	"github.com/dengjiawen8955/go_utils/erru"
	"github.com/dengjiawen8955/go_utils/logger"
	"github.com/dengjiawen8955/go_utils/stringu"
	"github.com/dengjiawen8955/go_utils/sysu"
)

func main() {
	logger.Info("info")
	logger.Debug("debug")
	logger.Error("error")
}
func t1() {
	fmt.Printf("%s\n", "start")
	var err error
	listen, err := sysu.Listen(9999)
	erru.Throw(err)
	for {
		conn, err := listen.Accept()
		erru.Throw(err)
		go func() {
			defer conn.Close()
			fmt.Printf("conn.ClientFd=%#v\n", conn.ClientFd)
			fmt.Printf("conn.ServerFd=%#v\n", conn.ServerFd)
			conn.Write([]byte("hi,sysu"))
		}()
	}
}
func t2() {
	str := `Content-Disposition: form-data; name="file"; filename="Snipaste_2021-07-07_15-49-56.png"`
	regx := `Content-Disposition: (.*?); name="(.*?)"; filename="(.*?)"`
	subs, err := stringu.GetSubStringByRegex(str, regx)
	if err != nil {
		return
	}
	for _, s := range subs {
		fmt.Printf("%s\n", s)
	}
}
