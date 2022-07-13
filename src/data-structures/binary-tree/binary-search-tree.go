// package BST

package main

type Node struct {
	Left, Right *Node
	Value       interface{}
}

type BST struct {
	Root *Node
}

func Insert(node *Node, value interface{}) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value.(int) < node.Value.(int) {
		node.Left = Insert(node.Left, value)
	} else if value.(int) > node.Value.(int) {
		node.Right = Insert(node.Right, value)
	}

	return node
}

func main() {
	bst := &BST{}
	bst.Root = &Node{Value: 1}
	Insert(bst.Root, 2)
	Insert(bst.Root, 3)
	Insert(bst.Root, 4)
}
