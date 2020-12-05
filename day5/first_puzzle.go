package day5

import (
	"advent-of-code/utils"
	"fmt"
)

var rows = 128
var cols = 8

func calculateBinary(line string, width int) int {
	start := 0
	end := width
	for i, char := range line {
		if char == 'F' || line[i] == 'L' {
			end = start + (end-start)/2
		} else {
			start = start + (end-start+1)/2
		}
	}
	return start
}

func calculateID(line string) int {
	row := calculateBinary(line[0:7], rows-1)
	col := calculateBinary(line[7:10], cols-1)
	return row*8 + col
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day5/data1")
	for _, line := range lines {
		id := calculateID(line)
		fmt.Println(id)
		if result < id {
			result = id
		}
	}
	return
}
