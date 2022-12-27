package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type pos struct {
	x int
	y int
}

func (p pos) add(x, y int) pos {
	return pos{p.x + x, p.y + y}
}

func (p *pos) move(x, y int) {
	p.x += x
	p.y += y
}

type line struct {
	verts []pos
}

type posset struct {
	min     pos
	max     pos
	xBuffer int
	yBuffer int
}

func (p *posset) update(r pos) {
	if r.x < p.min.x {
		p.min.x = r.x - p.xBuffer
	}
	if r.x > p.max.x {
		p.max.x = r.x + p.xBuffer
	}
	if r.y < p.min.y {
		p.min.y = r.y
	}
	if r.y > p.max.y {
		p.max.y = r.y + p.yBuffer
	}
}

func (p *posset) width() int {
	return (p.max.x - p.min.x) + 1
}

func (p *posset) height() int {
	return (p.max.y - p.min.y) + 1
}

func (p *posset) normalize(v pos) pos {
	v.x = v.x - p.min.x
	v.y = v.y - p.min.y
	return v
}

type grid [][]string

func (g *grid) at(p pos) string {
	return (*g)[p.y][p.x]
}

func (g *grid) set(p *pos, v string) {
	(*g)[p.y][p.x] = v
}

func (g *grid) printLines() {
	for _, l := range *g {
		for _, c := range l {
			if c == "" {
				fmt.Printf(".")
			} else {
				fmt.Printf("%v", c)
			}
		}
		fmt.Println()
	}
}

func (g *grid) drawLines(l []pos) {
	for i := 0; i < len(l)-1; i++ {
		la := l[i]
		lb := l[i+1]

		// (*g)[la.y-1][la.x-1] = "#"

		// fmt.Println(la, lb)
		if diff := la.x - lb.x; diff != 0 {
			// fmt.Println("x", diff)
			if diff < 0 {
				for j := 0; j >= diff; j-- {
					nx := la.x - j
					// fmt.Println("less x", nx, la.y)
					(*g)[la.y][nx] = "#"
				}
			} else {
				for j := 0; j <= diff; j++ {
					nx := la.x - j
					// fmt.Println("more x", nx, la.y)
					(*g)[la.y][nx] = "#"
				}
			}
		} else {
			diff := la.y - lb.y
			// fmt.Println("y", diff)
			if diff < 0 {
				for j := 0; j >= diff; j-- {
					ny := la.y - j
					// fmt.Println("less y", la.x, ny)
					(*g)[ny][la.x] = "#"
				}
			} else {
				for j := 0; j <= diff; j++ {
					ny := la.y - j
					// fmt.Println("more y", la.x, ny)
					(*g)[ny][la.x] = "#"
				}
			}
		}
	}
}

func makeGrid(width int, height int) grid {
	g := make([][]string, height)
	for i, _ := range g {
		g[i] = make([]string, width)
	}
	return g
}

func part1(paths []line) {
	drop := pos{500, 0}
	minmax := posset{
		min:     drop,
		max:     drop,
		xBuffer: 1,
		yBuffer: 1,
	}

	for _, r := range paths {
		for _, v := range r.verts {
			minmax.update(v)
		}
	}

	drop = minmax.normalize(drop)
	grid := makeGrid(minmax.width(), minmax.height())

	for _, r := range paths {
		normVerts := []pos{}
		for _, v := range r.verts {
			normVerts = append(normVerts, minmax.normalize(v))
		}

		grid.drawLines(normVerts)
	}
	grid[drop.y][drop.x] = "+"
	maxY := minmax.normalize(minmax.max).y

	i := 0
	for {
		i++

		grain := &pos{drop.x, drop.y}
		for {
			if grain.y == maxY {
				break
			} else if grid.at(grain.add(0, 1)) == "" {
				grain.move(0, 1)
			} else if grid.at(grain.add(-1, 1)) == "" {
				grain.move(-1, 1)
			} else if grid.at(grain.add(1, 1)) == "" {
				grain.move(1, 1)
			} else {
				break
			}
		}

		if grain.y == maxY {
			fmt.Println("made it", i-1)
			break
		}

		grid.set(grain, "o")
	}

	// grid.printLines()
}

func part2(paths []line) {
	drop := pos{500, 0}
	minmax := posset{
		min:     drop,
		max:     drop,
		xBuffer: 200,
		yBuffer: 1,
	}

	for _, r := range paths {
		for _, v := range r.verts {
			minmax.update(v)
		}
	}

	drop = minmax.normalize(drop)
	grid := makeGrid(minmax.width(), minmax.height())

	for _, r := range paths {
		for j, v := range r.verts {
			r.verts[j] = minmax.normalize(v)
		}

		grid.drawLines(r.verts)
	}
	grid[drop.y][drop.x] = "+"
	maxY := minmax.normalize(minmax.max).y

	i := 0
	for {
		i++

		grain := &pos{drop.x, drop.y}
		for {
			if grain.y == maxY {
				break
			} else if grid.at(grain.add(0, 1)) == "" {
				grain.move(0, 1)
			} else if grid.at(grain.add(-1, 1)) == "" {
				grain.move(-1, 1)
			} else if grid.at(grain.add(1, 1)) == "" {
				grain.move(1, 1)
			} else {
				break
			}
		}

		if grain.x == drop.x && grain.y == drop.y {
			fmt.Println("made it", i)
			break
		}

		grid.set(grain, "o")
	}

	// grid.printLines()
}

func main() {
	slines := utils.FileLines(os.Args[1])

	paths := []line{}
	for _, l := range slines {
		ls := strings.Split(l, " -> ")
		verts := []pos{}
		for _, p := range ls {
			posS := strings.Split(p, ",")
			x, _ := strconv.Atoi(posS[0])
			y, _ := strconv.Atoi(posS[1])
			verts = append(verts, pos{x, y})
		}
		paths = append(paths, line{verts})
	}

	part1(paths)
	part2(paths)
}
