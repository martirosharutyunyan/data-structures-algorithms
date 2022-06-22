package linkedList

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyList struct {
	Head *Node
	Tail *Node
}

func (list *SinglyList) Push(value interface{}) {
	newNode := &Node{Value: value}
	if list.Tail == nil {
		list.Tail = newNode
	}
	newNode.Next = list.Head
	list.Head = newNode
}

func (list *SinglyList) Append(value interface{}) {
	newNode := &Node{Value: value}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	list.Tail = newNode

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *SinglyList) Pop() {
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

func (list *SinglyList) Shift() {
	list.Head = list.Head.Next
}

func (list *SinglyList) Count() int {
	length := 0

	current := list.Head
	for current != nil {
		current = current.Next
		length++
	}

	return length
}

func (list *SinglyList) Reverse() {
	var prev, next *Node

	current := list.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev
}

func (list *SinglyList) NthNode(n int) interface{} {
	current := list.Head

	for i := 0; i <= n; i++ {
		if n == i {
			return current
		}
		current = current.Next
	}

	return nil
}

func (list *SinglyList) Print() {
	str := ""
	current := list.Head

	for current != nil {
		str += fmt.Sprintf("%d -> ", current.Value)
		current = current.Next
	}

	fmt.Println(str)
}
