package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type grid [][]int

func (g *grid) lookup(c utils.Coord) int {
	return (*g)[c.Y][c.X]
}

func (g *grid) neighborCoords(c utils.Coord) []utils.Coord {
	coords := []utils.Coord{}
	if c.X-1 >= 0 {
		coords = append(coords, c.Add(-1, 0))
	}

	if c.X+1 < len((*g)[0]) {
		coords = append(coords, c.Add(1, 0))
	}

	if c.Y+1 < len(*g) {
		coords = append(coords, c.Add(0, 1))
	}

	if c.Y-1 >= 0 {
		coords = append(coords, c.Add(0, -1))
	}

	return coords
}

func letterCode(letter string) int {
	return strings.Index("0abcdefghijklmnopqrstuvwxyz", letter)
}

func charGrid(lines []string) *grid {
	grid := make(grid, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			switch string(char) {
			case "S":
				grid[i][j] = 0
			case "E":
				grid[i][j] = 27
			default:
				grid[i][j] = letterCode(string(char))
			}
		}
	}
	return &grid
}

type nnode struct {
	pos  utils.Coord
	grid *grid
}

func (n nnode) gridVal() int {
	return n.grid.lookup(n.pos)
}

func (n nnode) PathNeighbors() []utils.Pather {
	ps := []utils.Pather{}
	baseEle := n.gridVal()
	for _, c := range n.grid.neighborCoords(n.pos) {
		if baseEle+1 >= n.grid.lookup(c) {
			ps = append(ps, nnode{c, n.grid})
		}
	}
	return ps
}

func (n nnode) PathNeighborCost(to utils.Pather) float64 {
	return 1
}

func (n nnode) PathEstimatedCost(to utils.Pather) float64 {
	tonnode := to.(nnode)
	return float64(n.pos.Distance(tonnode.pos).Mag())
}

func part1(lines []string) {
	grid := charGrid(lines)

	var start, end utils.Coord
	for yi, y := range *grid {
		for xi, v := range y {
			if v == 27 {
				end = utils.Coord{xi, yi}
			} else if v == 0 {
				start = utils.Coord{xi, yi}
			}
		}
	}

	_, cost, _ := utils.Path(nnode{start, grid}, nnode{end, grid})
	fmt.Println(cost)
}

func part2(lines []string) {
	starts := []utils.Coord{}
	grid := charGrid(lines)
	var end utils.Coord

	for yi, y := range *grid {
		for xi, v := range y {
			if v == 27 {
				end = utils.Coord{xi, yi}
			} else if v == 1 {
				starts = append(starts, utils.Coord{xi, yi})
			}
		}
	}

	minCost := 1000.0

	for _, start := range starts {
		_, cost, found := utils.Path(nnode{start, grid}, nnode{end, grid})
		if found && cost < minCost {
			minCost = cost
		}
	}

	fmt.Println(minCost)
}

func main() {
	lines := utils.FileLines(os.Args[1])

	// part1(lines)
	part2(lines)
}
