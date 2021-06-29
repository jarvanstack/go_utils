package go_utils

import (
	"fmt"
	"github.com/dengjiawen8955/go_utils/sysu"
	"github.com/dengjiawen8955/go_utils/throwu"
)

func main() {
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
