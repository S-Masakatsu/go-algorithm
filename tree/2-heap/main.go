// Mini Heap
// 小さいものが上に来る木構造
// Max Heapはその逆
//
// Input: 5, 6, 2, 9, 13, 11, 1
// Heap
// Index:  0, 1, 2, 3, 4, 5,  6,  7
// Array: -x, 1, 6, 2, 9, 13, 11, 5
// Parent = index / 2
// Left   = index * 2
// Right  = (index * 2) + 1
// ■ Insert
// 1) root に代入
// 2) left に代入  → left  と root を 比較して小さければ入れ替える
// 3) right に代入 → right と root を 比較して小さければ入れ替える
// 4) left1 をroot (2) を実行
// 5) left1 をroot (3) を実行
// 6) right1をroot (2) を実行
// 7) right1をroot (3) を実行
// ■ Pop
// root を取り除く → 一番小さな値
// 一番後ろのものをrootに持ってくる
// left  と root を 比較して小さければ入れ替える
// right と root を 比較して小さければ入れ替える
package main

import (
	"fmt"
	"math"
)

type MinHeap struct {
	data []int
	size int
}

type HeapImpl interface {
	Push(int)
	Pop() (int, bool)
	HeapData() []int
	Size() int
}

func NewMiniHeap() HeapImpl {
	return &MinHeap{[]int{-1 * math.MaxInt}, 0}
}

func (m *MinHeap) Push(value int) {
	m.data = append(m.data, value)
	m.size++
	m.heapUP(m.size)
}

func (m *MinHeap) Pop() (int, bool) {
	if len(m.data) == 1 {
		return 0, false
	}
	root := m.data[1]
	data := m.pop()
	if len(m.data) == 1 {
		return root, true
	}
	m.data[1] = data
	m.size--
	m.heapDown(1)
	return root, true
}

func (m *MinHeap) HeapData() []int {
	return m.data[1:]
}

func (m *MinHeap) Size() int {
	return m.size
}

func (m *MinHeap) parentIndex(i int) int {
	return i / 2
}

func (m *MinHeap) leftIndex(i int) int {
	return i * 2
}

func (m *MinHeap) rightIndex(i int) int {
	return (i * 2) + 1
}

func (m *MinHeap) pop() int {
	d := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return d
}

func (m *MinHeap) swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *MinHeap) less(i, j int) bool {
	return m.data[i] < m.data[j]
}

func (m *MinHeap) minIndex(i int) int {
	l, r := m.leftIndex(i), m.rightIndex(i)
	if r > m.size {
		return l
	}

	if m.less(l, r) {
		return l
	}
	return r
}

func (m *MinHeap) heapUP(i int) {
	for m.parentIndex(i) > 0 {
		parentIndex := m.parentIndex(i)
		if m.less(i, parentIndex) {
			m.swap(i, parentIndex)
		}
		i = parentIndex
	}
}

func (m *MinHeap) heapDown(i int) {
	for m.leftIndex(i) <= m.size {
		j := m.minIndex(i)
		if !m.less(i, j) {
			m.swap(i, j)
		}
		i = j
	}
}

func main() {
	input := []int{5, 6, 2, 9, 13, 11, 1}
	mh := NewMiniHeap()
	for _, i := range input {
		mh.Push(i)
	}
	fmt.Println("Head Data:", mh.HeapData())
	if v, ok := mh.Pop(); ok {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Head Data:", mh.HeapData())
	fmt.Println("Head Size:", mh.Size())
}
