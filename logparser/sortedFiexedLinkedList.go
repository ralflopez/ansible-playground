package main

import "fmt"

// 5 -> 4 -> 3 -> 1
// A singly linked list that can store [capacity] amount of nodes in ascending order where head stores the smallest value value
type SortedFixedLinkedList struct {
	Head     *Node
	Capacity int
	length   int
}

func NewSortedFixedLinkedList(cap int) *SortedFixedLinkedList {
	return &SortedFixedLinkedList{
		Head:     nil,
		Capacity: cap,
		length:   0,
	}
}

func (ll *SortedFixedLinkedList) GetValue() []*Node {
	return getValue(ll.Head, make([]*Node, 0))
}

func getValue(currentNode *Node, value []*Node) []*Node {
	if currentNode == nil {
		return value
	}

	return getValue(currentNode.Next, append(value, currentNode))
}

func (ll *SortedFixedLinkedList) Print() {
	print(ll.Head)
}

func print(head *Node) {
	if head == nil {
		return
	}

	fmt.Println(head.Value)
	print(head.Next)
}

func (ll *SortedFixedLinkedList) Add(value int) {
	ll.Head = addInSortedPosition(ll.Head, nil, value)
	ll.length++

	if ll.length > ll.Capacity {
		ll.RemoveFirst()
	}
}

func addInSortedPosition(head *Node, prev *Node, value int) *Node {
	if head == nil {
		return NewNode(value)
	}

	if head.Next == nil {
		newNode := NewNode(value)
		if value > head.Value {
			head.Next = newNode
			return head
		}

		newNode.Next = head
		return newNode
	}

	// position found
	if head.Value >= value {
		newNode := NewNode(value)
		if prev == nil {
			// first item
			newNode.Next = head
		} else {
			prev.Next = newNode
			newNode.Next = head
		}
		return newNode
	}

	head.Next = addInSortedPosition(head.Next, head, value)
	return head
}

func add(head *Node, value int) *Node {
	if head == nil {
		return NewNode(value)
	}

	if head.Next == nil {
		head.Next = NewNode(value)
		return head
	}

	head.Next = add(head.Next, value)
	return head
}

func (ll *SortedFixedLinkedList) RemoveLast() {
	ll.Head = removeLast(ll.Head)
	ll.length--
}

func removeLast(head *Node) *Node {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	head.Next = removeLast(head.Next)
	return head
}

func (ll *SortedFixedLinkedList) RemoveFirst() {
	ll.Head = removeFirst(ll.Head)
	ll.length--
}

func removeFirst(head *Node) *Node {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	newHead := head.Next
	head.Next = nil
	return newHead
}

func (ll *SortedFixedLinkedList) IsFullCapacity() bool {
	return ll.length >= ll.Capacity
}

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func main() {
	ll := NewSortedFixedLinkedList(5)

	ll.Add(5)
	ll.Add(10)
	ll.Add(6)
	ll.Add(100)
	ll.Add(1)
	ll.Add(0)
	ll.Add(101)
	//ll.A dd(7)
	//ll.Add(130)
	//ll.Add(150)

	fmt.Println(len(ll.GetValue()))
}
