package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func part1(grid [][]int) {
	outerCt := (len(grid) * 2) + ((len(grid[0]) - 2) * 2)

	visibleCt := 0
	for y := 1; y < len(grid)-1; y++ {
	ToY:
		for x := 1; x < len(grid[0])-1; x++ {
			curr := grid[y][x]
			// fmt.Println(x, y, curr)

			// Xleft
			for xl := 0; xl < x; xl++ {
				if grid[y][xl] >= curr {
					break
				}

				if xl == x-1 {
					// fmt.Println(x, y, curr, "left")
					visibleCt++
					continue ToY
				}
			}

			// Xright
			for xr := x + 1; xr < len(grid[0]); xr++ {
				if grid[y][xr] >= curr {
					break
				}

				if xr == len(grid[0])-1 {
					// fmt.Println(x, y, curr, "right")
					visibleCt++
					continue ToY
				}
			}

			// Ytop
			for yt := 0; yt < y; yt++ {
				if grid[yt][x] >= curr {
					break
				}

				if yt == y-1 {
					// fmt.Println(x, y, curr, "top")
					visibleCt++
					continue ToY
				}
			}
			// Ybottom
			for yb := y + 1; yb < len(grid); yb++ {
				if grid[yb][x] >= curr {
					break
				}

				if yb == len(grid)-1 {
					// fmt.Println(x, y, curr, "bottom")
					visibleCt++
					continue ToY
				}
			}
		}
	}

	fmt.Println(outerCt + visibleCt)
}

func part2(grid [][]int) {
	maxScenicScore := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			// for y := 3; y == 3; y++ {
			// 	for x := 2; x == 2; x++ {
			curr := grid[y][x]
			// fmt.Println(x, y, curr)

			// Xleft
			xlScore := 0
			for xl := x - 1; xl >= 0; xl-- {
				xlScore++

				// fmt.Println(xl, y, "xl")
				if grid[y][xl] >= curr {
					break
				}
			}
			// fmt.Println("xl", xlScore)

			// Xright
			xrScore := 0
			for xr := x + 1; xr < len(grid[0]); xr++ {
				xrScore++

				// fmt.Println(xr, y, "xr")
				if grid[y][xr] >= curr {
					break
				}
			}
			// fmt.Println("xr", xrScore)

			// Ytop
			ytScore := 0
			for yt := y - 1; yt >= 0; yt-- {
				ytScore++

				// fmt.Println(x, yt, "yt")
				if grid[yt][x] >= curr {
					break
				}
			}
			// fmt.Println("yt", ytScore)

			// Ybottom
			ybScore := 0
			for yb := y + 1; yb < len(grid); yb++ {
				ybScore++

				// fmt.Println(x, yb, "yb")
				if grid[yb][x] >= curr {
					break
				}
			}
			// fmt.Println("yb", ybScore)

			scenicScore := xlScore * xrScore * ytScore * ybScore
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println(maxScenicScore)
}

func main() {
	lines := utils.FileLines(os.Args[1])

	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[0]))
	}

	for x := range lines {
		for y, v := range lines[x] {
			n, _ := strconv.Atoi(string(v))
			grid[x][y] = n
		}
	}

	// for _, v := range grid {
	// 	fmt.Println(v)
	// }

	part1(grid)
	part2(grid)
}
