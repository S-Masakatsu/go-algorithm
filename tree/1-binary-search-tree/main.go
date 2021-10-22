// Binary Search Tree (二分探索木)
package main

import "fmt"

type Tree struct {
	root *Node
}

type TreeImpl interface {
	Insert(int)
	Search(int) bool
	Remove(int)
	PrintPreOrder()
	PrintPostOrder()
	PrintInOrder()
}

func NewTree() TreeImpl {
	return &Tree{}
}

func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
	} else {
		t.root.insert(value)
	}
}

func (t *Tree) Search(value int) bool {
	return t.root.search(value)
}

func (t *Tree) Remove(value int) {
	t.root.remove(value)
}

func (t *Tree) PrintPreOrder() {
	t.root.printPreOrder()
	fmt.Println()
}

func (t *Tree) PrintPostOrder() {
	t.root.printPostOrder()
	fmt.Println()
}

func (t *Tree) PrintInOrder() {
	t.root.printInOrder()
	fmt.Println()
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	node := &Node{value: value}
	if value <= n.value {
		if n.left == nil {
			n.left = node
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = node
		} else {
			n.right.insert(value)
		}
	}
}

func (n *Node) search(value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	} else if n.value > value {
		return n.left.search(value)
	} else {
		return n.right.search(value)
	}
}

func (n *Node) remove(value int) *Node {
	if n == nil {
		return n
	}
	if n.value > value {
		n.left = n.left.remove(value)
	} else if n.value < value {
		n.right = n.right.remove(value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		tmp := n.right.min()
		n.value = tmp.value
		n.right = n.right.remove(tmp.value)
	}
	return n
}

func (n *Node) min() *Node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

// root → left → right
func (n *Node) printPreOrder() {
	if n != nil {
		fmt.Printf("%d ", n.value)
		n.left.printPreOrder()
		n.right.printPreOrder()
	}
}

// left → right → root
func (n *Node) printPostOrder() {
	if n != nil {
		n.left.printPostOrder()
		n.right.printPostOrder()
		fmt.Printf("%d ", n.value)
	}
}

// left → root → right
func (n *Node) printInOrder() {
	if n != nil {
		n.left.printInOrder()
		fmt.Printf("%d ", n.value)
		n.right.printInOrder()
	}
}

func main() {
	t := NewTree()
	for _, n := range []int{1, 3, 6, 2, 5, 8, 3, 10, 9, 12} {
		t.Insert(n)
	}
	fmt.Println("Search [7]: ", t.Search(7))
	fmt.Println("Search [5]: ", t.Search(5))
	fmt.Println("Remove 5")
	t.Remove(5)
	fmt.Println("Search [5]: ", t.Search(5))
	fmt.Println("================================")
	fmt.Println("Print functions - Tree Traversal")
	fmt.Print("Pre Order:  ")
	t.PrintPreOrder()
	fmt.Print("Post Order: ")
	t.PrintPostOrder()
	fmt.Print("In Order:   ")
	t.PrintInOrder()
}
