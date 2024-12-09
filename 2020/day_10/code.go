package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func part1(nums []int) {
	oc := utils.NewOccurrenceSet[int]()
	for i := 0; i < len(nums)-1; i++ {
		oc.Add(nums[i+1] - nums[i])
	}

	one, _ := oc.Get(1)
	three, _ := oc.Get(3)

	fmt.Println(one * three)
}

// [0 1 4 5 6 7 10 11 12 15 16 19 22]
// [ 1 3 1 1 1 3  1  1  3  1  3  3 ]
//   0     4        2      0

// (0), 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, (22)
// (0), 1, 4,    6, 7, 10, 11, 12, 15, 16, 19, (22)
// (0), 1, 4,       7, 10, 11, 12, 15, 16, 19, (22)
// (0), 1, 4, 5,    7, 10, 11, 12, 15, 16, 19, (22)
// (0), 1, 4,       7, 10,     12, 15, 16, 19, (22)

// [1 1 1 1 3 1 1 1 1 3 3 1 1 1 3 1 1 3 3 1 1 1 1 3 1 3 3 1 1 1 1 3]
//      7 				 7   			 4       2         7               7
//
// TBH, I'm unsure why this works, but if you count the number of sequential
// 1's and map them to:
//
// 1 = 0
// 2 = 2
// 3 = 4
// 4 = 7
//
// Then multiply them all together, it gets the correct result
//
func part2(nums []int) {
	oneLen := 0
	var oneLens []int
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff == 1 {
			oneLen++
		} else {
			if oneLen > 1 {
				oneLens = append(oneLens, oneLen)
			}
			oneLen = 0
		}
	}

	mult := 1
	for _, v := range oneLens {
		switch v {
		case 2:
			mult *= 2
		case 3:
			mult *= 4
		case 4:
			mult *= 7
		}
	}

	fmt.Println(mult)
}

func main() {
	input := utils.FileLines(os.Args[1])

	ints := []int{0}
	for _, v := range input {
		n, _ := strconv.Atoi(v)
		ints = append(ints, n)
	}

	sort.Ints(ints)
	ints = append(ints, ints[len(ints)-1]+3)

	part1(ints)
	part2(ints)
}
