package restful_util

import (
	"fmt"
	"testing"
)

func TestNotFound(t *testing.T) {
	data := NotFound("fuck")
	fmt.Printf("data=%#v\n", string(data))
}
