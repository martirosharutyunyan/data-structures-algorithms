package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) Push(value interface{}) {
	newNode := &Node{Value: value}
	if list.Tail == nil {
		list.Tail = newNode
	}
	newNode.Next = list.Head
	list.Head = newNode
}

func (list *LinkedList) Append(value interface{}) {
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

func (list *LinkedList) Pop() {
	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = nil
}

func (list *LinkedList) Shift() {
	list.Head = list.Head.Next
}

func main() {
	var list LinkedList = LinkedList{}
	list.Append(10)
	list.Pop()
	fmt.Printf("%+v", list.Head)
}
