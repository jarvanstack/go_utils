package sysu

// Conn
//serverFd int
//clientFd int
type Conn interface {
	Read([]byte)(readN int,err error)
	Write([]byte)(writeN int,err error)
	Close() error
}
