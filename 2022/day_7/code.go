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
	reCd     = regexp.MustCompile(`\$ cd (.*)`)
	fileSize = regexp.MustCompile(`(\d+) \w+`)
)

func processLines(line string, lines []string, paths []string, sizes map[string]int) {
	switch {
	case reCd.MatchString(line):
		nextPath := reCd.FindStringSubmatch(line)[1]
		if nextPath == ".." {
			paths = paths[:len(paths)-1]
		} else if nextPath == "/" {
			paths = append(paths, "root")
		} else {
			paths = append(paths, nextPath)
		}
	case fileSize.MatchString(line):
		sizeString := fileSize.FindStringSubmatch(line)[1]
		size, _ := strconv.Atoi(sizeString)

		for i := range paths {
			path := strings.Join(paths[:i+1], ":")

			currSize, found := sizes[path]
			if !found {
				sizes[path] = size
			} else {
				sizes[path] = currSize + size
			}
		}
	}

	if len(lines) > 0 {
		processLines(lines[0], lines[1:], paths, sizes)
	}
}

func part1(sizes map[string]int) {
	sum := 0
	for _, v := range sizes {
		if v <= 100000 {
			sum += v
		}
	}
	fmt.Println(sum)
}

func part2(sizes map[string]int) {
	currentSpace, _ := sizes["root"]
	toRemove := 30000000 - (70000000 - currentSpace)

	minRemoved := currentSpace
	for _, v := range sizes {
		if v >= toRemove && v <= minRemoved {
			minRemoved = v
		}
	}

	fmt.Println(minRemoved)
}

func main() {
	lines := utils.FileLines(os.Args[1])

	sizes := map[string]int{}
	processLines(lines[0], lines[1:], []string{}, sizes)

	part1(sizes)
	part2(sizes)
}
