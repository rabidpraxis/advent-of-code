package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var part2Shapes = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var shapes = map[string]int{
	"A": 1,
	"X": 1,
	"B": 2,
	"Y": 2,
	"C": 3,
	"Z": 3,
}

func score(a string, b string) int {
	p1 := shapes[a]
	p2 := shapes[b]

	if p1 == p2 {
		return 3
	} else if (p1 == 3 && p2 == 1) || (p1 == 1 && p2 == 2) || (p1 == 2 && p2 == 3) {
		return 0
	} else {
		return 6
	}
}

func part1(lines *bufio.Scanner) int {
	sum := 0
	for lines.Scan() {
		round := strings.Split(lines.Text(), " ")
		sum += score(round[1], round[0]) + shapes[round[1]]
	}

	return sum
}

func part2(lines *bufio.Scanner) int {
	sum := 0
	for lines.Scan() {
		round := strings.Split(lines.Text(), " ")
		shape := shapes[round[0]]
		score := part2Shapes[round[1]]

		// lose
		if score == 0 {
			shape = (shape - 1)
			if shape == 0 {
				shape = 3
			}
			// win
		} else if score == 6 {
			shape = shape + 1
			if shape == 4 {
				shape = 1
			}
		}

		sum += (shape + score)
	}

	return sum
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	fScan := bufio.NewScanner(f)

	// fmt.Println(part1(fScan))
	fmt.Println(part2(fScan))
}
