package bTree

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value       any
	Left, Right *Node
}

type Height struct {
	h int
}

func PreorderRecursive(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)

	PreorderRecursive(root.Left)
	PreorderRecursive(root.Right)
}

func PreorderIterative(root *Node) []any {
	preorder := []any{}
	current := root
	if root == nil {
		return preorder
	}

	stack := []*Node{}

	for current != nil || len(stack) != 0 {
		if current != nil {
			preorder = append(preorder, current.Value)
			stack = append(stack, current)
			current = current.Left
		} else {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			current = current.Right
		}
	}

	return preorder
}

func InorderRecursive(root *Node) {
	if root == nil {
		return
	}

	InorderRecursive(root.Left)
	fmt.Println(root.Value)
	InorderRecursive(root.Right)
}

func InorderIterative(root *Node) []any {
	inorder := []any{}

	if root == nil {
		return inorder
	}

	stack := []*Node{}

	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		inorder = append(inorder, current.Value)
		current = current.Right
	}

	return inorder
}

func PostorderRecursive(root *Node) {
	if root == nil {
		return
	}

	PostorderRecursive(root.Left)
	PostorderRecursive(root.Right)

	fmt.Println(root.Value)
}

func PostOrderIterative(root *Node) []any {
	postorder := []any{}

	firstStack, secondStack := []*Node{}, []*Node{}

	if root == nil {
		return postorder
	}
	firstStack = append(firstStack, root)

	var temp *Node
	for len(firstStack) > 0 {
		temp = firstStack[len(firstStack)-1]
		firstStack = firstStack[:len(firstStack)-1]
		secondStack = append(secondStack, temp)

		if temp.Left != nil {
			firstStack = append(firstStack, temp.Left)
		}

		if temp.Right != nil {
			firstStack = append(firstStack, temp.Right)
		}
	}

	for len(secondStack) > 0 {
		temp = secondStack[len(secondStack)-1]
		secondStack = secondStack[:len(secondStack)-1]
		postorder = append(postorder, temp.Value)
	}

	return postorder
}

func Levelorder(root *Node) []any {
	levelorder := []any{}

	q := []*Node{root}
	var node *Node

	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		levelorder = append(levelorder, node.Value)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return levelorder
}

func Search(root *Node, data any) int {
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

func Insert(temp *Node, value any) {
	if temp == nil {
		temp = &Node{Value: value}
		return
	}

	stack := []*Node{temp}

	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if temp.Left == nil {
			temp.Left = &Node{Value: value}
			break
		} else {
			stack = append(stack, temp.Left)
		}

		if temp.Right == nil {
			temp.Right = &Node{Value: value}
			break
		} else {
			stack = append(stack, temp.Right)
		}
	}
}

func DeleteDeepest(root *Node, delNode *Node) {
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

func Delete(root *Node, value any) {
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
		DeleteDeepest(root, temp)
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

func IsFullBtree(root *Node) bool {
	q := []*Node{root}
	var node *Node

	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}

		if (node.Right != nil && node.Left == nil) || (node.Right == nil && node.Left != nil) {
			return false
		}
	}

	return true
}

func IntToBool(x int) bool {
	if x == 0 {
		return false
	}
	return true
}

func BoolToInt(boolean bool) int {
	if boolean {
		return 1
	}
	return 0
}

func IsBalanced(root *Node, height *int) int {
	leftHeight, rightHeight := 0, 0

	left, right := 0, 0

	if root == nil {
		*height = 0
		return 1
	}

	left = IsBalanced(root.Left, &leftHeight)
	right = IsBalanced(root.Right, &rightHeight)

	if leftHeight > rightHeight {
		*height = leftHeight
	} else {
		*height = rightHeight
	}
	*height += 1

	if math.Abs(float64(leftHeight)-float64(rightHeight)) >= 2 {
		return 0
	}

	return BoolToInt(IntToBool(left) && IntToBool(right))
}

func TreeFromInorderAndPreorderTraversals(inorder, preorder []any, start, end int, hashmap *map[any]int, preIndex *int) *Node {
	if start > end {
		return nil
	}

	newNode := &Node{Value: preorder[*preIndex]}
	*(preIndex)++

	if start == end {
		return newNode
	}

	inIndex := (*hashmap)[newNode.Value]

	newNode.Left = TreeFromInorderAndPreorderTraversals(inorder, preorder, start, inIndex-1, hashmap, preIndex)
	newNode.Right = TreeFromInorderAndPreorderTraversals(inorder, preorder, inIndex+1, end, hashmap, preIndex)

	return newNode
}

func BuildTreeFromInorderAndPreorderTraversalsWrap(inorder, preorder []any, start, end int) *Node {
	hashmap := make(map[any]int)

	for i := 0; i < len(inorder); i++ {
		hashmap[inorder[i]] = i
	}

	var preIndex int = 0

	return TreeFromInorderAndPreorderTraversals(inorder, preorder, start, end, &hashmap, &preIndex)
}
