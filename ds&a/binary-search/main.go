package main

import "fmt"

var count int

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}
	if n.key < k {
		return n.right.Search(k)
	} else if n.key > k {
		return n.left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{key: 100}

	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(276))
	fmt.Println(count)

}
