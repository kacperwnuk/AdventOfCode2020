package day4

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var additionalValidation = map[string]func(string) bool{
	"byr": checkByr,
	"iyr": checkIyr,
	"eyr": checkEyr,
	"hgt": checkHgt,
	"hcl": checkHcl,
	"ecl": checkEcl,
	"pid": checkPid,
}

var totalBad = 0

func is_fully_valid(lines []string) bool {
	if is_valid(lines) {
		for _, line := range lines {
			codes := strings.Split(line, " ")
			for _, code := range codes {
				info := strings.Split(code, ":")
				id := info[0]
				value := info[1]
				if fun, ok := additionalValidation[id]; ok {
					if !fun(value) {
						fmt.Println(id, value)
						totalBad++
						return false
					}
				}
			}
		}
	} else {
		totalBad++
		return false
	}

	return true
}

func checkByr(value string) bool {
	val, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return val >= 1920 && val <= 2002
}

func checkIyr(value string) bool {
	val, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return val >= 2010 && val <= 2020
}

func checkEyr(value string) bool {
	val, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return val >= 2020 && val <= 2030
}

func checkHgt(value string) bool {

	unit := value[len(value)-2:]
	val, err := strconv.Atoi(value[:len(value)-2])
	if err != nil {
		return false
	}
	if unit == "cm" {
		return val >= 150 && val <= 193
	}
	if unit == "in" {
		return val >= 59 && val <= 76
	}
	return false
}

func checkHcl(value string) bool {
	match, _ := regexp.MatchString("^#([0-9, a-f]+)$", value)
	return match && len(value) == 7
}

func checkEcl(value string) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range colors {
		if value == color {
			return true
		}
	}
	return false
}

func checkPid(value string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", value)
	return match && len(value) == 9
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day4/data1")

	total := 0
	totalValid := 0
	passportInfo := []string{}
	for i := range lines {
		if lines[i] == "" {
			total++
			if is_fully_valid(passportInfo) {
				totalValid++
			}
			passportInfo = []string{}
		} else {
			passportInfo = append(passportInfo, lines[i])
		}
	}
	result = totalValid
	fmt.Println(totalBad, total)
	return
}
