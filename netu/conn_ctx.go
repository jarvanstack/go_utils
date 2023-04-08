package netu

import (
	"bufio"
	"io"
	"net"

	"github.com/jarvanstack/go_utils/baseu"
	"github.com/jarvanstack/go_utils/erru"
)

//封装粘包协议 + reader + writer 缓冲
type ConnCtx struct {
	//继承所有的这个功能. read write
	R *bufio.Reader
	W *bufio.Writer
}

//构建 ConnCtx 对象
func NewConnCtx(conn net.Conn) *ConnCtx {
	return &ConnCtx{
		R: bufio.NewReader(conn),
		W: bufio.NewWriter(conn),
	}
}

//粘包协议数据发送,自带 flush
func (c *ConnCtx) WriteData(data []byte) error {
	var err error
	l, err := baseu.IntToBytes(len(data), 4)
	if err != nil {
		return erru.NewError("baseu.IntToBytes", "transition err while Write to client")
	}
	c.W.Write(l)
	c.W.Write(data)
	//没有flush不行
	c.W.Flush()
	return nil
}

//粘包协议数据读取
func (c *ConnCtx) ReadData() ([]byte, error) {
	var err error
	lb := make([]byte, 4)
	//使用 io.ReadFull 可以完全读取字节
	io.ReadFull(c.R, lb)
	l, err := baseu.BytesToInt(lb, false)
	if err != nil {
		return nil, erru.NewError("baseu.IntToBytes", "transition err while read from client")
	}
	data := make([]byte, l)
	io.ReadFull(c.R, data)
	return data, nil
}
