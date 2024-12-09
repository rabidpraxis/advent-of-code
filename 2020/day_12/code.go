package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
	"golang.org/x/exp/slices"
)

var (
	dirs = []string{"N", "E", "S", "W"}
)

type instruction struct {
	op  string
	num int
}

type coords struct {
	x int
	y int
}

func nextDir(curDir string, deg int) string {
	idx := slices.Index(dirs, curDir)
	affect := deg / 90
	frn := (idx + affect)
	var newIdx int
	if frn < 0 {
		newIdx = frn + 4
	} else {
		newIdx = frn % 4
	}
	return dirs[newIdx]
}

func updateCoords(dir string, num int, coords *coords) {
	switch dir {
	case "N":
		coords.y += num
	case "S":
		coords.y -= num
	case "E":
		coords.x += num
	case "W":
		coords.x -= num
	}
}

func part1(insts []*instruction) {
	coords := &coords{0, 0}

	dir := "E"

	for _, inst := range insts {
		switch inst.op {
		case "N", "E", "S", "W":
			updateCoords(inst.op, inst.num, coords)
		case "F":
			updateCoords(dir, inst.num, coords)
		case "L":
			dir = nextDir(dir, -inst.num)
		case "R":
			dir = nextDir(dir, inst.num)
		}
	}

	fmt.Println(utils.Abs(coords.x) + utils.Abs(coords.y))
}

func part2(insts []*instruction) {
	wp := &coords{10, 1}
	coords := &coords{0, 0}

	for _, inst := range insts {
		switch inst.op {
		case "N", "E", "S", "W":
			updateCoords(inst.op, inst.num, wp)
		case "F":
			for i := 0; i < inst.num; i++ {
				coords.x += wp.x
				coords.y += wp.y
			}
		case "L":
			for i := 0; i < inst.num/90; i++ {
				nx := wp.x
				wp.x = -wp.y
				wp.y = nx
			}
		case "R":
			for i := 0; i < inst.num/90; i++ {
				nx := wp.x
				wp.x = wp.y
				wp.y = -nx
			}
		}
	}

	fmt.Println(utils.Abs(coords.x) + utils.Abs(coords.y))
}

func main() {
	input := utils.FileLines(os.Args[1])

	var instructions []*instruction
	for _, v := range input {
		op := v[:1]
		num, _ := strconv.Atoi(v[1:])
		instructions = append(instructions, &instruction{op, num})
	}

	part1(instructions)
	part2(instructions)
}
