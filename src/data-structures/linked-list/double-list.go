package linkedList

import "fmt"

type Doubly struct {
	Head *Node
}

func NewDoubly() *Doubly {
	return &Doubly{nil}
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value, Prev: nil, Next: nil}
}

func (ll *Doubly) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head

	if ll.Head != nil {
		ll.Head.Prev = n
	}

	ll.Head = n

}

func (ll *Doubly) AddAtEnd(val interface{}) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
	n.Prev = cur
}

func (ll *Doubly) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next

	if ll.Head != nil {
		ll.Head.Prev = nil
	}
	return cur.Value
}

func (ll *Doubly) DelAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Value
	cur.Next = nil
	return retval
}

func (ll *Doubly) Count() interface{} {
	var ctr int = 0

	for cur := ll.Head; cur != nil; cur = cur.Next {
		ctr += 1
	}

	return ctr
}

func (ll *Doubly) Reverse() {
	var Prev, Next *Node
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = Prev
		cur.Prev = Next
		Prev = cur
		cur = Next
	}

	ll.Head = Prev
}

func (ll *Doubly) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Value, " ")
	}

	fmt.Print("\n")
}

func (ll *Doubly) DisplayReverse() {
	if ll.Head == nil {
		return
	}
	var cur *Node
	for cur = ll.Head; cur.Next != nil; cur = cur.Next {
	}

	for ; cur != nil; cur = cur.Prev {
		fmt.Print(cur.Value, " ")
	}

	fmt.Print("\n")
}
