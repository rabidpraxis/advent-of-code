package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func part1(lines []string) {
	max := 0
	curr := 0

	for _, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			if curr >= max {
				max = curr
			}

			curr = 0
		} else {
			curr += num
		}
	}

	fmt.Println(max)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func part2(lines []string) {
	curr := 0
	var elves []int

	for _, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			elves = append(elves, curr)
			curr = 0
		} else {
			curr += num
		}
	}

	elves = append(elves, curr)

	sort.Ints(elves)
	fmt.Println(sum(elves[len(elves)-3:]))
}

func main() {
	lines := utils.FileLines(os.Args[1])

	part1(lines)
	part2(lines)
}
