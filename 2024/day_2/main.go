package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/2024/utils"
)

func reports() [][]int {
	lines := utils.FileLines(os.Args[1])

	reports := make([][]int, 0)

	for _, line := range lines {
		nums := strings.Split(line, " ")
		report := make([]int, 0)
		for _, num := range nums {
			num, _ := strconv.Atoi(num)
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	return reports
}

func part1() {
	safeCt := 0
	for _, report := range reports() {
		prev := report[0]
		decreasing := report[0] > report[1]
		okay := true

		for i := 1; i < len(report); i++ {
			diff := report[i] - prev
			absDiff := math.Abs(float64(diff))

			if (decreasing && diff > 0) || (!decreasing && diff < 0) {
				okay = false
				break
			} else if absDiff < 1 || absDiff > 3 {
				okay = false
				break
			}

			prev = report[i]
		}

		if okay {
			safeCt++
		}
	}

	fmt.Println(safeCt)
}

// This was a failed attempt
func part2() {
	safeCt := 0
	for _, report := range reports() {
		prev := report[0]
		decreasing := report[0] > report[1]
		skipped := false
		okay := true

		for i := 1; i < len(report); i++ {
			diff := report[i] - prev
			absDiff := math.Abs(float64(diff))

			if (decreasing && diff > 0) || (!decreasing && diff < 0) {
				if skipped {
					okay = false
				}
				skipped = true
				continue
			} else if absDiff < 1 || absDiff > 3 {
				if skipped {
					okay = false
				}
				skipped = true
				continue
			}

			prev = report[i]
		}

		if okay {
			safeCt++
		}
	}

	fmt.Println(safeCt)
}

func okay(report []int) bool {
	prev := report[0]
	decreasing := report[0] > report[1]

	for i := 1; i < len(report); i++ {
		diff := report[i] - prev
		absDiff := math.Abs(float64(diff))

		if (decreasing && diff > 0) || (!decreasing && diff < 0) {
			return false
		} else if absDiff < 1 || absDiff > 3 {
			return false
		}

		prev = report[i]
	}

	return true
}

func checkReport(report []int) bool {
	if okay(report) {
		return true
	}

	original := make([]int, len(report))
	copy(original, report)

	for i := 0; i < len(report); i++ {
		// Create a fresh copy for each iteration
		temp := make([]int, len(original))
		copy(temp, original)

		// Remove element at index i
		modified := append(temp[:i], temp[i+1:]...)
		if okay(modified) {
			return true
		}
	}

	return false
}

func part2a() {
	safeCt := 0
	for _, report := range reports() {
		if checkReport(report) {
			safeCt++
		}
	}

	fmt.Println(safeCt)
}

func main() {
	part1()
	part2a()
}
