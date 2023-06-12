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
}
