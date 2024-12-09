package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

func part2(lines []string) {
	width := len(lines[0])
	repeat := (len(lines) * 10) / width

	var trees []string
	for i := range lines {
		trees = append(trees, strings.Repeat(lines[i], repeat))
	}

	rules := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var ruleResults []int

	for _, rule := range rules {
		right := rule[0]
		down := rule[1]

		treesCt := 0
		for i := 1; i < len(trees); i++ {
			if i*down > len(trees) {
				break
			}
			s := trees[i*down]
			if string(s[i*right]) == "#" {
				treesCt++
			}
		}

		ruleResults = append(ruleResults, treesCt)
	}

	mult := 1
	for _, v := range ruleResults {
		mult *= v
	}

	fmt.Println(mult)
}

func part1(lines []string) {
	width := len(lines[0])
	repeat := (len(lines) * 4) / width

	var trees []string
	for i := range lines {
		trees = append(trees, strings.Repeat(lines[i], repeat))
	}

	treesCt := 0
	for i := 1; i < len(trees); i++ {
		s := trees[i]
		rep := "O"
		if string(s[i*3]) == "#" {
			rep = "X"
			treesCt++
		}
		trees[i] = s[:(i*3)] + rep + s[(i*3)+1:]
	}

	// For the debugging
	for _, v := range trees {
		fmt.Println(v)
	}

	fmt.Println(treesCt)
}

func main() {
	input := utils.FileLines(os.Args[1])

	part1(input)
	part2(input)
}
