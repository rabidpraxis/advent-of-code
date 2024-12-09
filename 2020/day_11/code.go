package main

import (
	"fmt"
	"os"

	"github.com/rabidpraxis/advent-of-code/utils"
)

var (
	input    = utils.FileLines(os.Args[1])
	xLen     = len(input[0])
	yLen     = len(input)
	occupied = "#"
	empty    = "L"
)

func newGrid() [][]string {
	grid := make([][]string, yLen)
	for i := range grid {
		grid[i] = make([]string, xLen)
	}

	return grid
}

func printGrid(grid [][]string) {
	for _, v := range grid {
		fmt.Println(v)
	}
}

type iterConfig struct {
	follow     bool
	occupiedCt int
}

func searchDirection(grid [][]string, x int, y int, offsetX int, offsetY int, conf iterConfig) (string, bool) {
	nx := x + offsetX
	ny := y + offsetY

	if nx < 0 || nx > len(grid[0])-1 {
		return "", false
	}

	if ny < 0 || ny > len(grid)-1 {
		return "", false
	}

	curr := grid[ny][nx]

	if conf.follow {
		for {
			if nx < 0 || nx > len(grid[0])-1 {
				return "", false
			}

			if ny < 0 || ny > len(grid)-1 {
				return "", false
			}

			curr = grid[ny][nx]

			if curr != "." {
				return curr, true
			}

			nx += offsetX
			ny += offsetY
		}
	} else {
		return curr, true
	}
}

func iterate(grid [][]string, conf iterConfig) ([][]string, int) {
	updatedGrid := newGrid()
	changes := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			curr := grid[y][x]

			if curr == "." {
				updatedGrid[y][x] = curr
				continue
			}

			occupiedCt := 0

			l, f := searchDirection(grid, x, y, -1, 0, conf)
			if f && l == occupied {
				occupiedCt++
			}
			r, f := searchDirection(grid, x, y, 1, 0, conf)
			if f && r == occupied {
				occupiedCt++
			}
			lt, f := searchDirection(grid, x, y, -1, -1, conf)
			if f && lt == occupied {
				occupiedCt++
			}
			lb, f := searchDirection(grid, x, y, -1, 1, conf)
			if f && lb == occupied {
				occupiedCt++
			}
			t, f := searchDirection(grid, x, y, 0, -1, conf)
			if f && t == occupied {
				occupiedCt++
			}
			b, f := searchDirection(grid, x, y, 0, 1, conf)
			if f && b == occupied {
				occupiedCt++
			}
			rt, f := searchDirection(grid, x, y, 1, -1, conf)
			if f && rt == occupied {
				occupiedCt++
			}
			rb, f := searchDirection(grid, x, y, 1, 1, conf)
			if f && rb == occupied {
				occupiedCt++
			}

			// if y > 0 {
			// 	surround = append(surround, grid[y-1][x])

			// 	if x > 0 {
			// 		surround = append(surround, grid[y-1][x-1])
			// 	}
			// 	if x < xLen-1 {
			// 		surround = append(surround, grid[y-1][x+1])
			// 	}
			// }

			// if y < yLen-1 {
			// 	surround = append(surround, grid[y+1][x])

			// 	if x > 0 {
			// 		surround = append(surround, grid[y+1][x-1])
			// 	}
			// 	if x < xLen-1 {
			// 		surround = append(surround, grid[y+1][x+1])
			// 	}
			// }

			// if x > 0 {
			// 	surround = append(surround, grid[y][x-1])
			// }
			// if x < xLen-1 {
			// 	surround = append(surround, grid[y][x+1])
			// }

			// occupied := 0
			// for _, v := range surround {
			// 	if v == "#" {
			// 		occupied++
			// 	}
			// }

			if curr == occupied && occupiedCt >= conf.occupiedCt {
				changes++
				updatedGrid[y][x] = empty
			} else if curr == empty && occupiedCt == 0 {
				changes++
				updatedGrid[y][x] = occupied
			} else {
				updatedGrid[y][x] = curr
			}
		}
	}

	return updatedGrid, changes
}

func run(grid [][]string, conf iterConfig) {
	iteration, changes := iterate(grid, conf)

	for changes > 0 {
		iteration, changes = iterate(iteration, conf)
	}

	occupiedCt := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if iteration[y][x] == occupied {
				occupiedCt++
			}
		}
	}

	fmt.Println(occupiedCt)
}

func part1(grid [][]string) {
	run(grid, iterConfig{
		follow:     false,
		occupiedCt: 4,
	})
}

func part2(grid [][]string) {
	run(grid, iterConfig{
		follow:     true,
		occupiedCt: 5,
	})
}

func main() {
	grid := newGrid()

	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			grid[y][x] = string(input[y][x])
		}
	}

	// grid[0][8] = "#"
	// grid[2][6] = "S"
	// fmt.Println(grid[1][7])
	// printGrid(grid)
	// fmt.Println(searchDirection(grid, 6, 2, 1, -1, true))
	// fmt.Println(searchDirection(grid, 0, 6, -1, 0, true))

	part1(grid)
	part2(grid)
}
