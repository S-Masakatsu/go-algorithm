// Stack (スタック)
// 後入れ先出し(LIFO: Last In First Out)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stack struct {
	data []interface{}
	size int
}

type StackImpl interface {
	Push(interface{})
	Pop() interface{}
	Top() interface{}
	IsEmpty() bool
	Size() int
	String() string
}

func NewStack() StackImpl {
	return &Stack{}
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	s.size--
	ret := s.data[s.size]
	s.data = s.data[:s.size]
	return ret
}

func (s *Stack) Top() interface{} {
	return s.data[s.size-1]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) String() string {
	return fmt.Sprint(s.data)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	stack := NewStack()
	for i := 1; i <= 5; i++ {
		stack.Push(rand.Intn(50))
	}
	fmt.Println("Size: ", stack.Size())
	fmt.Println("Data: ", stack.String())
	fmt.Println("Top:  ", stack.Top())
	fmt.Println("Pop:  ", stack.Pop())
	fmt.Println("Pop:  ", stack.Pop())
	fmt.Println("Size: ", stack.Size())
	fmt.Println("Data: ", stack.String())
}
