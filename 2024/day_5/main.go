package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/2024/utils"
)

func stuff() ([][]int, [][]int) {
	lines := utils.FileLines(os.Args[1])

	pageOrderRules := [][]int{}
	updates := [][]int{}

	for _, line := range lines {
		if strings.Contains(line, "|") {
			nums := strings.Split(line, "|")
			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])
			pageOrderRules = append(pageOrderRules, []int{n1, n2})
		} else if line != "" {
			nums := strings.Split(line, ",")

			update := []int{}
			for _, n := range nums {
				num, _ := strconv.Atoi(n)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	return pageOrderRules, updates
}

func part1() {
	pageOrderRules, updates := stuff()

	acc := 0
	for _, update := range updates {
		good := true
		for _, rule := range pageOrderRules {
			s1 := slices.Index(update, rule[0])
			s2 := slices.Index(update, rule[1])
			if s1 == -1 || s2 == -1 {
				continue
			} else if s1 > s2 {
				good = false
			}
		}

		if good {
			idx := math.Ceil(float64(len(update) / 2))
			acc += int(update[int(idx)])
		}
	}

	fmt.Println(acc)
}

// move item from atIdx to before beforeIdx
func moveItem(arr []int, beforeIdx, atIdx int) {
	val := arr[atIdx]
	arr = append(arr[:atIdx], arr[atIdx+1:]...)
	arr = append(arr[:beforeIdx], append([]int{val}, arr[beforeIdx:]...)...)
}

func part2() {
	pageOrderRules, updates := stuff()

	acc := 0
	for _, update := range updates {
		good := true
		// Run the rules 3 times (brute force baby)
		for _, rule := range pageOrderRules {
			s1 := slices.Index(update, rule[0])
			s2 := slices.Index(update, rule[1])
			if s1 == -1 || s2 == -1 {
				continue
			} else if s1 > s2 {
				good = false
				moveItem(update, s1, s2)
			}
		}

		for _, rule := range pageOrderRules {
			s1 := slices.Index(update, rule[0])
			s2 := slices.Index(update, rule[1])
			if s1 == -1 || s2 == -1 {
				continue
			} else if s1 > s2 {
				good = false
				moveItem(update, s1, s2)
			}
		}

		for _, rule := range pageOrderRules {
			s1 := slices.Index(update, rule[0])
			s2 := slices.Index(update, rule[1])
			if s1 == -1 || s2 == -1 {
				continue
			} else if s1 > s2 {
				good = false
				moveItem(update, s1, s2)
			}
		}

		if !good {
			idx := math.Ceil(float64(len(update) / 2))
			acc += int(update[int(idx)])
		}
	}

	fmt.Println(acc)
}

func main() {
	part1()
	part2()
}
