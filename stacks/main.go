package main

import (
	"fmt"

	"github.com/hugovallada/go-demos/stacks/stack"
)

func main() {
	stck := stack.Stack[string]{}
	stck.Add("a")
	stck.Add("b")
	fmt.Println(stck.List())
	stck.Pop()
	fmt.Println(stck.List())
}
