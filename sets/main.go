package main

import (
	"fmt"

	"github.com/hugovallada/go-demos/sets/set"
)

type Pessoa struct {
	Name string
}

// Exemplo de uso
func main() {
	persons := set.New[Pessoa]()
	p := Pessoa{Name: "Hugo"}
	p1 := Pessoa{Name: "Lopes"}
	p2 := Pessoa{Name: "Adalberto"}
	pessoas := []Pessoa{
		{Name: "Maicon"},
		{Name: "Hugo"},
		{Name: "Maicon"},
	}
	persons.Add(p1)
	persons.AddAll(p, p2)
	persons.AddCollection(pessoas)

	fmt.Println(persons.Contains(p))
	fmt.Println(persons.List())

	numeros := set.New(1, 2, 2, 3, 2, 4, 5, 1, 4, 5, 6, 6, 6, 7)
	fmt.Println(numeros.List())

	pessoasImmutable := set.NewImmutableSet[Pessoa](persons.List()...)
	fmt.Println(pessoasImmutable.List())

	locked := set.NewLockableSet("a", "b", "c", "c", "d", "e", "e")
	fmt.Println(locked.List())
	locked.Add("z")
	fmt.Println(locked.List())
	locked.Lock()
	locked.Add("x")
	fmt.Println(locked.List())

	newLocked := set.CreateLockedSet(numeros.List())
	newLocked.Add(1000)
	fmt.Println(newLocked.List())

}
