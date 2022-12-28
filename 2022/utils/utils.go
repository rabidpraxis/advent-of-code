package utils

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

func FileLines(fName string) []string {
	data, _ := os.ReadFile(fName)
	split := strings.Split(string(data), "\n")
	return split[:len(split)-1]
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MakeGrid(width int, height int) [][]string {
	grid := make([][]string, height)
	for i, _ := range grid {
		grid[i] = make([]string, width)
	}
	return grid
}

func PrintGrid(grid [][]string) {
	for _, v := range grid {
		fmt.Println(v)
	}
}
