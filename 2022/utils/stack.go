package utils

// filo

type Stack[T any] []T

func (q *Stack[T]) Len() int {
	return len(*q)
}

func (q *Stack[T]) Push(x T) {
	*q = append(*q, x)
}

func (q *Stack[T]) Pop() T {
	h := *q
	var el T
	l := len(h)
	el, *q = h[l-1], h[0:l-1]
	return el
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}
