package testu

import (
	"fmt"
	"testing"
	"time"
)

//做一个测试类，用来方便测试
//只需要一步的执行时间
//* 计算出 OPS
//* 循环次数 + 总用时 + ops
//* 计算出 memory.
//------run---------
//(1)[][][] TIME
//(2)[]LOOP TIMES
//(3)[]QPS
//------end---------
func Test_Float(t *testing.T) {
	fmt.Printf("float64(1/10)=%#v\n", float64(1)/float64(10))
}
//1.08
//(1.11s)
//(1.09s)
func Test_test_util(t *testing.T) {
	tu := NewTestUtil(1000)
	tu.Start()
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Microsecond)
	}
	tu.End()
}
//(1.14s)
//(1.11s)
func Test_sleep(t *testing.T) {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Microsecond)
	}
}