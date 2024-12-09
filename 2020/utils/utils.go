package utils

import (
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
