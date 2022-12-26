package utils

// fifo

type Queue[T any] []T

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Push(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Pop() T {
	h := *q
	var el T
	el, *q = h[0], h[1:len(h)]
	return el
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}
