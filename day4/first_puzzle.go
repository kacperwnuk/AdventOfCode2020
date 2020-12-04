package day4

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func is_valid(lines []string) bool {
	idsFound := make(map[string]bool)
	for _, line := range lines {
		codes := strings.Split(line, " ")
		for _, code := range codes {
			info := strings.Split(code, ":")
			id := info[0]
			idsFound[id] = true
		}
	}
	for _, requiredField := range requiredFields {
		if !idsFound[requiredField] {
			return false
		}
	}
	return true
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day4/data1")

	total := 0
	totalValid := 0
	passportInfo := []string{}
	for i := range lines {
		if lines[i] == "" {
			if is_valid(passportInfo) {
				totalValid++
			} else {
				fmt.Println(total)
			}
			total++
			passportInfo = []string{}
		} else {
			passportInfo = append(passportInfo, lines[i])
		}
	}
	result = totalValid
	return
}
