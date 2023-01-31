package main

import "fmt"

var count int

// Node represents components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a Node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}
	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Right.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(100)
	tree.Insert(50)
	tree.Insert(600)
	tree.Insert(300)

	fmt.Println(tree)

	fmt.Println(tree.Search(100))
	fmt.Println(count)
}
