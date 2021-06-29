package containeru

import (
	"fmt"
	"math"
)

//这是一个工具类，封装需要用到的数据结构
//一个基本栈的实现.
//Pop()
//push()
//IsEmpty()
//Peek()只看不取出来.
//size()

type Stack struct {
	//存放数据的动态数组.
	data[] interface{}
	//栈顶的指针
	topIndex int
	//最大长度容量，用来看看是否扩容.
	cap int
}
func  NewStack() *Stack {
	stack := &Stack{}
	//初始化初始容量为8
	stack.cap = 8
	stack.topIndex = -1
	stack.data = make([]interface{},8)
	return stack
}
//入栈.push
func (s *Stack) Push(item interface{})error {
	//判空.
	if s == nil {
		return fmt.Errorf("error:%s", "stack is null")
	}
	if s.cap >= math.MaxInt64 -1 {
		return fmt.Errorf("error:%s", "index out of int64 of stack array")
	}
	//先判断是否需要扩容
	if  s.topIndex >= s.cap-1{
		if s.cap<1024 {
			s.cap <<= 1
			temp := make([]interface{},s.cap)
			copy(temp[:],s.data)
			s.data = temp
		}else {
			s.cap = s.cap + (s.cap >> 2)
			temp := make([]interface{},s.cap)
			copy(temp[:],s.data)
			s.data = temp
		}
	}
	//赋值
	s.topIndex++
	s.data[s.topIndex] = item
	return nil
}
//出栈pop
func (s *Stack) Pop()(interface{})  {
	//判空.
	if s == nil {
		fmt.Printf("error:%s", "stack is null")
		return nil
	}
	//如果栈顶没有元素
	if s.topIndex == -1 {
		return nil
	}
	//出栈
	result := s.data[s.topIndex]
	s.data[s.topIndex] = nil
	s.topIndex--
	return result
}
//判断这个栈是否为空.
func (s *Stack) IsEmpty()bool  {
	if s == nil || s.topIndex == -1 {
		return true
	}
	return false
}
//只是看看不删除栈顶元素.
func (s *Stack) Peek() interface{} {
	//判空.
	if s == nil {
		return nil
	}
	//如果栈顶没有元素
	if s.topIndex == -1 {
		return nil
	}
	return s.data[s.topIndex]
}
//返回s.topIndex.
func (s *Stack) size()int  {
	if s == nil {
		return -1
	}
	return s.topIndex + 1
}

