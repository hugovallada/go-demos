package main

import (
	"fmt"

	"github.com/hugovallada/go-demos/queues/queue"
)

func main() {
	q := queue.Queue[string]{}
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	fmt.Println(q.List())
	q.Dequeue()
	fmt.Println(q.List())

	pq := queue.PriorityQueue[string]{}
	pq.Enqueue("a", queue.HighPriority)
	pq.Enqueue("b", queue.LowPriority)
	pq.Enqueue("c", queue.HighPriority)
	pq.Enqueue("d", queue.HighPriority)
	pq.Enqueue("e", queue.LowPriority)

	low, high := pq.List()

	fmt.Println("Low:", low)
	fmt.Println("High:", high)

	pq.Dequeue()
	pq.Dequeue()
	pq.Dequeue()

	low, high = pq.List()

	fmt.Println("Low:", low)
	fmt.Println("High:", high)

	pq.Dequeue()
	low, high = pq.List()
	fmt.Println("Low:", low)
	fmt.Println("High:", high)
}
