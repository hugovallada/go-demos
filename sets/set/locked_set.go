package set

type LockableSet[T comparable] struct {
	elements map[T]Void
	readonly bool
}

func NewLockableSet[T comparable](elements ...T) LockableSet[T] {
	var lockedSet LockableSet[T]
	lockedSet.elements = make(map[T]Void)
	for _, element := range elements {
		lockedSet.elements[element] = Void{}
	}
	lockedSet.readonly = false
	return lockedSet
}

func CreateLockedSet[T comparable](elements []T) LockableSet[T] {
	lck := NewLockableSet[T](elements...)
	lck.Lock()
	return lck
}

func (l *LockableSet[T]) Add(element T) {
	if !l.readonly {
		l.elements[element] = Void{}
	}
}

func (l *LockableSet[T]) Lock() {
	l.readonly = true
}

func (l LockableSet[T]) List() []T {
	keys := make([]T, 0, len(l.elements))
	for k := range l.elements {
		keys = append(keys, k)
	}
	return keys
}
