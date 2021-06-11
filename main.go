package go_utils

import (
	"fmt"
	"github.com/dengjiawen8955/go_utils/syscall_util"
	"github.com/dengjiawen8955/go_utils/throw_util"
)

func main() {
	var err error
	listen, err := syscall_util.Listen(9999)
	throw_util.Throw(err)
	for {
		conn, err := listen.Accept()
		throw_util.Throw(err)
		go func() {
			defer conn.Close()
			fmt.Printf("conn.ClientFd=%#v\n", conn.ClientFd)
			fmt.Printf("conn.ServerFd=%#v\n", conn.ServerFd)
			conn.Write([]byte("hi,syscall_util"))
		}()
	}
}
