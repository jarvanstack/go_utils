package netu

import (
	"net"

	"github.com/jarvanstack/go_utils/logger"
)

func GetIpV4() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	ips = make([]string, 0)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// fmt.Println(ipnet.IP.String())
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return
}
