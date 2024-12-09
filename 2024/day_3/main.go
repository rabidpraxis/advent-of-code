package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/2024/utils"
)

func nashyBashy(mul string) int64 {
	nums := strings.Split(mul[4:len(mul)-1], ",")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	return int64(num1 * num2)
}

func beckyDoubleYou(line string) int64 {
	matcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := matcher.FindAll([]byte(line), -1)
	acc := 0
	for _, match := range matches {
		strMatch := string(match)
		s := strMatch[4:]
		// Remove the last character
		s = s[:len(s)-1]
		nums := strings.Split(s, ",")
		// Parse both numbers
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		// Multiply the numbers
		acc += num1 * num2
	}
	return int64(acc)
}

func part1() {
	lines := utils.FileLines(os.Args[1])
	acc := int64(0)
	for _, line := range lines {
		acc += beckyDoubleYou(line)
	}
	fmt.Println(acc)
}

func part2() {
	data, _ := os.ReadFile(os.Args[1])
	line := string(data)
	matcher := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	matches := matcher.FindAll([]byte(line), -1)
	actual := []string{}
	inDont := false
	for _, match := range matches {
		s := string(match)

		if s == "do()" {
			inDont = false
		} else if s == "don't()" {
			inDont = true
		} else if !inDont {
			actual = append(actual, s)
		}
	}

	acc := int64(0)
	for _, a := range actual {
		acc += nashyBashy(a)
	}

	fmt.Println(acc)
}

func main() {
	part2()
}
