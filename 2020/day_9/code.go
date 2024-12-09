package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func part1(nums []int, buf int) int {
	for idx := buf; idx < len(nums); idx++ {
		ss := utils.NewSet[int]()

		for i := idx - buf; i < idx; i++ {
			for j := i + 1; j < idx; j++ {
				ss.Add(nums[i] + nums[j])
			}
		}

		if !ss.Has(nums[idx]) {
			fmt.Println(nums[idx])
			return nums[idx]
		}
	}

	return 0
}

func part2(nums []int, invalid int) {
Outer:
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		min := n
		max := n

		for j := i + 1; j < len(nums); j++ {
			n += nums[j]
			if n == invalid {
				fmt.Println(min + max)
				return
			} else if n > invalid {
				continue Outer
			}

			if nums[j] > max {
				max = nums[j]
			} else if nums[j] < min {
				min = nums[j]
			}
		}
	}
}

func main() {
	lines := utils.FileLines(os.Args[1])
	var nums []int

	for _, v := range lines {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	invalid := part1(nums, 25)
	part2(nums, invalid)
}
