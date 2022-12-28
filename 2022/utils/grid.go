package utils

import "fmt"

type Grid[T comparable] [][]T

func NewGrid[T comparable](w, h int) *Grid[T] {
	grid := make(Grid[T], h)
	for i, _ := range grid {
		grid[i] = make([]T, w)
	}
	return &grid
}

func (g *Grid[T]) Lookup(c Coord) T {
	return (*g)[c.Y][c.X]
}

func (g *Grid[T]) Set(p Coord, v T) {
	(*g)[p.Y][p.X] = v
}

func PrintLines(g *Grid[string]) {
	for _, l := range *g {
		for _, c := range l {
			if c == "" {
				fmt.Printf(".")
			} else {
				fmt.Printf("%v", c)
			}
		}
		fmt.Println()
	}
}
