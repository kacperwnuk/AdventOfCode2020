package day10

import (
	"advent-of-code/utils"
	"sort"
	"strconv"
)

func findJoltage(values []int) int {
	sort.Ints(values)
	max := values[len(values)-1]
	values = append(values, max+3)
	oneDiff := 0
	threeDiff := 0

	for i, value := range values {
		if i > 0 {
			if value-values[i-1] == 1 {
				oneDiff++
			} else if value-values[i-1] == 3 {
				threeDiff++
			}
		}
	}
	return oneDiff * threeDiff
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day10/data1")

	values := []int{}
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		values = append(values, v)
	}
	values = append(values, 0)

	result = findJoltage(values)

	return
}
