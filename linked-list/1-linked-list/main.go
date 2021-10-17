// Linked List (単方向リスト)
// 要素同士を前後双方向のリンクで参照するリンクリスト
// https://ja.wikipedia.org/wiki/%E9%80%A3%E7%B5%90%E3%83%AA%E3%82%B9%E3%83%88
package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node // root Node
}

func NewLenkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Append(data interface{}) {
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
}

func (l *LinkedList) Insert(data interface{}) {
	node := &Node{data: data}
	node.next = l.head
	l.head = node
}

func (l *LinkedList) Get(index int) *Node {
	current := l.head
	if current == nil {
		return nil
	}

	var ptr *Node = &Node{}
	for i := 0; i < index; i++ {
		ptr = current
		current = current.next
	}
	return ptr
}

func (l *LinkedList) Remove(data interface{}) error {
	current := l.head
	// 先頭に削除する要素がある場合
	if current != nil && current.data == data {
		l.head = current.next
		// メモリから解放する
		current = nil
		return nil
	}

	var prev *Node = &Node{}
	for current != nil && current.data != data {
		prev = current
		current = current.next
	}

	// 削除するデータがない場合
	if current == nil {
		return fmt.Errorf("remove data not found: %v", data)
	}

	prev.next = current.next
	current = nil
	return nil
}

func (l *LinkedList) Print() {
	for cur := l.head; cur != nil; {
		if cur.next != nil {
			fmt.Printf("%v,", cur.data)
		} else {
			fmt.Println(cur.data)
		}
		cur = cur.next
	}
}

func main() {
	list := NewLenkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Insert(0)
	list.Print()
}
