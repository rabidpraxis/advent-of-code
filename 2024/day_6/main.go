package main

import (
	"fmt"
	"os"

	"github.com/rabidpraxis/advent-of-code/2024/utils"
)

type Directions struct {
	Up, Down, Left, Right utils.Vec2
}

var dirs = Directions{
	Up:    utils.Vec2{X: 0, Y: -1},
	Down:  utils.Vec2{X: 0, Y: 1},
	Left:  utils.Vec2{X: -1, Y: 0},
	Right: utils.Vec2{X: 1, Y: 0},
}

func nextDir(d utils.Vec2) utils.Vec2 {
	if d == dirs.Up {
		return dirs.Right
	} else if d == dirs.Right {
		return dirs.Down
	} else if d == dirs.Down {
		return dirs.Left
	} else {
		return dirs.Up
	}
}

func getDirRune(d utils.Vec2) string {
	if d == dirs.Up {
		return "^"
	} else if d == dirs.Right {
		return ">"
	} else if d == dirs.Down {
		return "v"
	} else {
		return "<"
	}
}

func getGridAndStart() (*utils.Grid[string], utils.Vec2) {
	lines := utils.FileLines(os.Args[1])

	grid := utils.NewGrid[string](len(lines[0]), len(lines))
	var start utils.Vec2
	for i, l := range lines {
		for j, c := range l {
			if string(c) == "^" {
				start = utils.NewVec2(j, i)
			}
			grid.Set(utils.NewVec2(j, i), string(c))
		}
	}

	return grid, start
}

func part1() {
	grid, start := getGridAndStart()
	moves := []utils.Vec2{}
	op := dirs.Up
	current := start
	steps := 0
	for {
		next := current.AddVec(op)
		if !grid.InBounds(next) {
			break
		} else if grid.Lookup(next) == "#" {
			op = nextDir(op)
		} else {
			steps++
			grid.Set(next, getDirRune(op))
			moves = append(moves, next)
			current = next
		}
	}

	uniqueMoves := utils.UniqueSlice(moves)
	fmt.Println(len(uniqueMoves) + 1)
}

type Bounce struct {
	Pos utils.Vec2
	Dir utils.Vec2
}

func part2() {
	grid, start := getGridAndStart()

	op := dirs.Up
	current := start
	trips := [][]Bounce{}
	bounces := 0

	for {
		next := current.AddVec(op)
		if !grid.InBounds(next) {
			break
		} else if grid.Lookup(next) == "#" {
			for i := 1; i < 3; i++ {
				nextIdx := len(trips) - i
				if nextIdx < 0 {
					break
				}

				trips[nextIdx] = append(trips[nextIdx], Bounce{Pos: next, Dir: op})
			}
			trips = append(trips, []Bounce{{Pos: next, Dir: op}})
			op = nextDir(op)
			bounces++
		} else {
			current = next
		}
	}

	curr := trips[2]
	for _, v := range curr {
		grid.Set(v.Pos, "+")
	}

	block, pos, dir := utils.Vec2{}, utils.Vec2{}, utils.Vec2{}
	if curr[0].Dir == dirs.Up {
		block = utils.Vec2{
			X: curr[0].Pos.X - 1,
			Y: curr[2].Pos.Y - 1,
		}
		dir = curr[0].Dir
		pos = block.AddVec(dirs.Right)
	} else if curr[0].Dir == dirs.Right {
		block = utils.Vec2{
			X: curr[2].Pos.X + 1,
			Y: curr[0].Pos.Y - 1,
		}
		dir = curr[0].Dir
		pos = block.AddVec(dirs.Down)
	} else if curr[0].Dir == dirs.Down {
		block = utils.Vec2{
			X: curr[0].Pos.X + 1,
			Y: curr[2].Pos.Y + 1,
		}
		dir = curr[0].Dir
		pos = block.AddVec(dirs.Left)
	} else {
		block = utils.Vec2{
			X: curr[2].Pos.X - 1,
			Y: curr[0].Pos.Y + 1,
		}
		dir = curr[0].Dir
		pos = block.AddVec(dirs.Up)
	}

	target := curr[0].Pos
	next := pos.AddVec(dir)
	cool := false
	for {
		if next == target {
			cool = true
			break
		} else if grid.Lookup(next) == "#" {
			break
		}
		next = next.AddVec(dir)
	}

	grid.Set(block, "*")
	grid.Set(pos, "e")
	fmt.Println(cool)
	fmt.Println(block, pos, dir)

	// }

	// loop through all hits and see if we can find intersection points to add
	// which will close any loops
	utils.PrintGridLines(grid)
	// for _, trip := range trips {
	// 	fmt.Println(trip)
	// }
}

func main() {
	// part1()
	part2()
}
