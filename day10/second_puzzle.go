package day10

import (
	"advent-of-code/utils"
	"sort"
	"strconv"
)

func findCombinations(values []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	results := map[int]int{}

	for i, value := range values {
		if i == 0 {
			results[value] = 1
		} else {
			sum := 0
			for j := i - 3; j < i; j++ {
				if j >= 0 {
					if values[j]-value <= 3 {
						sum += results[values[j]]
					}
				}
			}
			results[value] = sum
		}
	}
	return results[values[len(values)-1]]
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day10/data1")

	values := []int{}
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		values = append(values, v)
	}
	values = append(values, 0)
	sort.Ints(values)
	values = append(values, values[len(values)-1]+3)
	result = findCombinations(values)

	return
}
