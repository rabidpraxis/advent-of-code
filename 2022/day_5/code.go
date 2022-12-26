package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func extractBoard(lines []string) ([][]string, []string) {
	var boardGrid [][]string
	var boardIdx int

	for idx, line := range lines {
		lineLen := len(line)

		if line[0:3] == " 1 " {
			// Skip past grid footer into instructions
			boardIdx += 3
			break
		}

		var boardRow []string

		for i := 0; i < (lineLen/4)+1; i++ {
			after := (i + 1) * 4
			if after > lineLen {
				after = lineLen
			}

			curLine := line[i*4 : after]

			boardRow = append(boardRow, string(curLine[1]))
		}

		boardIdx = idx
		boardGrid = append(boardGrid, boardRow)
	}

	return boardGrid, lines[boardIdx:len(lines)]
}

func transposeReverse(grid [][]string) []*Row {
	xl := len(grid[0])
	yl := len(grid)

	ngrid := make([]*Row, xl)
	for rowI := range ngrid {
		ngrid[rowI] = &Row{}
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			val := grid[utils.Abs(j-yl+1)][i]
			if val != " " {
				ngrid[i].row = append(ngrid[i].row, val)
			}
		}
	}

	return ngrid
}

type Row struct {
	row []string
}

func (r *Row) TakeReverse(n int) []string {
	toLen := len(r.row)
	toN := toLen - n

	var ret []string
	for i := toLen; i > toN; i-- {
		ret = append(ret, r.row[i-1])
	}

	r.row = r.row[0:toN]
	return ret
}

func (r *Row) Take(n int) []string {
	toLen := len(r.row)
	toN := toLen - n

	ret := r.row[toN:toLen]

	r.row = r.row[0:toN]
	return ret
}

func (r *Row) PushSlice(vs []string) {
	for _, v := range vs {
		r.row = append(r.row, v)
	}
}

func (r *Row) TopLetter() string {
	return r.row[len(r.row)-1]
}

type Instruction struct {
	move int
	from int
	to   int
}

func processInstructions(lines []string) []Instruction {
	var instructions []Instruction
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)

		move, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])

		instructions = append(instructions, Instruction{move, from - 1, to - 1})
	}

	return instructions
}

func part1(lines []string) {
	rowGrid, rest := extractBoard(lines)
	colGrid := transposeReverse(rowGrid)
	instructions := processInstructions(rest)

	for _, inst := range instructions {
		colGrid[inst.to].PushSlice(colGrid[inst.from].TakeReverse(inst.move))
	}

	var b strings.Builder
	for _, grid := range colGrid {
		b.WriteString(grid.TopLetter())
	}

	fmt.Println(b.String())
}

func part2(lines []string) {
	rowGrid, rest := extractBoard(lines)
	colGrid := transposeReverse(rowGrid)
	instructions := processInstructions(rest)

	for _, inst := range instructions {
		colGrid[inst.to].PushSlice(colGrid[inst.from].Take(inst.move))
	}

	var b strings.Builder
	for _, grid := range colGrid {
		b.WriteString(grid.TopLetter())
	}

	fmt.Println(b.String())
}

func main() {
	lines := utils.FileLines(os.Args[1])

	part1(lines)
	part2(lines)
}
