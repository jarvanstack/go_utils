package syscall_util

import (
	"fmt"
	"github.com/dengjiawen8955/go_utils/throw_util"
	"net"
	"testing"
)

func Test_server(t *testing.T) {
	var err error
	listen, err := Listen(9999)
	throw_util.Throw(err)
	for  {
		conn := listen.Accept()
		go func() {
			defer conn.Close()
			conn.Write([]byte("hi,syscall_util"))
		}()
	}

}
func Test_client(t *testing.T) {
	var err error
	conn, err := net.Dial("tcp", ":9999")
	throw_util.Throw(err)
	buf := make([]byte,1024)
	readN, err := conn.Read(buf)
	throw_util.Throw(err)
	fmt.Printf("buf[:readN]=%s\n", buf[:readN])

}