package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func part1(lines *bufio.Scanner) int {
	max := 0
	curr := 0

	for lines.Scan() {
		num, err := strconv.Atoi(lines.Text())

		if err != nil {
			if curr >= max {
				max = curr
			}

			curr = 0
		} else {
			curr += num
		}
	}

	return max
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func part2(lines *bufio.Scanner) int {
	curr := 0
	var elves []int

	for lines.Scan() {
		num, err := strconv.Atoi(lines.Text())

		if err != nil {
			elves = append(elves, curr)
			curr = 0
		} else {
			curr += num
		}
	}

	elves = append(elves, curr)

	sort.Ints(elves)
	return sum(elves[len(elves)-3:])
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	fScan := bufio.NewScanner(f)

	fmt.Println(part2(fScan))
}
