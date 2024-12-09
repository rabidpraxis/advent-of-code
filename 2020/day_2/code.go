package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type PasswordRule struct {
	start    int
	end      int
	letter   string
	password string
}

func part1(rules []*PasswordRule) {
	validCt := 0
	for _, rule := range rules {
		ct := 0
		for _, r := range rule.password {
			if rule.letter == string(r) {
				ct++
			}
		}

		if ct >= rule.start && ct <= rule.end {
			validCt++
		}
	}
	fmt.Println(validCt)
}

func part2(rules []*PasswordRule) {
	validCt := 0
	for _, rule := range rules {
		matchStart := rule.letter == string(rule.password[rule.start-1])
		matchEnd := rule.letter == string(rule.password[rule.end-1])

		if (matchEnd || matchStart) && !(matchEnd && matchStart) {
			validCt++
		}
	}
	fmt.Println(validCt)
}

func main() {
	input := utils.FileLines(os.Args[1])
	re := regexp.MustCompile(`(\d+)-(\d+) (\w+): (\w+)`)
	var rules []*PasswordRule
	for _, rule := range input {
		m := re.FindStringSubmatch(rule)

		min, _ := strconv.Atoi(m[1])
		max, _ := strconv.Atoi(m[2])
		letter := m[3]
		password := m[4]

		passwordRule := &PasswordRule{min, max, letter, password}
		rules = append(rules, passwordRule)
	}

	part1(rules)
	part2(rules)
}
