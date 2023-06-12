package set

type Void struct{}

type Set[T comparable] struct {
	elements map[T]Void
}

func New[T comparable](elements ...T) *Set[T] {
	var set Set[T]
	set.elements = make(map[T]Void)
	for _, element := range elements {
		set.elements[element] = Void{}
	}
	return &set
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = Void{}
}

func (s *Set[T]) AddAll(elements ...T) {
	for _, element := range elements {
		s.elements[element] = Void{}
	}
}

func (s *Set[T]) AddCollection(elements []T) {
	for _, element := range elements {
		s.elements[element] = Void{}
	}
}

func (s Set[T]) List() []T {
	keys := make([]T, 0, len(s.elements))
	for k := range s.elements {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[T]) Contains(element T) bool {
	for k := range s.elements {
		if k == element {
			return true
		}
	}
	return false
}
