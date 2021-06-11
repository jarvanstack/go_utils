package syscall_util

// Conn
//serverFd int
//clientFd int
type Conn interface {
	Read([]byte)(readN int,err error)
	Write([]byte)(writeN int,err error)
	Close() error
}
type ServerFd interface {
	Accept() (conn *Conn,err error)
}