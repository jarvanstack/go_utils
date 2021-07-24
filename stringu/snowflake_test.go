package stringu

import (
	"fmt"
	"testing"
)

func Test_snowflake(t *testing.T) {
	w := NewWorker(0, 0)
	fmt.Println(w.NextID())
	//156004662263152640
}
