package main

import (
	"fmt"
)

// Singly Linked List

type SLLNode struct {
	Value int
	Next  *SLLNode
}

type LinkedList struct {
	Head *SLLNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(value int) {
	node := &SLLNode{Value: value}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (l *LinkedList) Remove(value int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
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
	list.Add(5)

	// Get the value at index 1
	fmt.Println("Value at index 1:", list.Get(1))

	// Remove an element from the linked list
	list.Remove(3)

	// Print the linked list
	list.Print()
}
