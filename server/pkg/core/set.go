package core

type SetElement interface {
	string | int
}

type Set[T SetElement] map[T]struct{}

func NewSet[T SetElement]() *Set[T] {
	set := make(Set[T])
	return &set
}

func (set Set[T]) Add(t T) {
	(set)[t] = struct{}{}
}

func (set Set[T]) Remove(t T) {
	delete(set, t)
}

func (set Set[T]) Has(t T) bool {
	_, has := set[t]
	return has
}
