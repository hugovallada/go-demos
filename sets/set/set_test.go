package set

import "testing"

func TestNewSetWithoutGivingElementsCreateANewSetWithoutAnyElements(t *testing.T) {
	newSet := New[string]()
	if len(newSet.elements) > 0 {
		t.Errorf("o valor esperado de elements era 0, mas tinha: %d", len(newSet.elements))
	}
}

func TestNewSetWithElementsCreateANewSetWithElements(t *testing.T) {
	entryArray := []string{"a", "b", "c"}
	newSet := New(entryArray...)
	if len(newSet.elements) != len(entryArray) {
		t.Errorf("o valor esperado de elements era %d, mas tinha: %d", len(entryArray), len(newSet.elements))
	}
}

func TestAddShouldAddANewElement(t *testing.T) {
	newSet := New[string]()
	newSet.Add("mundo")
	if len(newSet.elements) != 1 {
		t.Errorf("numero diferente de elementos esperados: O esperado era 1, mas tinha: %d", len(newSet.elements))
	}
	if newSet.List()[0] != "mundo" {
		t.Errorf("o valor esperado era 'mundo' , mas encontrou '%s'", newSet.List()[0])
	}
}
