package main

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

// func (list *CycleList) Shift() {
// 	list.Head = list.Head.Next
// }

// func (list *CycleList) Count() int {
// 	length := 0

// 	current := list.Head
// 	for current != nil {
// 		current = current.Next
// 		length++
// 	}

// 	return length
// }

// func (list *CycleList) Reverse() {
// 	var prev, next *Node

// 	current := list.Head

// 	for current != nil {
// 		next = current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}

// 	list.Head = prev
// }

// func (list *CycleList) NthNode(n int) interface{} {
// 	current := list.Head

// 	for i := 1; i <= n; i++ {
// 		if n == i {
// 			return current
// 		}
// 		current = current.Next
// 	}

// 	return nil
// }

func print(list *CycleList) {
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

// func main() {
// 	cycle := CycleList{}
// 	cycle.Append(6)
// 	cycle.Append(6)
// 	cycle.Append(6)
// 	cycle.Append(6)
// 	cycle.Append(7)

// 	print(&cycle)
// }
