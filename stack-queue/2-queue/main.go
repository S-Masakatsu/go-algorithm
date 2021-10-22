// Queue (キュー)
// 先入れ先出し(FIFO: First In First Out)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Queue struct {
	data []interface{}
	size int
}

type QueueImpl interface {
	Enqueue(interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
	String() string
}

func NewQueue() QueueImpl {
	return &Queue{}
}

func (q *Queue) Enqueue(data interface{}) {
	q.data = append(q.data, data)
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.data[0]
	q.data = q.data[1:]
	q.size--
	return ret
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) String() string {
	return fmt.Sprint(q.data)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	queue := NewQueue()
	for i := 1; i <= 5; i++ {
		queue.Enqueue(rand.Intn(50))
	}
	fmt.Println("Size:   ", queue.Size())
	fmt.Println("Data:   ", queue.String())
	fmt.Println("Dequeue:", queue.Dequeue())
	fmt.Println("Dequeue:", queue.Dequeue())
	fmt.Println("Size:   ", queue.Size())
	fmt.Println("Data:   ", queue.String())
}
