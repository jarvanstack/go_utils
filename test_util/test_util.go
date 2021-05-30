package test_util

import (
	"fmt"
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
const NS_TO_S = 1000000000
const NS_TO_MS = 1000000
type TU struct {
	start  time.Time
	n uint32 //number of loop times.
	comments string
}

func NewTestUtil(n uint32)*TU {
	return &TU{n: n}
}
func (test *TU)Start()  {
	fmt.Printf("%s\n", "------run---------")
	test.start = time.Now()
}
func  (test *TU)End()  {
	end := time.Now()
	useNano := end.UnixNano() - test.start.UnixNano()
	//(1)[%ds][%dms][%dns] TIME
	second := float64(useNano)/float64(NS_TO_S)
	fmt.Printf("(1)[%fs][%fms][%dns] TIME\n",second, float64(useNano)/float64(NS_TO_MS),useNano)
	//(2)[%d]LOOP TIMES
	fmt.Printf("(2)[%d]LOOP TIMES\n",test.n)
	//(3)[%d]QPS
	fmt.Printf("(3)[%f]QPS",float64(test.n)/second)
	fmt.Printf("%s\n", "------end---------")
}