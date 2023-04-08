package netu

import (
	"fmt"
	"testing"

	"github.com/jarvanstack/go_utils/logger"
)

func TestGetIpv4(t *testing.T) {
	ips, err := GetIpV4()
	if err != nil {
		logger.Error(err)
	}
	fmt.Printf("ips[0]: %v\n", ips[0])
	for i := 0; i < 100; i++ {
		ips2, err := GetIpV4()
		if err != nil {
			logger.Error(err)
		}
		if ips2[0] != ips[0] {
			fmt.Printf("%s\n", "重复匹配失败")
		}
	}
	fmt.Printf("%s\n", "重复匹配成功")
}
