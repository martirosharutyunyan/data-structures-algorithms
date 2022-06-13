package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func (list *List) Push(value interface{}) {
	newNode := &Node{Value: value}
	if list.Tail == nil {
		list.Tail = newNode
	}
	newNode.Next = list.Head
	list.Head = newNode
}

func (list *List) Append(value interface{}) {
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

func (list *List) Pop() {
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

func (list *List) Shift() {
	list.Head = list.Head.Next
}

func (list *List) Count() int {
	length := 0

	current := list.Head
	for current != nil {
		current = current.Next
		length++
	}

	return length
}

func (list *List) Reverse() {
	var prev, next *Node

	current := list.Head

	for current != nil {
		next = current.Next
		prev = next

	}
}

func main() {
	var list List = List{}
	list.Append(10)
	list.Append(10)
	list.Append(10)
	list.Append(10)
	// list.Pop()
	// fmt.Printf("%+v", list.Head)
	fmt.Println(list.Count())
}
