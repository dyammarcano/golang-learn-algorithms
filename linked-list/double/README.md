# Double Linked list

* [Introduction](../../README.md)
* [Linked Lists](../README.md)

## How it works

````mermaid
flowchart LR
    A[[Data 1]] <-.-> B[[Data 2]]
    B <-.-> C[[Data 3]]
    C <-.-> D[[Data 4]]
    D <-.-> E[[Data 5]]
    E <-.-> F[[Nil]]
````

## Inside Code

The provided code is a simple implementation of a singly linked list in Go. A linked list is a linear data structure where each element is a separate object. Each element (or node) of a list consists of two items: the data and a reference to the next node.

The Node struct is defined with two fields: Value, which holds the integer data, and Next, which is a pointer to the next node in the linked list.

```go
type DLLNode struct {
    Value int
    Next  *DLLNode
    Prev  *DLLNode
}
```

The LinkedList struct is defined with a single field Head, which is a pointer to the first node in the linked list.

```go
type LinkedList struct {
    Head *DLLNode
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
```

The Remove method removes the first node in the list that has a given value. It handles three cases: when the list is empty, when the head node has the given value, and when a node other than the head has the given value.

```go
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
go run double-linked-list.go
```