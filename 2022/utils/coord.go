package utils

type Coord struct {
	X int
	Y int
}

func (c Coord) Add(x, y int) Coord {
	return Coord{c.X + x, c.Y + y}
}

func (c Coord) Distance(c2 Coord) Coord {
	return Coord{
		Abs(c.X - c2.X),
		Abs(c.Y - c2.Y),
	}
}

func (c Coord) Mag() int {
	return Abs(c.X) + Abs(c.Y)
}

func (c *Coord) Move(x, y int) {
	c.X += x
	c.Y += y
}

type CoordColl struct {
	Min    Coord
	Max    Coord
	Coords []Coord
}

func NewCoordColl() CoordColl {
	return CoordColl{
		Min:    Coord{9999, 9999},
		Max:    Coord{-9999, -9999},
		Coords: []Coord{},
	}
}

func (c *CoordColl) Update(r Coord) {
	c.Coords = append(c.Coords, r)

	if r.X < c.Min.X {
		c.Min.X = r.X
	}
	if r.X > c.Max.X {
		c.Max.X = r.X
	}
	if r.Y < c.Min.Y {
		c.Min.Y = r.Y
	}
	if r.Y > c.Max.Y {
		c.Max.Y = r.Y
	}
}

func (c *CoordColl) Normalize(v Coord) Coord {
	v.X = v.X - c.Min.X
	v.Y = v.Y - c.Min.Y
	return v
}

func (p *CoordColl) Width() int {
	return (p.Max.X - p.Min.X) + 1
}

func (p *CoordColl) Height() int {
	return (p.Max.Y - p.Min.Y) + 1
}
