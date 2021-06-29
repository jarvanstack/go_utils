package baseu

import (
	"fmt"
	"testing"
)

func Test_bytes_to_int(t *testing.T) {
	bs, _ := IntToBytes(1<<10,8)
	fmt.Printf("bs=%b\n", bs)
}

//UINT8_MAX=255
//UINT16_MAX=65535
//UINT32_MAX=4294967295
//UINT64_MAX=18446744073709551615
//INT8_MAX=127
//INT16_MAX=32767
//INT32_MAX=2147483647
//INT64_MAX=9223372036854775807
func Test_int_max_const(t *testing.T)  {
	fmt.Printf("UINT8_MAX=%d\n", UINT8_MAX)
	fmt.Printf("UINT16_MAX=%d\n", UINT16_MAX)
	fmt.Printf("UINT32_MAX=%d\n", UINT32_MAX)
	fmt.Printf("UINT64_MAX=%d\n", UINT64_MAX)
	fmt.Printf("INT8_MAX=%d\n", INT8_MAX)
	fmt.Printf("INT16_MAX=%d\n", INT16_MAX)
	fmt.Printf("INT32_MAX=%d\n", INT32_MAX)
	fmt.Printf("INT64_MAX=%d\n", INT64_MAX)
}