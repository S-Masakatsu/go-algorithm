// Doubly Linked List (双方向連結リスト)
package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node // root Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) Insert(data interface{}) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
		return
	}

	l.head.prev = node
	node.next = l.head
	l.head = node
}

func (l *DoublyLinkedList) Append(data interface{}) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = node
	node.prev = lastNode
}

func (l *DoublyLinkedList) Get(index int) (*Node, bool) {
	current := l.head
	if current == nil {
		return nil, false
	}

	ptr := &Node{}
	for i := 0; i < index; i++ {
		ptr = current
		if current.next == nil {
			return nil, false
		}
		current = current.next
	}
	return ptr, ptr != nil
}

func (l *DoublyLinkedList) Remove(date interface{}) error {
	current := l.head
	if current != nil && current.data == date {
		// Nodeが１つしかない場合
		if current.next == nil {
			l.head = nil
			current = nil
			return nil
		}
		next := current.next
		next.prev = nil
		l.head = next
		current = nil
		return nil
	}

	for current != nil && current.data != date {
		current = current.next
	}

	if current == nil {
		return fmt.Errorf("not found: %v", date)
	}

	// 末尾のNodeを削除する場合
	if current.next == nil {
		prev := current.prev
		prev.next = nil
		current = nil
		return nil
	}

	next := current.next
	prev := current.prev
	prev.next = next
	next.prev = prev
	current = nil
	return nil
}

// 逆順に並び替え(forを使った手法)
func (l *DoublyLinkedList) ReverseIterative() {
	prev := &Node{}
	current := l.head
	for current != nil {
		prev = current.prev
		current.prev = current.next
		current.next = prev
		current = current.prev
	}
	if prev != nil {
		l.head = prev.prev
	}
}

// 逆順に並び替え(再帰を使った手法)
func (l *DoublyLinkedList) ReverseRecursive() {
	var fn func(*Node) *Node
	fn = func(current *Node) *Node {
		if current == nil {
			return nil
		}
		prev := current.prev
		current.prev = current.next
		current.next = prev
		if current.prev == nil {
			return current
		}
		return fn(current.prev)
	}

	l.head = fn(l.head)
}

func (l *DoublyLinkedList) Print() {
	for current := l.head; current != nil; {
		if current.next != nil {
			fmt.Printf("%v,", current.data)
		} else {
			fmt.Println(current.data)
		}
		current = current.next
	}
}

func main() {
	list := NewDoublyLinkedList()
	for i := 1; i <= 5; i++ {
		list.Append(i)
	}
	list.Insert(0)
	list.Print()
	if n, ok := list.Get(3); ok {
		fmt.Println(n.data)
	}
	if n, ok := list.Get(10); ok {
		fmt.Println(n.data)
	}
	list.Remove(3)
	list.Print()
	list.ReverseIterative()
	list.Print()
	list.ReverseRecursive()
	list.Print()
}
