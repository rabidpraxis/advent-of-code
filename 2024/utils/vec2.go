package utils

// Vec2 represents a 2D vector or point
type Vec2 struct {
	X, Y int
}

// New creates a new Vec2
func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

// Add returns a new Vec2 with x,y added to the components
func (v Vec2) Add(x, y int) Vec2 {
	return Vec2{v.X + x, v.Y + y}
}

// AddVec returns a new Vec2 that is the sum of this vector and another
func (v Vec2) AddVec(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

// Distance returns a Vec2 representing the absolute distance between vectors
func (v Vec2) Distance(other Vec2) Vec2 {
	return Vec2{
		Abs(v.X - other.X),
		Abs(v.Y - other.Y),
	}
}

// Manhattan returns the Manhattan distance (L1 norm) of the vector
func (v Vec2) Manhattan() int {
	return Abs(v.X) + Abs(v.Y)
}

// Move updates the vector's components in place
func (v *Vec2) Move(x, y int) {
	v.X += x
	v.Y += y
}

// Vec2Set represents a set of 2D vectors with bounds
type Vec2Set struct {
	Min    Vec2
	Max    Vec2
	Points []Vec2
}

// NewVec2Set creates a new Vec2Set with initialized bounds
func NewVec2Set() Vec2Set {
	return Vec2Set{
		Min:    Vec2{X: 1<<31 - 1, Y: 1<<31 - 1},
		Max:    Vec2{X: -1 << 31, Y: -1 << 31},
		Points: make([]Vec2, 0),
	}
}

// Add adds a new vector and updates the set bounds
func (s *Vec2Set) Add(v Vec2) {
	s.Points = append(s.Points, v)

	if v.X < s.Min.X {
		s.Min.X = v.X
	}
	if v.X > s.Max.X {
		s.Max.X = v.X
	}
	if v.Y < s.Min.Y {
		s.Min.Y = v.Y
	}
	if v.Y > s.Max.Y {
		s.Max.Y = v.Y
	}
}

// Normalize returns a new Vec2 normalized relative to the set's bounds
func (s *Vec2Set) Normalize(v Vec2) Vec2 {
	return Vec2{
		X: v.X - s.Min.X,
		Y: v.Y - s.Min.Y,
	}
}

// Width returns the width of the set bounds
func (s *Vec2Set) Width() int {
	return (s.Max.X - s.Min.X) + 1
}

// Height returns the height of the set bounds
func (s *Vec2Set) Height() int {
	return (s.Max.Y - s.Min.Y) + 1
}
