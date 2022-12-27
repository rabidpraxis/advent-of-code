package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type Coords struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func extract(s string) *Coords {
	v := strings.Split(s, ",")
	g1 := strings.Split(v[0], "-")
	g2 := strings.Split(v[1], "-")

	x1, _ := strconv.Atoi(g1[0])
	y1, _ := strconv.Atoi(g1[1])
	x2, _ := strconv.Atoi(g2[0])
	y2, _ := strconv.Atoi(g2[1])

	return &Coords{x1, y1, x2, y2}
}

func covered(coords *Coords) bool {
	if (coords.x1 <= coords.x2 && coords.y1 >= coords.y2) ||
		(coords.x2 <= coords.x1 && coords.y2 >= coords.y1) {
		return true
	}
	return false
}

func part1(lines []string) {
	ct := 0
	for _, line := range lines {
		coords := extract(line)
		if covered(coords) {
			ct += 1
		}
	}
	fmt.Println(ct)
}

func overlaps(c *Coords) bool {
	if ((c.x1 >= c.x2) && (c.x1 <= c.y2)) ||
		((c.y1 >= c.x2) && (c.y1 <= c.y2)) ||
		((c.x1 >= c.x2) && (c.y1 <= c.y2)) ||
		((c.x2 >= c.x1) && (c.y2 <= c.y1)) {
		return true
	}
	return false
}

func part2(lines []string) {
	ct := 0
	for _, line := range lines {
		coords := extract(line)
		if overlaps(coords) {
			ct += 1
		}
	}
	fmt.Println(ct)
}

func main() {
	lines := utils.FileLines(os.Args[1])

	part1(lines)
	part2(lines)
}
