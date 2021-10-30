package stringu

import (
	"fmt"
	"testing"
)

func TestTokenMarshal(t *testing.T) {
	m := make(map[string]interface{})
	m["k"] = "value"
	s, err := TokenMarshal(m, "config.PrivateKey", int64(60))
	fmt.Printf("err: %v\n", err)
	fmt.Printf("s: %v\n", s)
	mc, err := TokenUnmarshal(s, "config.PrivateKey")
	fmt.Printf("err: %v\n", err)
	i := mc["exp"]
	fmt.Printf("i: %v\n", i)
	i2 := mc["k"]
	fmt.Printf("i2: %v\n", i2)
}
