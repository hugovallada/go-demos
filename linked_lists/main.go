package main

import (
	"fmt"

	"github.com/hugovallada/go-demos/linked-lists/linked"
)

func main() {
	ll := linked.LinkedList[string]{}
	ll.Insert("one")
	ll.Insert("two")
	ll.Insert("three")
	ll.List()
	ll.DeleteFirst()
	ll.List()
	one := ll.Search("one")
	fmt.Println(one.Value)
	ll.Insert("four")
	ll.List()
	ll.Delete("two")
	ll.List()
}
