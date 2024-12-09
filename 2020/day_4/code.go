package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
	"golang.org/x/exp/maps"
)

var (
	passportRegex = regexp.MustCompile(`(?:^| )(\w+):([^ ]*)`)
	hgtRegex      = regexp.MustCompile(`(\d+)(cm|in)`)
	hclRegex      = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eclRegex      = regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)
	pidRegex      = regexp.MustCompile(`^\d{9}$`)

	requiredKeys = utils.NewSet[string]()
)

func parsePassports(lines []string) []map[string]string {
	currPassport := map[string]string{}
	var passports []map[string]string

	for _, line := range lines {
		if !passportRegex.MatchString(line) {
			passports = append(passports, currPassport)
			currPassport = map[string]string{}
		} else {
			matches := passportRegex.FindAllStringSubmatch(line, -1)
			// fmt.Println(strings.Split(line, " "))
			for _, match := range matches {
				currPassport[match[1]] = match[2]
			}
		}
	}

	passports = append(passports, currPassport)
	return passports
}

func hasRequiredKeys(passport map[string]string) bool {
	ss := utils.NewSet[string]()
	ss.AddMany(maps.Keys(passport))

	if ss.Subset(requiredKeys) {
		return true
	}
	return false
}

func part1(passports []map[string]string) {
	valid := 0
	for _, passport := range passports {
		if hasRequiredKeys(passport) {
			valid++
		}
	}
	fmt.Println(valid)
}

func part2(passports []map[string]string) {
	validCt := 0
	for _, passport := range passports {
		if !hasRequiredKeys(passport) {
			continue
		}

		fmt.Println(passport)

		byr, _ := passport["byr"]
		byrv, err := strconv.Atoi(byr)
		if err != nil {
			fmt.Println("byr strconv: ", byr)
			continue
		}
		if byrv < 1920 || byrv > 2002 {
			fmt.Println("byr check: ", byrv)
			continue
		}

		iyr, _ := passport["iyr"]
		iyrv, err := strconv.Atoi(iyr)
		if err != nil {
			fmt.Println("iyr strconv: ", iyr)
			continue
		}
		if iyrv < 2010 || iyrv > 2020 {
			fmt.Println("iyrv check: ", iyrv)
			continue
		}

		eyr, _ := passport["eyr"]
		eyrv, err := strconv.Atoi(eyr)
		if err != nil {
			fmt.Println("eyr strconv: ", eyr)
			continue
		}
		if eyrv < 2020 || eyrv > 2030 {
			fmt.Println("eyrv check: ", eyrv)
			continue
		}

		hgt, _ := passport["hgt"]

		if !hgtRegex.MatchString(hgt) {
			fmt.Println("hgt match: ", hgt)
			continue
		}
		hgtMatch := hgtRegex.FindStringSubmatch(hgt)
		hgtv, err := strconv.Atoi(hgtMatch[1])
		if err != nil {
			fmt.Println("hgt strconf: ", hgtMatch[1])
			continue
		}

		if hgtMatch[2] == "in" && (hgtv < 59 || hgtv > 76) {
			fmt.Println("hgt incheck: ", hgtv)
			continue
		} else if hgtMatch[2] == "cm" && (hgtv < 150 || hgtv > 193) {
			fmt.Println("hgt cmcheck: ", hgtv)
			continue
		}

		hcl, _ := passport["hcl"]
		if !hclRegex.MatchString(hcl) {
			fmt.Println("hcl match: ", hcl)
			continue
		}

		ecl, _ := passport["ecl"]
		if !eclRegex.MatchString(ecl) {
			fmt.Println("ecl match: ", ecl)
			continue
		}

		pid, _ := passport["pid"]
		if !pidRegex.MatchString(pid) {
			fmt.Println("pid match: ", pid)
			continue
		}

		validCt++
	}

	fmt.Println(validCt)
}

func main() {
	input := utils.FileLines(os.Args[1])
	requiredKeys.AddMany([]string{"hcl", "iyr", "eyr", "ecl", "pid", "byr", "hgt"})

	passports := parsePassports(input)

	part1(passports)
	part2(passports)
}
