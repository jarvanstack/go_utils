package sysu

import (
	"fmt"
	"net"
	"testing"

	"github.com/jarvanstack/go_utils/erru"
)

func Test_server(t *testing.T) {
	var err error
	listen, err := Listen(9999)
	erru.Throw(err)
	for {
		conn, err := listen.Accept()
		erru.Throw(err)
		erru.Throw(err)
		go func() {
			defer conn.Close()
			conn.Write([]byte("hi,sysu"))
		}()
	}

}
func Test_client(t *testing.T) {
	var err error
	conn, err := net.Dial("tcp", ":9999")
	erru.Throw(err)
	buf := make([]byte, 1024)
	readN, err := conn.Read(buf)
	erru.Throw(err)
	fmt.Printf("buf[:readN]=%s\n", buf[:readN])

}
