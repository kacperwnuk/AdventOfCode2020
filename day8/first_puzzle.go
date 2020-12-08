package day8

import (
	"advent-of-code/utils"
	"strconv"
	"strings"
)

func parseOperation(operationStr string) (opType string, symbol int, amount int) {
	strs := strings.Split(operationStr, " ")
	opType = strs[0]
	s := strs[1][0]
	if s == '-' {
		symbol = -1
	} else {
		symbol = 1
	}
	amount, _ = strconv.Atoi(strs[1][1:len(strs[1])])
	return
}

func findLoop(operations []string) int {
	i := 0
	visited := map[int]struct{}{}
	acc := 0

	for i <= len(operations) {
		opType, symbol, amount := parseOperation(operations[i])
		if _, ok := visited[i]; ok {
			return acc
		} else {
			visited[i] = struct{}{}
		}

		switch opType {
		case "nop":
			i += 1
		case "acc":
			i += 1
			acc += symbol * amount
		case "jmp":
			i += symbol * amount
		}
	}
	return acc
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day8/data1")
	result = findLoop(lines)
	return
}
