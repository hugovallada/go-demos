package set

type ImmutableSet[T comparable] struct {
	elements map[T]Void
}

func NewImmutableSet[T comparable](elements ...T) ImmutableSet[T] {
	var set ImmutableSet[T]
	set.elements = make(map[T]Void)
	for _, element := range elements {
		set.elements[element] = Void{}
	}
	return set
}

func (is ImmutableSet[T]) Contains(element T) bool {
	for k := range is.elements {
		if element == k {
			return true
		}
	}
	return false
}

func (is ImmutableSet[T]) List() []T {
	keys := make([]T, 0, len(is.elements))
	for k := range is.elements {
		keys = append(keys, k)
	}
	return keys
}
