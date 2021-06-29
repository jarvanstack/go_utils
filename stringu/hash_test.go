package stringu

import (
	"fmt"
	"testing"
)

func Test_hash_sha1(t *testing.T) {
	b := GetSha1ByStr("hello")
	md := GetMd5ByBytes(b)
	fmt.Printf("md=%#v\n", md)
}
