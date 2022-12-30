package utils

import (
	"testing"

	"gotest.tools/assert"
)

type person struct {
	name  string
	score int
}

func TestNeat(t *testing.T) {
	h := NewHeap[person](func(a, b person) bool {
		if a.score > b.score {
			return true
		}
		return false
	})

	h.Push(person{"kevin", 2})
	h.Push(person{"webster", 1})
	h.Push(person{"winner", 100})

	assert.Equal(t, "winner", h.Pop().name)
	assert.Equal(t, "kevin", h.Pop().name)
}
