package utils

import "maps"

type OccurrenceSet[T comparable] struct {
	set map[T]int
}

func NewOccurrenceSet[T comparable]() *OccurrenceSet[T] {
	oc := OccurrenceSet[T]{}
	oc.set = make(map[T]int)
	return &oc
}

func (set *OccurrenceSet[T]) Clone() *OccurrenceSet[T] {
	return &OccurrenceSet[T]{
		set: maps.Clone(set.set),
	}
}

func (set *OccurrenceSet[T]) Add(s T) {
	i, found := set.set[s]
	if found {
		set.set[s] = i + 1
	} else {
		set.set[s] = 1
	}
}

func (set *OccurrenceSet[T]) AddAll(s []T) {
	for _, v := range s {
		set.Add(v)
	}
}

func (set *OccurrenceSet[T]) AtLeastOccurred(i int) []T {
	var ret []T
	for k, v := range set.set {
		if v >= i {
			ret = append(ret, k)
		}
	}
	return ret
}

func (set *OccurrenceSet[T]) Get(k T) (int, bool) {
	v, err := set.set[k]
	return v, err
}

func (set *OccurrenceSet[T]) MostOccurred() T {
	max := 0
	var most T
	for k, v := range set.set {
		if v > max {
			max = v
			most = k
		}
	}
	return most
}
