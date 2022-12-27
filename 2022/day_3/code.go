package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func letterCode(letter string) int {
	return strings.Index("0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", letter)
}

func findShared(t string) string {
	letterSet := utils.NewSet[string]()
	mid := len(t) / 2

	for idx, r := range t {
		s := string(r)

		if idx >= mid {
			if letterSet.Has(s) {
				return s
			}
		} else {
			letterSet.Add(s)
		}
	}

	return ""
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += letterCode(findShared(line))
	}
	fmt.Println(sum)
}

type OccurrenceSet struct {
	set map[string]int
}

func NewOccurrenceSet() *OccurrenceSet {
	return &OccurrenceSet{make(map[string]int)}
}

func (set *OccurrenceSet) Add(s string) {
	i, found := set.set[s]
	if found {
		set.set[s] = i + 1
	} else {
		set.set[s] = 1
	}
}

func (set *OccurrenceSet) AddAll(s []string) {
	for _, v := range s {
		set.Add(v)
	}
}

func (set *OccurrenceSet) MostOccurred() string {
	max := 0
	var most string
	for k, v := range set.set {
		if v > max {
			max = v
			most = k
		}
	}
	return most
}

func part2(lines []string) {
	occurrenceSet := NewOccurrenceSet()
	sum := 0
	for idx, line := range lines {
		lineSet := utils.NewSet[string]()

		for _, r := range line {
			lineSet.Add(string(r))
		}

		occurrenceSet.AddAll(lineSet.ToSlice())

		if ((idx + 1) % 3) == 0 {
			sum += letterCode(occurrenceSet.MostOccurred())
			occurrenceSet = NewOccurrenceSet()
		}
	}

	fmt.Println(sum)
}

func main() {
	lines := utils.FileLines(os.Args[1])

	part1(lines)
	part2(lines)
}
