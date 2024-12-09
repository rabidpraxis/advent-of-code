package utils

type Set[T comparable] struct {
	set map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]bool)}
}

func (set *Set[T]) Add(s T) {
	set.set[s] = true
}

func (set *Set[T]) AddMany(s []T) {
	for _, v := range s {
		set.Add(v)
	}
}

func (set *Set[T]) Has(s T) bool {
	_, found := set.set[s]
	return found
}

func (set *Set[T]) Length() int {
	return len(set.set)
}

func (set *Set[T]) ToSlice() []T {
	var o []T
	for k, _ := range set.set {
		o = append(o, k)
	}
	return o
}

func (set *Set[T]) Subset(other *Set[T]) bool {
	for k, _ := range other.set {
		_, found := set.set[k]
		if !found {
			return false
		}
	}

	return true
}
