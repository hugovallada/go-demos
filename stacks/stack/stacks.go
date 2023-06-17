package stack

import "errors"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Add(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() error {
	if len(s.elements) == 0 {
		return errors.New("stack is empty")
	}
	s.elements = s.elements[:len(s.elements)-1]
	return nil
}

func (s Stack[T]) List() []T {
	return s.elements
}
