package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func part1(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				fmt.Println(nums[i] * nums[j])
				return
			}
		}
	}
}

func part2(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					fmt.Println(nums[i] * nums[j] * nums[k])
					return
				}
			}
		}
	}
}

func main() {
	input := utils.FileLines(os.Args[1])

	var nums []int
	for _, line := range input {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}

	part1(nums)
	part2(nums)
}
