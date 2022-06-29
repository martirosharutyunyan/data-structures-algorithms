package main

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

type Height struct {
	h int
}

func Preorder(root *Node) {
	if root == nil {
		return
	}

	q := []*Node{}

	for root != nil || len(q) != 0 {
		if root != nil {
			fmt.Println(root.Value)
			q = append(q, root)
			root = root.Left
		} else {
			root = q[len(q)-1]
			q = q[:len(q)-1]
			root = root.Right
		}
	}
}

func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Println(root.Value)
	Inorder(root.Right)
}

func Postorder(root *Node) []interface{} {
	if root == nil {
		return nil
	}
	left := Postorder(root.Left)
	right := Postorder(root.Right)

	output := make([]interface{}, 0)

	output = append(output, left...)
	output = append(output, right...)
	output = append(output, root.Value)

	return output
}

func Levelorder(root *Node) {
	q := []*Node{root}
	var node *Node

	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		fmt.Println(node.Value)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
}

func Search(root *Node, data interface{}) int {
	q := []*Node{root}
	var node *Node

	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		if node.Value == data {
			return 1
		}

		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return -1
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

func GetHeight(root *Node) int {
	var x, y int

	if root != nil {
		x = GetHeight(root.Left)
		y = GetHeight(root.Right)

		if x > y {
			return x + 1
		}
		return y + 1
	}

	return 0
}

func DiameterOpt(root *Node, height *Height) int {
	leftHeight := &Height{}
	rightHeight := &Height{}

	if root == nil {
		height.h = 0
		return 0
	}

	leftDiameter := DiameterOpt(root.Left, leftHeight)
	rightDiameter := DiameterOpt(root.Right, rightHeight)
	height.h = int(math.Max(float64(leftDiameter), float64(rightDiameter))) + 1

	return int(math.Max(float64(leftHeight.h)+float64(rightHeight.h), math.Max(float64(leftDiameter), float64(rightDiameter))))
}

func Diameter(root *Node) int {
	height := &Height{}

	return DiameterOpt(root, height)
}

func main() {
	tree := BinaryTree{Root: &Node{Value: 1}}
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
	node3.Right = &node7
	//		 1
	//   2      3
	// 4   5  6   7
	// Preorder(tree.Root)
	// Inorder(tree.Root)
	// Levelorder(tree.Root)
}
