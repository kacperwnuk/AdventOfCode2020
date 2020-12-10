package day9

import (
	"advent-of-code/utils"
	"sort"
	"strconv"
)

var sumToFind = 85848519

func findRange(values []int) []int {
	for i, val := range values {
		sum := val
		for j := i + 1; j < len(values); j++ {
			sum += values[j]
			if sum == sumToFind {
				return values[i : j+1]
			}
		}
	}
	return nil
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day9/data1")

	values := []int{}
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		values = append(values, val)
	}

	r := findRange(values)
	sort.Ints(r)
	result = r[0] + r[len(r)-1]
	return
}
