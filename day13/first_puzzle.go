package day13

import (
	"advent-of-code/utils"
	"sort"
	"strconv"
	"strings"
)

func findNearestBus(timestamp int, busNumbers []int) int {
	sort.Slice(busNumbers, func(i, j int) bool {
		iTimeToWait := busNumbers[i] - (timestamp % busNumbers[i])
		jTimeToWait := busNumbers[j] - (timestamp % busNumbers[j])
		return iTimeToWait < jTimeToWait
	})
	winningNumber := busNumbers[0]
	timeToWait := winningNumber - (timestamp % winningNumber)
	return timeToWait * winningNumber
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day13/data1")
	timestamp, _ := strconv.Atoi(lines[0])
	busNumbers := []int{}
	for _, value := range strings.Split(lines[1], ",") {
		if value != "x" {
			busNumber, _ := strconv.Atoi(value)
			busNumbers = append(busNumbers, busNumber)
		}
	}

	result = findNearestBus(timestamp, busNumbers)
	return
}
