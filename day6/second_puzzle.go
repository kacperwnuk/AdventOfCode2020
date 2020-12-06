package day6

import "advent-of-code/utils"

func countYesAnswer(lines []string) int {
	answered := make(map[rune]int)
	for _, line := range lines {
		for _, chr := range line {
			answered[chr] += 1
		}
	}
	total := 0
	for i := 'a'; i <= 'z'; i += 1 {
		if value, ok := answered[i]; ok {
			if value == len(lines) {
				total++
			}
		}
	}
	return total
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day6/data1")

	group := []string{}
	for _, line := range lines {
		if line == "" {
			amount := countYesAnswer(group)
			result += amount
			group = []string{}
		} else {
			group = append(group, line)
		}
	}

	return
}
