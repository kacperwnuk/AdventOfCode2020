package day3

import (
	"advent-of-code/utils"
	"fmt"
)

var width = 31
var height = 323

var stepRight = 3

func isTree(line string, x int) bool {
	pos := x % 31
	return line[pos] == '#'
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day3/data1")
	fmt.Println(lines)
	totalTrees := 0
	currentPosition := 0
	for i := 1; i < len(lines); i++ {
		currentPosition += 3
		if isTree(lines[i], currentPosition) {
			totalTrees++
		}
	}

	return totalTrees
}
