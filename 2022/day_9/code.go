package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type instruction struct {
	direction string
	mag       int
}

type coord struct {
	x int
	y int
}

type node struct {
	pos     coord
	name    string
	track   bool
	history []coord
}

func (c *node) Move(pos string) {
	switch pos {
	case "U":
		c.pos.y--
	case "D":
		c.pos.y++
	case "R":
		c.pos.x++
	case "L":
		c.pos.x--
	}

	if c.track {
		c.history = append(c.history, c.pos)
	}
}

func (to *node) Follow(from *node) {
	toPosX := to.pos.x
	toPosY := to.pos.y
	fromPosX := from.pos.x
	fromPosY := from.pos.y

	if utils.Abs(fromPosX-toPosX) > 1 || utils.Abs(fromPosY-toPosY) > 1 {
		if toPosX == fromPosX {
			if fromPosY-toPosY > 0 {
				to.pos.y++
			} else {
				to.pos.y--
			}
		} else if toPosY == fromPosY {
			if fromPosX-toPosX > 0 {
				to.pos.x++
			} else {
				to.pos.x--
			}
		} else {
			if fromPosX-toPosX < 0 {
				to.pos.x--
			} else {
				to.pos.x++
			}
			if fromPosY-toPosY < 0 {
				to.pos.y--
			} else {
				to.pos.y++
			}
		}

		if to.track {
			to.history = append(to.history, to.pos)
		}
	}
}

func takeSteps(inst instruction, nodes []*node) {
	for v := 1; v < inst.mag+1; v++ {
		nodes[0].Move(inst.direction)
		for j := 0; j < len(nodes)-1; j++ {
			nodes[j+1].Follow(nodes[j])
		}
	}
}

func lastNodeHistoryLen(nodes []*node) int {
	h := nodes[len(nodes)-1].history
	s := utils.NewSet[coord]()
	s.AddMany(h)
	return s.Length()
}

func part2(lines []string, insts []instruction) {
	var nodes []*node

	for i := 0; i < 9; i++ {
		nodes = append(nodes, &node{
			name: strconv.Itoa(i),
		})
	}

	nodes = append(nodes, &node{
		track:   true,
		name:    "T",
		history: []coord{{0, 0}},
	})

	for _, inst := range insts {
		takeSteps(inst, nodes)
	}

	fmt.Println(lastNodeHistoryLen(nodes))
}

func part1(lines []string, insts []instruction) {
	nodes := []*node{
		&node{
			name: "H",
		},
		&node{
			track:   true,
			name:    "T",
			history: []coord{{0, 0}},
		},
	}

	for _, inst := range insts {
		takeSteps(inst, nodes)
	}

	fmt.Println(lastNodeHistoryLen(nodes))
}

func main() {
	lines := utils.FileLines(os.Args[1])

	var insts []instruction
	for _, line := range lines {
		direction := line[:1]
		mag, _ := strconv.Atoi(line[2:])
		insts = append(insts, instruction{direction, mag})
	}

	part1(lines, insts)
	part2(lines, insts)
}

// func printGrid(nodes []*node, x int, y int) {
// 	grid := make([][]string, y)
// 	for i, _ := range grid {
// 		grid[i] = make([]string, x)
// 	}

// 	minX := 100
// 	minY := 100
// 	for _, node := range nodes {
// 		if node.pos.x < minX {
// 			minX = node.pos.x
// 		}
// 		if node.pos.y < minY {
// 			minY = node.pos.y
// 		}
// 	}

// 	for y := 0; y < len(grid); y++ {
// 		for x := 0; x < len(grid[0]); x++ {
// 			grid[y][x] = "."
// 		}
// 	}

// 	for _, node := range nodes {
// 		grid[node.pos.y-minY][node.pos.x-minX] = node.name
// 	}

// 	for _, v := range grid {
// 		fmt.Println(v)
// 	}
// }
