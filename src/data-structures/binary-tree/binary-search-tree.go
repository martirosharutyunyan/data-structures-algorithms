package bTree

import mergeSort "DSA/src/algorithms/sorting-algorithms/merge-sort"

type BST struct {
	Root *Node
}

func InsertRecursive(node *Node, value interface{}) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value.(int) < node.Value.(int) {
		node.Left = InsertRecursive(node.Left, value)
	} else if value.(int) > node.Value.(int) {
		node.Right = InsertRecursive(node.Right, value)
	}

	return node
}

func (bst *BST) InsertIterative(value interface{}) {
	newNode := &Node{Value: value}

	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	var temp, prev *Node = bst.Root, nil

	for temp != nil {
		if temp.Value.(int) > value.(int) {
			prev = temp
			temp = temp.Left
		} else if temp.Value.(int) < value.(int) {
			prev = temp
			temp = temp.Right
		}
	}

	if prev.Value.(int) > value.(int) {
		prev.Left = newNode
	} else {
		prev.Right = newNode
	}
}

func (root *Node) MinValue() interface{} {
	minV := root.Value
	temp := root
	for temp.Left != nil {
		minV = temp.Left.Value
		temp = temp.Left
	}

	return minV
}

func (root *Node) MaxValue() interface{} {
	maxValue := root.Value
	temp := root

	for temp.Right != nil {
		maxValue = temp.Right.Value
		temp = temp.Right
	}

	return maxValue
}

func (root *Node) DeleteNodeIterative(value interface{}) *Node {
	if root == nil {
		return root
	}

	var current, prev *Node = root, nil

	for current != nil && current.Value != value {
		prev = current

		if value.(int) < current.Value.(int) {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	if current == nil {
		return root
	}

	if current.Left == nil || current.Right == nil {
		var newCurrent *Node

		if current.Left == nil {
			newCurrent = current.Right
		} else {
			newCurrent = current.Left
		}

		if prev == nil {
			return newCurrent
		}

		if current == prev.Left {
			prev.Left = newCurrent
		} else {
			prev.Right = newCurrent
		}
	} else {
		var p, temp *Node
		temp = current.Right

		for temp.Left != nil {
			p = temp
			temp = temp.Left
		}

		if p != nil {
			p.Left = temp.Right
		} else {
			current.Right = temp.Right
		}

		current.Value = temp.Value
	}

	return root
}

func DeleteNodeRecursive(root *Node, value interface{}) *Node {
	if root == nil {
		return root
	}

	if value.(int) < root.Value.(int) {
		root.Left = DeleteNodeRecursive(root.Left, value)
	} else if value.(int) > root.Value.(int) {
		root.Right = DeleteNodeRecursive(root.Right, value)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			temp := root.Right
			return temp
		} else if root.Right == nil {
			temp := root.Left
			return temp
		}

		temp := root.Right.MinValue()
		root.Value = temp
		root.Right = DeleteNodeRecursive(root.Right, temp)
	}

	return root
}

func ConstructBSTFromPreorderUtil(preorder []interface{}, preIndex *int, value interface{}, min, max, size int) *Node {
	if *preIndex >= size {
		return nil
	}

	var root *Node

	if value.(int) > min && value.(int) < max {
		root = &Node{Value: value}
		*preIndex++

		if *preIndex < size {
			root.Left = ConstructBSTFromPreorderUtil(preorder, preIndex, preorder[*preIndex], min, value.(int), size)
		}

		if *preIndex < size {
			root.Right = ConstructBSTFromPreorderUtil(preorder, preIndex, preorder[*preIndex], value.(int), max, size)
		}
	}

	return root
}

func ConstructBSTFromPreorder(preorder []interface{}) *Node {
	preIndex := 0

	return ConstructBSTFromPreorderUtil(preorder, &preIndex, preorder[0], 0, 1000, len(preorder))
}

func ToIntArray(array []interface{}) []int {
	temp := []int{}
	for _, value := range array {
		temp = append(temp, value.(int))
	}

	return temp
}

func ToInterfaceArray(array []int) []interface{} {
	temp := []interface{}{}

	for _, value := range array {
		temp = append(temp, value)
	}

	return temp
}

func ArrayToBST(array []interface{}, root *Node, index *int) *Node {
	if root == nil {
		return root
	}
	ArrayToBST(array, root.Left, index)

	root.Value = array[*index]
	*index++

	ArrayToBST(array, root.Right, index)

	return root
}

func ArraySortInPreorder(array []interface{}) []interface{} {
	result := []interface{}{}
	queue := [][]interface{}{array}

	for len(queue) > 0 {
		partition := queue[0]
		queue = queue[1:]
		mid := (len(partition) - 1) / 2
		result = append(result, partition[mid])
		leftPartition := partition[:mid]
		rightPartition := partition[:mid+1]
		if len(leftPartition) > 0 {
			queue = append(queue, leftPartition)
		}

		if len(rightPartition) > 0 {
			queue = append(queue, rightPartition)
		}
	}

	return result
}

func BTreeToBST(root *Node) *BST {
	temp := InorderIterative(root)

	sortedArray := mergeSort.MergeIterative(ToIntArray(temp))

	bst := &BST{}
	for _, value := range sortedArray {
		bst.InsertIterative(value)
	}

	return bst
}

func KthLargestElementBST(root *Node, k int) interface{} {
	temp := InorderIterative(root)

	return temp[k]
}
