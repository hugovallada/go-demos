package linked

import "fmt"

type LinkedList[T comparable] struct {
	Head *Node[T]
	Size int
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (l *LinkedList[T]) Insert(element T) {
	node := Node[T]{
		Value: element,
		Next:  l.Head,
	}
	l.Head = &node
	l.Size++
}

func (l *LinkedList[T]) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

func (l LinkedList[T]) List() {
	current := l.Head

	for current != nil {
		fmt.Printf("%v\n", current)
		current = current.Next
	}
}

func (l LinkedList[T]) Search(element T) *Node[T] {
	current := l.Head
	for current != nil {
		if current.Value == element {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList[T]) Delete(element T) {
	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Value == element {
			previous.Next = current.Next
		}
		previous = current
		current = current.Next
	}
}
