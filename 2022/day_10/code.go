package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type instruction struct {
	op  string
	num int
}

type cycle struct {
	inst     *instruction
	register int
	cycle    int
}

func buildCycles(insts []*instruction) []*cycle {
	cycleVal := 0
	xVal := 1
	var cycles []*cycle

	for _, inst := range insts {
		cycleVal++
		if inst.op == "addx" {
			cycleVal++
			xVal += inst.num
		}

		cycles = append(cycles, &cycle{
			inst:     inst,
			register: xVal,
			cycle:    cycleVal,
		})
	}

	return cycles
}

func buildInstructions(lines []string) []*instruction {
	var insts []*instruction
	for _, line := range lines {
		op := line[:4]
		inst := &instruction{op, 0}

		if op == "addx" {
			num, _ := strconv.Atoi(line[5:])
			inst.num = num
		}

		insts = append(insts, inst)
	}

	return insts
}

func part1(cycles []*cycle) {
	currentSignal := 20
	signalTotal := 0
	for i := 0; i < len(cycles)-1; i++ {
		if cycles[i].cycle < currentSignal && cycles[i+1].cycle >= currentSignal {
			signalTotal += currentSignal * cycles[i].register
			currentSignal += 40
		}
	}
	fmt.Println(signalTotal)
}

func part2(cycles []*cycle) {
	grid := utils.MakeGrid(40, 6)

	spritePos := 1
	cycleIdx := 0
	nextCycle := cycles[0].cycle

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			idx := (y * len(grid[0])) + x

			if idx >= nextCycle {
				nextCycle = cycles[cycleIdx+1].cycle
				spritePos = cycles[cycleIdx].register
				cycleIdx++
			}

			if x >= spritePos-1 && x <= spritePos+1 {
				grid[y][x] = "#"
			} else {
				grid[y][x] = "."
			}
		}
	}

	utils.PrintGrid(grid)
}

func main() {
	lines := utils.FileLines(os.Args[1])
	insts := buildInstructions(lines)
	cycles := buildCycles(insts)

	part1(cycles)
	part2(cycles)
}
