package unrolledList

type Node struct {
	Next   *Node
	Values []interface{}
}

type UnrolledList struct {
	Head *Node
}

func (list *UnrolledList) Push(values []interface{}) {
	newNode := &Node{Values: values}
	newNode.Next = list.Head
	list.Head = newNode
}

func (list *UnrolledList) Append(values []interface{}) {
	newNode := &Node{Values: values}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *UnrolledList) Pop() {
	if list.Head == nil {
		return
	}

	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	current := list.Head

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
}

func (list *UnrolledList) Shift() {
	list.Head = list.Head.Next
}

func (list *UnrolledList) Search(value interface{}) (*Node, int, int) {
	current := list.Head

	for current != nil {
		for i := 0; i < len(current.Values); i++ {
			if current.Values[i] == value {
				return current, i, 0
			}
		}
		current = current.Next
	}

	return nil, 0, 0
}
