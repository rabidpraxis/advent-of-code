package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/2024/utils"
)

func lists() ([]int, []int) {
	lines := utils.FileLines("./day1/input.txt")
	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, line := range lines {
		str := strings.Split(line, "   ")
		l1int, _ := strconv.Atoi(str[0])
		l1 = append(l1, l1int)
		l2int, _ := strconv.Atoi(str[1])
		l2 = append(l2, l2int)
	}
	return l1, l2
}

func part1() {
	l1, l2 := lists()
	slices.Sort(l1)
	slices.Sort(l2)
	total := 0.0
	for idx := range l1 {
		total += math.Abs(float64(l2[idx] - l1[idx]))
	}

	fmt.Println(int64(total))
}

func part2() {
	l1, l2 := lists()
	occurrences := utils.NewOccurrenceSet[int]()
	for _, v := range l2 {
		occurrences.Add(v)
	}

	total := 0
	for _, v := range l1 {
		o, _ := occurrences.Get(v)
		total += v * o
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
