package main

import (
	"fmt"
)

// Doubly Linked List

type DLLNode struct {
	Value int
	Next  *DLLNode
	Prev  *DLLNode
}

type LinkedList struct {
	Head *DLLNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(value int) {
	node := &DLLNode{Value: value}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	node.Prev = current
}

func (l *LinkedList) Remove(value int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			current.Next.Prev = current
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *LinkedList) Get(index int) int {
	if l.Head == nil {
		return -1
	}
	if index == 0 {
		return l.Head.Value
	}
	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func main() {
	// Create a new linked list
	list := NewLinkedList()

	// Add elements to the linked list
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	// Print the linked list
	list.Print()

	// Remove an element from the linked list
	list.Remove(3)

	// Get the value at index 1
	fmt.Println("Value at index 1:", list.Get(1))

	// Print the linked list
	list.Print()
}
