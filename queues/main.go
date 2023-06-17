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

}
