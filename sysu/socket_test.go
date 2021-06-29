package sysu

import (
	"fmt"
	"github.com/dengjiawen8955/go_utils/throwu"
	"net"
	"testing"
)

func Test_server(t *testing.T) {
	var err error
	listen, err := Listen(9999)
	throwu.Throw(err)
	for {
		conn, err := listen.Accept()
		throwu.Throw(err)
		throwu.Throw(err)
		go func() {
			defer conn.Close()
			conn.Write([]byte("hi,sysu"))
		}()
	}

}
func Test_client(t *testing.T) {
	var err error
	conn, err := net.Dial("tcp", ":9999")
	throwu.Throw(err)
	buf := make([]byte, 1024)
	readN, err := conn.Read(buf)
	throwu.Throw(err)
	fmt.Printf("buf[:readN]=%s\n", buf[:readN])

}
