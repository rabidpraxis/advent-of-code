package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func findMarkerIdx(input string, r int) int {
	inputLen := len(input)
	for i := r; i < inputLen; i++ {
		markerSet := input[i-r : i]
		ss := utils.NewSet[string]()
		ss.AddMany(strings.Split(markerSet, ""))

		if ss.Length() == r {
			return i
		}
	}

	return 0
}

func part1(input string) {
	fmt.Println(findMarkerIdx(input, 4))
}

func part2(input string) {
	fmt.Println(findMarkerIdx(input, 14))
}

func main() {
	input := utils.FileLines(os.Args[1])[0]

	part1(input)
	part2(input)
}
