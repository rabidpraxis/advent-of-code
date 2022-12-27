package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type vs interface{}

func check(a interface{}, b interface{}, status *status) {
	if status.satisfied {
		return
	}

	fmt.Printf("a %v\nb %v\n\n", a, b)

	switch at := a.(type) {
	case float64:
		switch bt := b.(type) {
		case float64:
			if at == bt {
				return
			}

			fmt.Printf("unmached = %t\n\n", at < bt)

			status.satisfied = true
			status.inOrder = at < bt
			return
		case []interface{}:
			check([]interface{}{a}, b, status)
		}
	case []interface{}:
		switch bt := b.(type) {
		case float64:
			check(a, []interface{}{b}, status)
		case []interface{}:
			if len(at) == 0 && len(bt) > 0 {
				fmt.Println("empty < next = True")
				status.satisfied = true
				status.inOrder = true
				return
			}

			for ati, av := range at {
				if ati >= len(bt) {
					fmt.Println("ati >= len(bt) = False")
					status.satisfied = true
					status.inOrder = false
					return
				}

				check(av, bt[ati], status)

				if status.satisfied {
					return
				}

				if (ati == len(at)-1) && ati < len(bt)-1 {
					fmt.Println("a < b = True")
					status.satisfied = true
					status.inOrder = true
					return
				}
			}
		}
	}
}

func parsePacket(s string) vs {
	var val vs
	if err := json.Unmarshal([]byte(s), &val); err != nil {
		panic(err)
	}
	return val
}

type status struct {
	satisfied bool
	inOrder   bool
}

func part1(lines []string) {
	statuses := []status{}
	for i := 0; i < (len(lines)+1)/3; i++ {
		ni := i * 3
		pa := parsePacket(lines[ni])
		pb := parsePacket(lines[ni+1])

		status := &status{false, false}
		check(pa, pb, status)
		statuses = append(statuses, *status)
	}

	fmt.Println(statuses)

	sum := 0
	for i, v := range statuses {
		if v.inOrder {
			sum += i + 1
		}
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	packets := []interface{}{
		parsePacket("[[2]]"),
		parsePacket("[[6]]"),
	}

	for _, l := range lines {
		if l != "" {
			packets = append(packets, parsePacket(l))
		}
	}

	sort.SliceStable(packets, func(i, j int) bool {
		status := &status{false, false}
		check(packets[i], packets[j], status)
		return status.inOrder
	})

	mult := 0
	for i, v := range packets {
		ret, _ := json.Marshal(v)
		if string(ret) == "[[2]]" {
			mult = i + 1
		} else if string(ret) == "[[6]]" {
			fmt.Println(mult * (i + 1))
		}
	}
}

func main() {
	lines := utils.FileLines(os.Args[1])

	part1(lines)
	part2(lines)
}
