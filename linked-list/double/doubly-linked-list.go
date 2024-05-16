package main

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

func (l *LinkedList) String() string {
	var str strings.Builder
	current := l.Head
	for current != nil {
		str.WriteString(strconv.Itoa(current.Value))
		current = current.Next
	}
	return str.String()
}

func main() {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	fmt.Println(l)
	l.Remove(3)
	fmt.Println(l)
	l.Remove(1)
	fmt.Println(l)
	l.Remove(5)
	fmt.Println(l)
	l.Remove(2)
	fmt.Println(l)
	l.Remove(4)
	fmt.Println(l)
}
