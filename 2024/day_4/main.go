package main

import (
	"fmt"
	"os"

	"github.com/rabidpraxis/advent-of-code/2024/utils"
)

type pt struct {
	x, y int
}

func extractStr(lines []string, p []pt) (string, bool) {
	res := ""
	for _, v := range p {
		if v.x < 0 || v.x >= len(lines[0]) || v.y < 0 || v.y >= len(lines) {
			return "", false
		}

		res += string(lines[v.y][v.x])
	}
	return res, true
}

func part1() {
	lines := utils.FileLines(os.Args[1])
	w := len(lines[0])
	h := len(lines)

	acc := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			f, match := extractStr(lines, []pt{{j, i}, {j + 1, i}, {j + 2, i}, {j + 3, i}})
			if match && f == "XMAS" {
				acc++
			}

			b, match := extractStr(lines, []pt{{j, i}, {j - 1, i}, {j - 2, i}, {j - 3, i}})
			if match && b == "XMAS" {
				acc++
			}

			u, match := extractStr(lines, []pt{{j, i}, {j, i - 1}, {j, i - 2}, {j, i - 3}})
			if match && u == "XMAS" {
				acc++
			}

			d, match := extractStr(lines, []pt{{j, i}, {j, i + 1}, {j, i + 2}, {j, i + 3}})
			if match && d == "XMAS" {
				acc++
			}

			dbr, match := extractStr(lines, []pt{{j, i}, {j + 1, i + 1}, {j + 2, i + 2}, {j + 3, i + 3}})
			if match && dbr == "XMAS" {
				acc++
			}

			dbl, match := extractStr(lines, []pt{{j, i}, {j - 1, i - 1}, {j - 2, i - 2}, {j - 3, i - 3}})
			if match && dbl == "XMAS" {
				acc++
			}

			dur, match := extractStr(lines, []pt{{j, i}, {j + 1, i - 1}, {j + 2, i - 2}, {j + 3, i - 3}})
			if match && dur == "XMAS" {
				acc++
			}

			dul, match := extractStr(lines, []pt{{j, i}, {j - 1, i + 1}, {j - 2, i + 2}, {j - 3, i + 3}})
			if match && dul == "XMAS" {
				acc++
			}
		}
	}
	fmt.Println(acc)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func part2() {
	lines := utils.FileLines(os.Args[1])
	w := len(lines[0])
	h := len(lines)

	acc := 0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			d1, m1 := extractStr(lines, []pt{{x - 1, y - 1}, {x, y}, {x + 1, y + 1}})
			d2, m2 := extractStr(lines, []pt{{x - 1, y + 1}, {x, y}, {x + 1, y - 1}})
			if m1 && m2 && (d1 == "MAS" || Reverse(d1) == "MAS") && (d2 == "MAS" || Reverse(d2) == "MAS") {
				acc++
			}
		}
	}

	fmt.Println(acc)
}

func main() {
	part1()
	part2()
}
