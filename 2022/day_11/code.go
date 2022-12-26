package main

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
	"golang.org/x/exp/slices"
)

type MonkeyBase struct {
	id        int
	op        string
	trueId    int
	falseId   int
	inspectCt int
}

type MonkeySmall struct {
	MonkeyBase

	items   []int
	opLeft  *int
	opRight *int
	testDiv int
}

type MonkeyBig struct {
	MonkeyBase

	items   []*big.Int
	opLeft  *big.Int
	opRight *big.Int
	testDiv *big.Int
}

func (m *MonkeyBig) mutateWorry(n *big.Int) *big.Int {
	left := n
	if m.opLeft != nil {
		left = m.opLeft
	}
	right := n
	if m.opRight != nil {
		left = m.opRight
	}
	switch m.op {
	case "*":
		n.Mul(left, right)
	case "+":
		n.Add(left, right)
	}
	return n
}

func (m *MonkeySmall) updateWorry(n int) int {
	left := n
	if m.opLeft != nil {
		left = *m.opLeft
	}
	right := n
	if m.opRight != nil {
		left = *m.opRight
	}

	switch m.op {
	case "*":
		return left * right
	case "+":
		return left + right
	}
	return 0
}

var (
	idRe    = regexp.MustCompile(`Monkey (\d+)`)
	itemsRe = regexp.MustCompile(`  Starting items: (.*)`)
	opRe    = regexp.MustCompile(`  Operation: new = ((?:\w|\d)+?) (.+) (.*)`)
	testRe  = regexp.MustCompile(`  Test: divisible by (\d+)`)
	trueRe  = regexp.MustCompile(`    If true: throw to monkey (\d+)`)
	falseRe = regexp.MustCompile(`    If false: throw to monkey (\d+)`)
)

func buildBigMonkies(lines []string) []*MonkeyBig {
	var monkies []*MonkeyBig
	monkey := &MonkeyBig{}
	for _, line := range lines {
		switch {
		case line == "":
			monkies = append(monkies, monkey)
			monkey = &MonkeyBig{}
		case idRe.MatchString(line):
			id, _ := strconv.Atoi(idRe.FindStringSubmatch(line)[1])
			monkey.id = id
		case itemsRe.MatchString(line):
			items := strings.Split(itemsRe.FindStringSubmatch(line)[1], ", ")
			var itemsInt []*big.Int
			for _, v := range items {
				itemInt, _ := strconv.Atoi(v)
				itemsInt = append(itemsInt, big.NewInt(int64(itemInt)))
			}
			monkey.items = itemsInt
		case opRe.MatchString(line):
			ops := opRe.FindStringSubmatch(line)
			opLeftInt, fl := strconv.Atoi(ops[1])
			if fl == nil {
				monkey.opLeft = big.NewInt(int64(opLeftInt))
			}
			monkey.op = ops[2]

			opRightInt, fr := strconv.Atoi(ops[3])
			if fr == nil {
				monkey.opRight = big.NewInt(int64(opRightInt))
			}
		case testRe.MatchString(line):
			div, _ := strconv.Atoi(testRe.FindStringSubmatch(line)[1])
			monkey.testDiv = big.NewInt(int64(div))
		case trueRe.MatchString(line):
			id, _ := strconv.Atoi(trueRe.FindStringSubmatch(line)[1])
			monkey.trueId = id
		case falseRe.MatchString(line):
			id, _ := strconv.Atoi(falseRe.FindStringSubmatch(line)[1])
			monkey.falseId = id
		}
	}
	monkies = append(monkies, monkey)
	return monkies
}

