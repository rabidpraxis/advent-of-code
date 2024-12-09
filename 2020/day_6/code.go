package main

import (
	"fmt"
	"os"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func part1(lines []string) {
	ss := utils.NewSet[string]()
	ct := 0
	for _, line := range lines {
		if line == "" {
			ct += ss.Length()
			ss = utils.NewSet[string]()
			continue
		}

		for _, l := range line {
			ss.Add(string(l))
		}
	}
	ct += ss.Length()
	fmt.Println(ct)
}

func part2(lines []string) {
	oc := utils.NewOccurrenceSet[string]()
	userCt := 0
	ct := 0
	for _, line := range lines {
		if line == "" {
			ct += len(oc.AtLeastOccurred(userCt))
			oc = utils.NewOccurrenceSet[string]()
			userCt = 0
			continue
		}

		for _, l := range line {
			oc.Add(string(l))
		}

		userCt++
	}

	ct += len(oc.AtLeastOccurred(userCt))
	fmt.Println(ct)
}

func main() {
	input := utils.FileLines(os.Args[1])

	part1(input)
	part2(input)
}
