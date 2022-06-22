package main

import "fmt"

type BinaryTree struct {
	Value interface{}
	Root  *Node
}

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

func Insert(temp *Node, value interface{}) {
	if temp == nil {
		temp = &Node{Value: value}
		return
	}

	q := []*Node{temp}

	for len(q) != 0 {
		temp := q[len(q)-1]
		q = q[:len(q)-1]

		if temp.Left == nil {
			temp.Left = &Node{Value: value}
			break
		} else {
			q = append(q, temp.Left)
		}

		if temp.Right == nil {
			temp.Right = &Node{Value: value}
			break
		} else {
			q = append(q, temp.Right)
		}
	}
}

func Print(root *Node) {
	q := []*Node{root}

	for len(q) != 0 {
		temp := *q[len(q)-1]
		q = q[:len(q)-1]
		if temp.Right != nil {
			q = append(q, temp.Right)
		}
		if temp.Left != nil {
			q = append(q, temp.Left)
		}
		fmt.Println(temp.Value)
	}
}

func DeleteDepest(root *Node, delNode *Node) {
	q := []*Node{root}

	var temp *Node

	for len(q) > 0 {
		temp = q[0]
		q = q[1:]

		if temp == delNode {
			temp = nil
			return
		}

		if temp.Right != nil {
			if temp.Right == delNode {
				temp.Right = nil
				return
			}
		} else {
			q = append(q, temp.Right)
		}

		if temp.Left != nil {
			if temp.Left == delNode {
				temp.Left = nil
				return
			}
		} else {
			q = append(q, temp.Left)
		}
	}
}

func Delete(root *Node, value interface{}) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Value == value {
			root = nil
		}
		return
	}

	q := []*Node{root}

	var temp, keyNode *Node

	for len(q) > 0 {
		temp = q[0]
		q = q[1:]

		if temp.Value == value {
			keyNode = temp
		}

		if temp.Left != nil {
			q = append(q, temp.Left)
		}

		if temp.Right != nil {
			q = append(q, temp.Right)
		}
	}

	if keyNode != nil {
		value := temp.Value
		DeleteDepest(root, temp)
		keyNode.Value = value
	}
}

func main() {
	tree := BinaryTree{Value: 1, Root: &Node{Value: 5}}
	node2 := Node{Value: 2}
	node3 := Node{Value: 3}
	node4 := Node{Value: 4}
	node5 := Node{Value: 5}
	node6 := Node{Value: 6}
	node7 := Node{Value: 7}
	tree.Root.Left = &node2
	tree.Root.Right = &node3
	node2.Left = &node4
	node2.Right = &node5
	node3.Left = &node6
	node4.Right = &node7
	Insert(tree.Root, 8)
	Delete(tree.Root, 5)
	Print(tree.Root)
}
