package utils

type Coord struct {
	X int
	Y int
}

func (c Coord) Add(x, y int) Coord {
	return Coord{c.X + x, c.Y + y}
}
