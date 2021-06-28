package base_util

import (
	"fmt"
	"testing"
)

func Test_bytes_to_int(t *testing.T) {
	bs, _ := IntToBytes(1<<10,8)
	fmt.Printf("bs=%b\n", bs)
}
func Test_int_max_const(t *testing.T)  {
	fmt.Printf("Uint8Max=%d\n", Uint8Max)
	fmt.Printf("Uint16Max=%d\n", Uint16Max)
	fmt.Printf("Uint32Max=%d\n", Uint32Max)
	fmt.Printf("Uint64Max=%d\n", Uint64Max)
	fmt.Printf("Int8Max=%d\n", Int8Max)
	fmt.Printf("Int16Max=%d\n", Int16Max)
	fmt.Printf("Int32Max=%d\n", Int32Max)
	fmt.Printf("Int64Max=%d\n", Int64Max)
}