package set

import (
	"testing"

	"github.com/hugovallada/go-demos/sets/tests"
)

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

func TestAddAllShouldAddAllGivenElements(t *testing.T) {
	newSet := New[string]()
	newSet.AddAll("ola", "mundo", "hello", "world")
	if len(newSet.elements) != 4 {
		t.Errorf("o numero de elementos esperados era 4, mas tinha: %d", len(newSet.elements))
	}
}

func TestAddCollectionShouldAddAllElementsFromACollection(t *testing.T) {
	newSet := New[string]()
	newSet.AddCollection([]string{"ola", "mundo", "hello", "world"})
	if len(newSet.elements) != 4 {
		t.Errorf("o numero de elementos esperados era 4, mas tinha: %d", len(newSet.elements))
	}
}

func TestListReturnAListWithTheSameNumberOfElements(t *testing.T) {
	newSet := New[string]()
	newSet.AddCollection([]string{"ola", "mundo", "hello", "world"})
	listOfElements := newSet.List()

	if len(newSet.elements) != len(listOfElements) {
		t.Errorf("o valor esperado de elements era %d, mas tinha: %d", len(newSet.elements), len(listOfElements))
	}
}

func TestContainsShouldReturnTrueWhenTheElementExistsInTheSet(t *testing.T) {
	newSet := New[string]()
	newSet.AddCollection([]string{"ola", "mundo", "hello", "world"})

	if !newSet.Contains("ola") {
		t.Error("o teste deveria conter a palavra 'ola'")
	}
}

func TestContainsShouldReturnFalseWhenTheElementDoesNotExistsInTheSet(t *testing.T) {
	newSet := New[string]()
	newSet.AddCollection([]string{"ola", "mundo", "hello", "world"})

	if newSet.Contains("hold") {
		t.Error("o teste não deveria conter a palavra 'hold'")
	}
}

func TestWhenAddingNewValuesWithRepeatedValuesIgnoreTheRepeatedOnes(t *testing.T) {
	tc := tests.PrepareTestCase()
	for _, testCase := range tc {
		newSet := New[string]()
		newSet.AddAll(testCase.Entry...)
		list := newSet.List()
		if len(list) != len(testCase.Expected) {
			t.Error("o numero de elementos é diferente do esperado")
		}
	}
}
