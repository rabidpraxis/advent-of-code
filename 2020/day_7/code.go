package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

var (
	containsRe = regexp.MustCompile(`(.*) (?:bags?) contain (.*).`)
	bagRe      = regexp.MustCompile(`(\d+) (.*) (?:bags?)`)
	emptyRe    = regexp.MustCompile(`(.*) (?:bag|bags) contain no other bags.`)
)

func buildInverseTree(lines []string) map[string][]string {
	inverse := make(map[string][]string)

	for _, line := range lines {
		if emptyRe.MatchString(line) {
			match := emptyRe.FindStringSubmatch(line)
			inverse[match[1]] = []string{}
		} else {
			match := containsRe.FindStringSubmatch(line)

			for _, bag := range strings.Split(match[2], ", ") {
				bagD := bagRe.FindStringSubmatch(bag)

				_, found := inverse[bagD[2]]
				if !found {
					inverse[bagD[2]] = []string{}
				}

				inverse[bagD[2]] = append(inverse[bagD[2]], match[1])
			}
		}
	}
	return inverse
}

func collectInverseBags(key string, tree map[string][]string, coll *utils.Set[string]) {
	bags, _ := tree[key]
	if len(bags) > 0 {
		for _, bag := range bags {
			coll.Add(bag)
			collectInverseBags(bag, tree, coll)
		}
	}
}

func part1(lines []string) {
	tree := buildInverseTree(lines)
	bags := utils.NewSet[string]()
	collectInverseBags("shiny gold", tree, bags)
	fmt.Println(bags.Length())
}

func calculateBagCounts(name string, tree map[string]map[string]int) int {
	bags, _ := tree[name]
	if bags == nil {
		return 1
	}

	final := 1
	for k, ct := range bags {
		final += ct * calculateBagCounts(k, tree)
	}
	return final
}

func buildCountTree(lines []string) map[string]map[string]int {
	ctTree := make(map[string]map[string]int)
	for _, line := range lines {
		if emptyRe.MatchString(line) {
			match := emptyRe.FindStringSubmatch(line)
			ctTree[match[1]] = nil
		} else {
			match := containsRe.FindStringSubmatch(line)

			bagMap, found := ctTree[match[1]]
			if !found {
				bagMap = map[string]int{}
				ctTree[match[1]] = bagMap
			}

			for _, bag := range strings.Split(match[2], ", ") {
				bagD := bagRe.FindStringSubmatch(bag)

				num, _ := strconv.Atoi(bagD[1])
				bagMap[bagD[2]] = num
			}
		}
	}
	return ctTree
}

func part2(lines []string) {
	tree := buildCountTree(lines)
	fmt.Println(calculateBagCounts("shiny gold", tree) - 1)
}

func main() {
	lines := utils.FileLines(os.Args[1])

	part1(lines)
	part2(lines)
}
