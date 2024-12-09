package main

import (
	"fmt"
	"os"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type Ticket struct {
	row    int
	column int
	seatId int
}

func updateRegion(code string, ranges []int, topCode string) {
	mid := ((ranges[1] - ranges[0]) / 2) + ranges[0]
	if code == topCode {
		ranges[0] = mid + 1
	} else {
		ranges[1] = mid
	}
}

func bspFind(extent int, codes string, topCode string) int {
	ranges := []int{0, extent}
	for _, v := range codes {
		updateRegion(string(v), ranges, topCode)
	}
	return ranges[0]
}

func part1(tickets []*Ticket) {
	max := 0
	for _, v := range tickets {
		if v.seatId > max {
			max = v.seatId
		}
	}
	fmt.Println(max)
}

func part2(tickets []*Ticket) {
	seatIds := utils.NewSet[int]()
	for _, v := range tickets {
		seatIds.Add(v.seatId)
	}

	var missing []int
	for i := 0; i < 1024; i++ {
		if !seatIds.Has(i) {
			missing = append(missing, i)
		}
	}

	for _, v := range missing {
		if seatIds.Has(v+1) && seatIds.Has(v-1) {
			fmt.Println("Seat ID:", v)
			return
		}
	}
}

func main() {
	input := utils.FileLines(os.Args[1])

	var tickets []*Ticket
	for _, code := range input {
		row := bspFind(127, code[:7], "B")
		column := bspFind(7, code[7:10], "R")
		ticket := &Ticket{row, column, row*8 + column}

		tickets = append(tickets, ticket)
	}

	part1(tickets)
	part2(tickets)
}
