package syscall_util

import (
	"fmt"
	"net"
	"syscall"
)



const (
	ip = "127.0.0.1"
	maxListenNumber = 10
	errCode = -1
)

type SocketServer struct {
	Fd int
}
func (s *SocketServer) Accept() (*SocketConn) {
	var err error
	clientFd, _, err := syscall.Accept(s.Fd)
	if err != nil {
		return nil
	}
	return &SocketConn{ServerFd: s.Fd,ClientFd: clientFd}
}


//SocketConn socketServerFd.Accept() to get
type SocketConn struct {
	ServerFd int
	ClientFd int
}
func (s *SocketConn) Write(buf []byte) (int, error) {
	return syscall.Write(s.ClientFd, buf)
}
func (s *SocketConn) Read(buf []byte) (int,error) {
	return syscall.Read(s.ClientFd,buf)
}
func (s *SocketConn) Close() error {
	return syscall.Close(s.ClientFd)
}
// Listen socket -> bind -> listen
// Return SocketServer pointer,nil if listen success.
// Return nil,err if Here have some errors.
func Listen(port int)(server *SocketServer,err error)  {
	var serverAddr *syscall.SockaddrInet4
	var serverFd int
	//1.socket
	syscall.ForkLock.Lock()
	serverFd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	syscall.ForkLock.Unlock()
	if err = syscall.SetsockoptInt(serverFd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		fmt.Printf("%s\n", "error: socket()")
		syscall.Close(serverFd)
		return nil,err
	}
	//2.bind
	serverAddr = &syscall.SockaddrInet4{Port: port}
	copy(serverAddr.Addr[:], net.ParseIP(ip))
	if err = syscall.Bind(serverFd, serverAddr); err != nil {
		fmt.Printf("error:PORT ALREADY IN USE:%d \n",port)
		return nil,err
	}
	//3.Listen.
	if err = syscall.Listen(serverFd, maxListenNumber); err != nil {
		fmt.Printf("%s\n", "error: listen()")
		return nil,err
	}
	return &SocketServer{Fd: serverFd},nil
}

