package main

import (
	"fmt"
	"strconv"
	"strings"
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

func (l *LinkedList) String() string {
	var b strings.Builder
	current := l.Head
	for current != nil {
		b.WriteString(strconv.Itoa(current.Value))
		if current.Next != nil {
			b.WriteString(" -> ")
		}
		current = current.Next
	}
	return b.String()
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

	// Print the linked list
	val := list.String()

	fmt.Println(val)

	// Remove an element from the linked list
	list.Remove(3)

	// Print the linked list
	list.Print()
}
