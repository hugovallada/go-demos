package queue

import "errors"

type Priority uint8

const (
	HighPriority Priority = 0
	LowPriority  Priority = 1
)

type PriorityQueue[T any] struct {
	low  []T
	high []T
}

func (q *PriorityQueue[T]) Enqueue(element T, priority Priority) {
	switch priority {
	case HighPriority:
		q.high = append(q.high, element)
	default:
		q.low = append(q.low, element)
	}
}

func (q *PriorityQueue[T]) Dequeue() (element T, err error) {
	if len(q.high) != 0 {
		element, q.high = q.high[0], q.high[1:]
		return element, nil
	} else if len(q.low) != 0 {
		element, q.low = q.low[0], q.low[1:]
		return element, nil
	} else {
		return element, errors.New("queue is empty")
	}
}

func (q PriorityQueue[T]) Peek() (element T, err error) {
	if len(q.high) != 0 {
		return q.high[0], nil
	} else if len(q.low) != 0 {
		return q.low[0], nil
	} else {
		return element, errors.New("queue is empty")
	}
}

func (pq PriorityQueue[T]) List() (low []T, high []T) {
	return pq.low, pq.high
}
