package day9

import (
	"advent-of-code/utils"
	"strconv"
)

func isCorrect(preamble []int, value int) bool {
	values := map[int]struct{}{}
	for _, el := range preamble {
		if _, ok := values[el]; ok {
			return true
		}
		values[value-el] = struct{}{}
	}
	return false
}

var size = 25

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day9/data1")

	values := []int{}
	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		if i < size {
			values = append(values, value)
			continue
		}
		if !isCorrect(values, value) {
			result = value
			break
		}
		values = append(values, value)
		values = values[1:]
	}

	return
}
