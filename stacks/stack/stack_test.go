package stack

import "testing"

func TestWhenFunctionAddIsCallShouldAddANewElementInTheEndOfTheStack(t *testing.T) {
	sut := Stack[string]{
		elements: []string{"a", "b", "c", "d", "e", "f"},
	}
	sut.Add("newWorld")
	lastElement := sut.elements[len(sut.elements)-1]
	if lastElement != "newWorld" {
		t.Errorf("expected newWorld, got %s", lastElement)
	}
}

func TestWhenFunctionPopIsCallShouldRemoveTheLastElementOfTheStack(t *testing.T) {
	sut := Stack[string]{
		elements: []string{"a", "b", "c", "d", "e", "f"},
	}
	sut.Pop()
	lastElement := sut.elements[len(sut.elements)-1]
	if lastElement != "e" {
		t.Errorf("expected e, got %s", lastElement)
	}
}
