# Linked list

## Description

A linked list is a data structure that consists of a sequence of elements, each of which points to the next element in the sequence. Each element in a linked list is called a node. A node consists of two parts: data and a reference to the next node in the sequence.

## How it works

````mermaid
flowchart LR
    A[Data 1] --> B[Data 2] {Next}
    B --> C[Data 3] {Next}
    C --> D[Data 4] {Next}
    D --> E[Data 5] {Next}
    E --> F[Nil]
````

## Comparison of Linked Lists with Arrays & Dynamic Arrays

| Operation | Linked List | Array | Dynamic Array |
| --- | --- | --- | --- |
| Access | O(n) | O(1) | O(1) |
| Insertion/Deletion at beginning | O(1) | O(n) | O(n) |
| Insertion/Deletion at end | O(n) | O(1) | O(1) |
| Insertion/Deletion in middle | O(n) | O(n) | O(n) |
| Size | O(n) | O(n) | O(n) |

Based on the comparison, linked lists are efficient for insertion and deletion operations but inefficient for access operations. Arrays are efficient for access operations but inefficient for insertion and deletion operations. Dynamic arrays provide a balance between access and insertion/deletion operations.

## Inside Code

The provided code is a simple implementation of a singly linked list in Go. A linked list is a linear data structure where each element is a separate object. Each element (or node) of a list consists of two items: the data and a reference to the next node.

The Node struct is defined with two fields: Value, which holds the integer data, and Next, which is a pointer to the next node in the linked list.

```go
type SLLNode struct {
    Value int
    Next  *SLLNode
}
```

The LinkedList struct is defined with a single field Head, which is a pointer to the first node in the linked list.

```go
type LinkedList struct {
    Head *SLLNode
}
```

The NewLinkedList function is a constructor that initializes an empty linked list with the head as nil.

```go
func NewLinkedList() *LinkedList {
    return &LinkedList{nil}
}
```

The Add method on the LinkedList struct adds a new node with a given value at the end of the list. If the list is empty (i.e., the head is nil), it sets the new node as the head. Otherwise, it iterates through the list to find the last node and sets its Next field to the new node.

```go
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
```

The Remove method removes the first node in the list that has a given value. It handles three cases: when the list is empty, when the head node has the given value, and when a node other than the head has the given value.

```go
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
```

The Print method prints the values of all nodes in the list from the head to the end.

```go
func (l *LinkedList) Print() {
    current := l.Head
    for current != nil {
        fmt.Println(current.Value)
        current = current.Next
    }
}
```

For practical use you can replace the `Print()` method with a `Get(index int)` method that returns a value at a given index.

```go
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
```

## Usage

To use the linked list implementation, you can create a new linked list with the `NewLinkedList` function and then add or remove nodes using the `Add` and `Remove` methods. You can print the contents of the list using the `Print` method or convert the list to a string using the `String` method.

The code implementing the complete linked list is in the linked-list.go file. You can run the code using the following command:

```bash
go run singly-linked-list.go
```