package queue

import "errors"

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (t T, err error) {
	if q.IsEmpty() {
		return t, errors.New("empty Queue")
	}

	var firstElement T
	firstElement, q.elements = q.elements[0], q.elements[1:]
	return firstElement, nil
}

func (q Queue[T]) Length() int {
	return len(q.elements)
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q Queue[T]) Peek() (t T, err error) {
	if q.IsEmpty() {
		return t, errors.New("empty queue")
	}
	return q.elements[0], nil
}

func (q Queue[T]) List() []T {
	return q.elements
}
