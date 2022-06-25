package linkedList

import (
	"fmt"
)

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
func (list *SinglyList) Search(value interface{}) (*Node, int) {
	current := list.Head

	for current != nil {
		if current.Value == value {
			return current, 0
		}
		current = current.Next
	}

	return nil, -1
}

func (list *SinglyList) ToArray() []interface{} {
	array := []interface{}{}

	current := list.Head

	for current != nil {
		array = append(array, current.Value)
		current = current.Next
	}

	return array
}

func FromArray(array []interface{}) *SinglyList {
	head := &Node{Value: array[0]}
	list := SinglyList{}
	list.Head = head

	current := list.Head

	for i := 1; i < len(array); i++ {
		current.Next = &Node{Value: array[i]}
		current = current.Next
	}

	return &list
}

func GetMiddle(head *Node) *Node {
	if head == nil {
		return head
	}

	var slow, fast = head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func SortedMerge(left *Node, right *Node) *Node {
	var result *Node

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	if left.Value.(int) <= right.Value.(int) {
		result = left
		result.Next = SortedMerge(left.Next, right)
	} else {
		result = right
		result.Next = SortedMerge(left, right.Next)
	}

	return result
}

func MergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	middle := GetMiddle(head)
	nextToMiddle := middle.Next

	middle.Next = nil

	left := MergeSort(head)
	right := MergeSort(nextToMiddle)

	listHead := SortedMerge(left, right)

	return listHead
}
