package day3

import (
	"advent-of-code/utils"
	"fmt"
)

func countTrees(lines []string, rightStep, downStep int) int {
	totalTrees := 0
	currentPosition := 0
	for i := downStep; i < len(lines); i += downStep {
		currentPosition += rightStep
		if isTree(lines[i], currentPosition) {
			totalTrees++
		}
	}
	return totalTrees
}

func RunSecondPuzzle() int {
	lines := utils.LoadFromFile("day3/data1")
	rightSteps := []int{1, 3, 5, 7, 1}
	downSteps := []int{1, 1, 1, 1, 2}
	result := 1
	for i := range rightSteps {
		trees := countTrees(lines, rightSteps[i], downSteps[i])
		result *= trees
		fmt.Println(trees)
	}

	return result
}
