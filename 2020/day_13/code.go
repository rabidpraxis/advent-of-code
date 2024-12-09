package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
	"golang.org/x/exp/slices"
)

func part1(id int, busses []*bus) {
	minWait := 9999
	waitId := 0
	for _, bus := range busses {
		v := id / bus.id
		wait := (v+1)*bus.id - id
		if wait < minWait {
			minWait = wait
			waitId = bus.id
		}
	}

	fmt.Println(waitId * minWait)
}

type bus struct {
	pos  int
	id   int
	diff int
}

func part2(id int, busses []*bus) {
	slices.SortFunc(busses, func(a, b *bus) bool {
		return a.id > b.id
	})

	bump := busses[0].id
	base := busses[0].pos
	ts := bump

	for _, bus := range busses[1:] {
		bus.diff = (bus.pos - base)
	}

	earlistDiff := 100
	for _, bus := range busses[1:] {
		if bus.diff < earlistDiff {
			earlistDiff = bus.diff
		}
	}

Root:
	for {
		ts += bump

		for _, bus := range busses[1:] {
			if (ts+bus.diff)%bus.id != 0 {
				continue Root
			}
		}

		break
	}

	fmt.Println(ts + earlistDiff)
}

func main() {
	input := utils.FileLines(os.Args[1])

	id, _ := strconv.Atoi(input[0])
	buses := strings.Split(input[1], ",")

	var busses []*bus
	for idx, v := range buses {
		busId, err := strconv.Atoi(v)
		if err == nil {
			busses = append(busses, &bus{idx, busId, 0})
		}
	}

	part1(id, busses)
	part2(id, busses)
}
