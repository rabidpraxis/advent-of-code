package utils

type Grid[T any] [][]T

func NewCharGrid(lines []string) Grid[string] {
	grid := make(Grid[string], len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
	}
	return grid
}

func (g *Grid[T]) Lookup(c Coord) T {
	return (*g)[c.Y][c.X]
}
