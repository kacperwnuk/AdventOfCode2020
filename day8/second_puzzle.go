package day8

import (
	"advent-of-code/utils"
	"fmt"
)

type Op struct {
	name   string
	symbol int
	amount int
	line   int
}

func repairLoop(operations []string, posChanged int) (int, map[int]Op, int) {
	i := 0
	visited := map[int]struct{}{}
	previousOp := map[int]Op{}
	acc := 0

	for i < len(operations) {
		opType, symbol, amount := parseOperation(operations[i])
		if _, ok := visited[i]; ok {
			fmt.Println("LOOP")
			return acc, previousOp, i
		} else {
			visited[i] = struct{}{}
		}
		oldI := i
		if i == posChanged {
			opType = "nop"
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
		previousOp[i] = Op{opType, symbol, amount, oldI}
	}
	return acc, nil, 0
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day8/data1")

	for _, line := range lines {
		_ = line
	}

	acc, previousOp, lastPos := repairLoop(lines, -1)

	posToOmit := 0
	currentOp := previousOp[lastPos]
	currentPos := currentOp.line
	for currentPos != lastPos {
		if currentOp.name == "jmp" {
			posToOmit = currentPos
			acc, previousOp, lastPos = repairLoop(lines, posToOmit)
			if previousOp == nil {
				break
			}
		}
		currentOp = previousOp[currentOp.line]
		currentPos = currentOp.line
	}

	result = acc

	return
}
