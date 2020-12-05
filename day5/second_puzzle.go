package day5

import (
	"advent-of-code/utils"
	"sort"
)

func findSeat(ids []int) int {
	currentID := ids[0]
	for i := 1; i < len(ids); i++ {
		if ids[i]-currentID > 1 {
			return ids[i] - 1
		} else {
			currentID = ids[i]
		}
	}
	return 0
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day5/data1")
	var ids = []int{}
	for _, line := range lines {
		id := calculateID(line)
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	result = findSeat(ids)
	return
}
