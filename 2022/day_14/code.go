package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type line struct {
	verts []utils.Coord
}

func drawLines(g *utils.Grid[string], l []utils.Coord) {
	for i := 0; i < len(l)-1; i++ {
		la := l[i]
		lb := l[i+1]

		if diff := la.X - lb.X; diff != 0 {
			if diff < 0 {
				for j := 0; j >= diff; j-- {
					nx := la.X - j
					(*g)[la.Y][nx] = "#"
				}
			} else {
				for j := 0; j <= diff; j++ {
					nx := la.X - j
					(*g)[la.Y][nx] = "#"
				}
			}
		} else {
			diff := la.Y - lb.Y
			if diff < 0 {
				for j := 0; j >= diff; j-- {
					ny := la.Y - j
					(*g)[ny][la.X] = "#"
				}
			} else {
				for j := 0; j <= diff; j++ {
					ny := la.Y - j
					(*g)[ny][la.X] = "#"
				}
			}
		}
	}
}

func part1(paths []line) {
	drop := utils.Coord{500, 0}
	vertSet := utils.NewCoordColl()

	for _, r := range paths {
		for _, v := range r.verts {
			vertSet.Update(v)
		}
	}

	vertSet.Update(drop)

	vertSet.Max.X = vertSet.Max.X + 1
	vertSet.Min.X = vertSet.Min.X - 1
	vertSet.Max.Y = vertSet.Max.Y + 1

	drop = vertSet.Normalize(drop)
	grid := utils.NewGrid[string](
		vertSet.Width(),
		vertSet.Height(),
	)

	for _, r := range paths {
		normVerts := []utils.Coord{}
		for _, v := range r.verts {
			normVerts = append(normVerts, vertSet.Normalize(v))
		}

		drawLines(grid, normVerts)
	}
	grid.Set(drop, "+")
	maxY := vertSet.Normalize(vertSet.Max).Y

	i := 0
	for {
		i++

		grain := utils.Coord{drop.X, drop.Y}
		for {
			if grain.Y == maxY {
				break
			} else if grid.Lookup(grain.Add(0, 1)) == "" {
				grain.Move(0, 1)
			} else if grid.Lookup(grain.Add(-1, 1)) == "" {
				grain.Move(-1, 1)
			} else if grid.Lookup(grain.Add(1, 1)) == "" {
				grain.Move(1, 1)
			} else {
				break
			}
		}

		if grain.Y == maxY {
			fmt.Println("made it", i-1)
			break
		}

		grid.Set(grain, "o")
	}

	utils.PrintLines(grid)
}

// func part2(paths []line) {
// 	drop := pos{500, 0}
// 	minmax := posset{
// 		min:     drop,
// 		max:     drop,
// 		xBuffer: 200,
// 		yBuffer: 1,
// 	}

// 	for _, r := range paths {
// 		for _, v := range r.verts {
// 			minmax.update(v)
// 		}
// 	}

// 	drop = minmax.normalize(drop)
// 	grid := makeGrid(minmax.width(), minmax.height())

// 	for _, r := range paths {
// 		for j, v := range r.verts {
// 			r.verts[j] = minmax.normalize(v)
// 		}

// 		drawLines(&grid, r.verts)
// 	}
// 	grid[drop.y][drop.x] = "+"
// 	maxY := minmax.normalize(minmax.max).y

// 	i := 0
// 	for {
// 		i++

// 		grain := &pos{drop.x, drop.y}
// 		for {
// 			if grain.y == maxY {
// 				break
// 			} else if grid.at(grain.add(0, 1)) == "" {
// 				grain.move(0, 1)
// 			} else if grid.at(grain.add(-1, 1)) == "" {
// 				grain.move(-1, 1)
// 			} else if grid.at(grain.add(1, 1)) == "" {
// 				grain.move(1, 1)
// 			} else {
// 				break
// 			}
// 		}

// 		if grain.x == drop.x && grain.y == drop.y {
// 			fmt.Println("made it", i)
// 			break
// 		}

// 		grid.set(grain, "o")
// 	}

// 	// grid.printLines()
// }

func main() {
	slines := utils.FileLines(os.Args[1])

	paths := []line{}
	for _, l := range slines {
		ls := strings.Split(l, " -> ")
		verts := []utils.Coord{}
		for _, p := range ls {
			posS := strings.Split(p, ",")
			x, _ := strconv.Atoi(posS[0])
			y, _ := strconv.Atoi(posS[1])
			verts = append(verts, utils.Coord{x, y})
		}
		paths = append(paths, line{verts})
	}

	part1(paths)
	// part2(paths)
}
