package string_util

import (
	"crypto/sha1"
	"hash"
	"sync"
)
//集成一些 hash 算法

var once sync.Once
var s1 hash.Hash
func GetSha1ByStr(value string)[]byte  {
	once.Do(func() {
		s1 = sha1.New()
	})
	s1.Write([]byte(value))
	return s1.Sum(nil)
}