func buildSmallMonkies(lines []string) []*MonkeySmall {
	var monkies []*MonkeySmall
	monkey := &MonkeySmall{}
	for _, line := range lines {
		switch {
		case line == "":
			monkies = append(monkies, monkey)
			monkey = &MonkeySmall{}
		case idRe.MatchString(line):
			id, _ := strconv.Atoi(idRe.FindStringSubmatch(line)[1])
			monkey.id = id
		case itemsRe.MatchString(line):
			items := strings.Split(itemsRe.FindStringSubmatch(line)[1], ", ")
			var itemsInt []int
			for _, v := range items {
				itemInt, _ := strconv.Atoi(v)
				itemsInt = append(itemsInt, itemInt)
			}
			monkey.items = itemsInt
		case opRe.MatchString(line):
			ops := opRe.FindStringSubmatch(line)
			opLeftInt, fl := strconv.Atoi(ops[1])
			if fl == nil {
				monkey.opLeft = &opLeftInt
			}
			monkey.op = ops[2]

			opRightInt, fr := strconv.Atoi(ops[3])
			if fr == nil {
				monkey.opRight = &opRightInt
			}
		case testRe.MatchString(line):
			div, _ := strconv.Atoi(testRe.FindStringSubmatch(line)[1])
			monkey.testDiv = div
		case trueRe.MatchString(line):
			id, _ := strconv.Atoi(trueRe.FindStringSubmatch(line)[1])
			monkey.trueId = id
		case falseRe.MatchString(line):
			id, _ := strconv.Atoi(falseRe.FindStringSubmatch(line)[1])
			monkey.falseId = id
		}
	}
	monkies = append(monkies, monkey)
	return monkies
}

func part1(monkies []*MonkeySmall) {
	for i := 0; i < 20; i++ {
		for _, monkey := range monkies {
			monkey.inspectCt += len(monkey.items)
			for _, worry := range monkey.items {
				p1 := monkey.updateWorry(worry) / 3

				if (p1 % monkey.testDiv) != 0 {
					monkies[monkey.falseId].items = append(
						monkies[monkey.falseId].items,
						p1,
					)
				} else {
					monkies[monkey.trueId].items = append(
						monkies[monkey.trueId].items,
						p1,
					)
				}
			}
			monkey.items = []int{}
		}
	}

	slices.SortFunc(monkies, func(a, b *MonkeySmall) bool {
		return a.inspectCt > b.inspectCt
	})

	fmt.Println(monkies[0].inspectCt * monkies[1].inspectCt)
}

func part2Big(monkies []*MonkeyBig) {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		for _, monkey := range monkies {
			monkey.inspectCt += len(monkey.items)
			for _, _ = range monkey.items {
				worry := monkey.items[0]
				monkey.items = monkey.items[1:]
				// p1 := new(big.Int).Div(monkey.updateWorry(worry), big.NewInt(3))
				monkey.mutateWorry(worry)

				// fmt.Println(p1, monkey.falseId, new(big.Int).Mod(p1, monkey.testDiv).Cmp(big.NewInt(0)) == 0)
				if (new(big.Int).Mod(worry, monkey.testDiv).Cmp(big.NewInt(0))) != 0 {
					monkies[monkey.falseId].items = append(
						monkies[monkey.falseId].items,
						worry,
					)
				} else {
					monkies[monkey.trueId].items = append(
						monkies[monkey.trueId].items,
						worry,
					)
				}
			}
		}
	}

	for _, monkey := range monkies {
		fmt.Println(monkey.inspectCt)
	}
	// spew.Dump(monkies)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], []int{})
	}

	return result
}

func part2Small(monkies []*MonkeySmall) {
	supermod := 1
	for _, monkey := range monkies {
		supermod *= monkey.testDiv
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkies {
			monkey.inspectCt += len(monkey.items)
			for _, worry := range monkey.items {
				p1 := monkey.updateWorry(worry) % supermod

				if (p1 % monkey.testDiv) != 0 {
					monkies[monkey.falseId].items = append(
						monkies[monkey.falseId].items,
						p1,
					)
				} else {
					monkies[monkey.trueId].items = append(
						monkies[monkey.trueId].items,
						p1,
					)
				}
			}
			monkey.items = []int{}
		}
	}

	for _, monkey := range monkies {
		fmt.Println(monkey.id, monkey.inspectCt)
	}

	slices.SortFunc(monkies, func(a, b *MonkeySmall) bool {
		return a.inspectCt > b.inspectCt
	})

	fmt.Println(monkies[0].inspectCt * monkies[1].inspectCt)
}

func main() {
	lines := utils.FileLines(os.Args[1])

	part1(buildSmallMonkies(lines))
	part2Small(buildSmallMonkies(lines))
	// part2Big(buildBigMonkies(lines))
}
