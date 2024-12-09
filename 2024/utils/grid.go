package utils

import "fmt"

type Grid[T comparable] [][]T

func NewGrid[T comparable](w, h int) *Grid[T] {
	grid := make(Grid[T], h)
	for i := range grid {
		grid[i] = make([]T, w)
	}
	return &grid
}

func (g *Grid[T]) Lookup(v Vec2) T {
	return (*g)[v.Y][v.X]
}

func (g *Grid[T]) Set(v Vec2, val T) {
	(*g)[v.Y][v.X] = val
}

// Width returns the width of the grid
func (g *Grid[T]) Width() int {
	if len(*g) == 0 {
		return 0
	}
	return len((*g)[0])
}

// Height returns the height of the grid
func (g *Grid[T]) Height() int {
	return len(*g)
}

// InBounds checks if a vector is within the grid boundaries
func (g *Grid[T]) InBounds(v Vec2) bool {
	return v.X >= 0 && v.X < g.Width() && v.Y >= 0 && v.Y < g.Height()
}

func (g *Grid[T]) LookupAll(start, end Vec2) []T {
	var items []T
	for y := start.Y; y < end.Y; y++ {
		for x := start.X; x < end.X; x++ {
			items = append(items, (*g)[y][x])
		}
	}
	return items
}

// Fill sets all cells in the grid to the given value
func (g *Grid[T]) Fill(val T) {
	for y := range *g {
		for x := range (*g)[y] {
			(*g)[y][x] = val
		}
	}
}

// Copy returns a deep copy of the grid
func (g *Grid[T]) Copy() *Grid[T] {
	newGrid := NewGrid[T](g.Width(), g.Height())
	for y := range *g {
		copy((*newGrid)[y], (*g)[y])
	}
	return newGrid
}

// FindAll returns all positions where the value matches the target
func (g *Grid[T]) FindAll(target T) []Vec2 {
	var positions []Vec2
	for y := range *g {
		for x := range (*g)[y] {
			if (*g)[y][x] == target {
				positions = append(positions, Vec2{X: x, Y: y})
			}
		}
	}
	return positions
}

// Neighbors returns all valid adjacent positions (orthogonal)
func (g *Grid[T]) Neighbors(v Vec2) []Vec2 {
	dirs := []Vec2{
		{X: 0, Y: -1}, // up
		{X: 1, Y: 0},  // right
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
	}

	var valid []Vec2
	for _, dir := range dirs {
		next := v.AddVec(dir)
		if g.InBounds(next) {
			valid = append(valid, next)
		}
	}
	return valid
}

// DiagonalNeighbors returns all valid diagonal positions
func (g *Grid[T]) DiagonalNeighbors(v Vec2) []Vec2 {
	dirs := []Vec2{
		{X: -1, Y: -1}, // up-left
		{X: 1, Y: -1},  // up-right
		{X: -1, Y: 1},  // down-left
		{X: 1, Y: 1},   // down-right
	}

	var valid []Vec2
	for _, dir := range dirs {
		next := v.AddVec(dir)
		if g.InBounds(next) {
			valid = append(valid, next)
		}
	}
	return valid
}

// AllNeighbors returns all valid adjacent positions (including diagonals)
func (g *Grid[T]) AllNeighbors(v Vec2) []Vec2 {
	return append(g.Neighbors(v), g.DiagonalNeighbors(v)...)
}

func PrintGridLines(g *Grid[string]) {
	for _, line := range *g {
		for _, cell := range line {
			if cell == "" {
				fmt.Print(".")
			} else {
				fmt.Printf("%v", cell)
			}
		}
		fmt.Println()
	}
}
