package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

var (
	instructionRe = regexp.MustCompile(`(\w+) (.*)`)
)

type Instruction struct {
	op      string
	val     int
	visited bool
}

func part1(ins []*Instruction) {
	acc := 0
	idx := 0
	for {
		instr := ins[idx]

		if instr.visited {
			fmt.Println(acc)
			return
		}

		switch instr.op {
		case "nop":
			idx++
		case "acc":
			acc += instr.val
			idx++
		case "jmp":
			idx += instr.val
		}

		instr.visited = true
	}
}

func reset(ins []*Instruction) {
	for _, in := range ins {
		in.visited = false
	}
}

func run(swapIdx int, ins []*Instruction) bool {
	insLen := len(ins)
	acc := 0
	idx := 0
	for {
		instr := ins[idx]

		if instr.visited {
			return false
		}

		op := instr.op
		if idx == swapIdx {
			if op == "nop" {
				op = "jmp"
			} else {
				op = "nop"
			}
		}

		switch op {
		case "nop":
			idx++
		case "acc":
			acc += instr.val
			idx++
		case "jmp":
			idx += instr.val
		}

		if idx == insLen {
			fmt.Println(acc)
			return true
		}

		instr.visited = true
	}
}

func part2(ins []*Instruction) {
	reset(ins)

	var changes []int
	for i, in := range ins {
		if in.op == "nop" || in.op == "jmp" {
			changes = append(changes, i)
		}
	}

	for _, idx := range changes {
		if run(idx, ins) {
			return
		} else {
			reset(ins)
		}
	}
}

func main() {
	lines := utils.FileLines(os.Args[1])

	var instructions []*Instruction

	for _, line := range lines {
		match := instructionRe.FindStringSubmatch(line)
		val, _ := strconv.Atoi(match[2])
		instructions = append(instructions, &Instruction{
			op:      match[1],
			val:     val,
			visited: false,
		})
	}

	part1(instructions)
	part2(instructions)
}
