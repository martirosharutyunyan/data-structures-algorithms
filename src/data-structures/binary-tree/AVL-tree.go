package bTree

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type AVLNode struct {
	Left, Right *AVLNode
	Height      int
	Value       int
}

func GetAVLNodeHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func RightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2
	y.Height = Max(GetAVLNodeHeight(y.Left), GetAVLNodeHeight(y.Right)) + 1
	x.Height = Max(GetAVLNodeHeight(x.Left), GetAVLNodeHeight(x.Right)) + 1

	return x
}

func LeftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = Max(GetAVLNodeHeight(x.Left), GetAVLNodeHeight(x.Right)) + 1
	y.Height = Max(GetAVLNodeHeight(y.Left), GetAVLNodeHeight(y.Right)) + 1

	return y
}

func GetBalance(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return GetAVLNodeHeight(node.Left) - GetAVLNodeHeight(node.Right)
}

func InsertAVL(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return &AVLNode{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = InsertAVL(node.Left, value)
	} else if value > node.Value {
		node.Right = InsertAVL(node.Right, value)
	} else {
		return node
	}

	node.Height = 1 + Max(GetAVLNodeHeight(node.Left), GetAVLNodeHeight(node.Right))

	balance := GetBalance(node)

	if balance > 1 && value < node.Left.Value {
		return RightRotate(node)
	}

	if balance < -1 && value > node.Right.Value {
		return LeftRotate(node)
	}

	if balance > 1 && value > node.Left.Value {
		node.Left = LeftRotate(node.Left)
		return RightRotate(node)
	}

	if balance < -1 && value < node.Right.Value {
		node.Right = RightRotate(node.Right)
		return LeftRotate(node)
	}

	return node
}

func AVLDeleteNode(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return node
	}

	if value < node.Value {
		node.Left = AVLDeleteNode(node.Left, value)
	} else if value < node.Value {
		node.Right = AVLDeleteNode(node.Right, value)
	} else {
		if node.Left == nil || node.Right == nil {
			var temp *AVLNode
			if node.Left == nil {
				temp = node.Left
			} else {
				temp = node.Right
			}

			if temp == nil {
				temp = node
				node = nil
			} else {
				*node = *temp
			}
		} else {
			temp := AVLMinNodeValue(node.Right)
			node.Value = temp.Value
			node.Right = AVLDeleteNode(node.Right, temp.Value)
		}
	}

	if node == nil {
		return node
	}

	node.Height = 1 + Max(GetAVLNodeHeight(node.Left), GetAVLNodeHeight(node.Right))

	balance := GetBalance(node)

	if balance > 1 && value < node.Left.Value {
		return RightRotate(node)
	}

	if balance < -1 && value > node.Right.Value {
		return LeftRotate(node)
	}

	if balance > 1 && value > node.Left.Value {
		node.Left = LeftRotate(node.Left)
		return RightRotate(node)
	}

	if balance < -1 && value < node.Right.Value {
		node.Right = RightRotate(node.Right)
		return LeftRotate(node)
	}

	return node
}

func AVLMinNodeValue(node *AVLNode) *AVLNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}

	return current
}

func AVLInorderIterative(root *AVLNode) []int {
	var inorder []int

	if root == nil {
		return inorder
	}

	var stack []*AVLNode

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
