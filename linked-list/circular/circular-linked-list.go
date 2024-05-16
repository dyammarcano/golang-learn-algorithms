package main

import "fmt"

// Circular Linked List

type CLLNode struct {
	Value int
	Next  *CLLNode
}

type LinkedList struct {
	Head *CLLNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(value int) {
	node := &CLLNode{Value: value}
	if l.Head == nil {
		l.Head = node
		node.Next = node
		return
	}
	current := l.Head
	for current.Next != l.Head {
		current = current.Next
	}
	current.Next = node
	node.Next = l.Head
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
	for current.Next != l.Head {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) Print() {
	if l.Head == nil {
		return
	}
	current := l.Head
	for {
		fmt.Println(current.Value)
		current = current.Next
		if current == l.Head {
			break
		}
	}
}

func (l *LinkedList) Length() int {
	if l.Head == nil {
		return 0
	}
	length := 1
	current := l.Head
	for current.Next != l.Head {
		length++
		current = current.Next
	}
	return length
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

	// Print the linked list
	list.Print()

	// Get the length of the linked list
	fmt.Println("Length:", list.Length())

	// Get the value at index 1
	fmt.Println("Value at index 1:", list.Get(1))

	// Remove an element from the linked list
	list.Remove(2)

	// Print the linked list
	list.Print()
}
