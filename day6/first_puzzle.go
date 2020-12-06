package day6

import "advent-of-code/utils"

func countYesAnswers(lines []string) int {
	answered := make(map[rune]bool)
	for _, line := range lines {
		for _, chr := range line {
			answered[chr] = true
		}
	}
	total := 0
	for i := 'a'; i <= 'z'; i += 1 {
		if answered[i] {
			total++
		}
	}
	return total
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day6/data1")

	group := []string{}
	for _, line := range lines {
		if line == "" {
			amount := countYesAnswers(group)
			result += amount
			group = []string{}
		} else {
			group = append(group, line)
		}
	}
	return
}
