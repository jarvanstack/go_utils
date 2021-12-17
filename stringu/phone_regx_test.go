package stringu

import (
	"fmt"
	"testing"
)

func TestIsPhone(t *testing.T) {
	s := []string{"18505921256", "1833023069", "19159029321"}
	for _, v := range s {
		fmt.Printf("p:%s,r=%v\n", v, IsPhone(v))
	}
}
