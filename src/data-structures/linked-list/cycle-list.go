package cycleList

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type CycleList struct {
	Head *Node
	Tail *Node
}

func (list *CycleList) Push(value interface{}) {
	newNode := &Node{Value: value}
	if list.Tail == nil {
		list.Tail = newNode
	}

	if list.Head == nil {
		list.Head = newNode
		newNode.Next = newNode
		return
	}

	newNode.Next = list.Head
	list.Head = newNode
	list.Tail.Next = list.Head
}

func (list *CycleList) Append(value interface{}) {
	newNode := &Node{Value: value}

	if list.Head == nil {
		list.Head = newNode
		newNode.Next = newNode
		return
	}

	list.Tail = newNode

	current := list.Head
	for true {
		if current.Next == list.Head {
			current.Next = newNode
			newNode.Next = list.Head
			break
		}
		current = current.Next
	}
}

func (list *CycleList) Pop() {
	if list.Head == nil {
		return
	}

	current := list.Head
	if current.Next == list.Head {
		list.Head = nil
		return
	}

	for true {
		if current.Next.Next == list.Head {
			current.Next = list.Head
			break
		}
		current = current.Next
	}
}

func (list *CycleList) Shift() {
	if list.Head == nil {
		return
	}

	if list.Head == list.Head.Next {
		list.Head = nil
		return
	}

	list.Head = list.Head.Next
	list.Pop()
}

func (list *CycleList) Count() int {
	length := 0

	if list.Head == nil {
		return length

	}

	current := list.Head
	for true {
		length++
		if current.Next == list.Head {
			break
		}
		current = current.Next
	}

	return length
}

func (list *CycleList) NthNode(n int) interface{} {
	current := list.Head

	for i := 0; i <= n; i++ {
		if n == i {
			return current
		}
		if current.Next == list.Head {
			return nil
		}
		current = current.Next
	}

	return nil
}

func (list *CycleList) Print() {
	str := ""
	current := list.Head

	for current != nil {
		str += fmt.Sprintf("%d -> ", current.Value)
		current = current.Next
		if current == list.Head {
			break
		}
	}

	fmt.Println(str)
}
