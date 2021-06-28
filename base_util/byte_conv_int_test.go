package base_util

import (
	"fmt"
	"testing"
)

func Test_bytes_to_int(t *testing.T) {
	bs, _ := Int64ToBytes(1<<10)
	fmt.Printf("bs=%b\n", bs)
}
