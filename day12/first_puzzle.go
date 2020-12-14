package day12

import (
	"advent-of-code/utils"
	"math"
	"strconv"
)

//
//Action N means to move north by the given value.
//Action S means to move south by the given value.
//Action E means to move east by the given value.
//Action W means to move west by the given value.
//Action L means to turn left the given number of degrees.
//Action R means to turn right the given number of degrees.
//Action F means to move forward by the given value in the direction the ship is currently facing.

type ActionType int

const (
	N = iota
	E
	S
	W
	R
	L
	F
)

type Action struct {
	name  ActionType
	value int
}

func calculateManhattanDistance(actions []Action) int {
	direction := E
	x := 0
	y := 0

	for _, action := range actions {
		switch action.name {
		case N:
			y += action.value
		case S:
			y -= action.value
		case E:
			x += action.value
		case W:
			x -= action.value
		case L:
			direction = (direction + (4 - (action.value / 90))) % 4
		case R:
			direction = (direction + (action.value / 90)) % 4
		case F:
			switch direction {
			case N:
				y += action.value
			case S:
				y -= action.value
			case E:
				x += action.value
			case W:
				x -= action.value
			}
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day12/data1")

	actions := []Action{}
	for _, line := range lines {
		name := line[0]
		value, _ := strconv.Atoi(line[1:])
		var actionType ActionType
		switch name {
		case 'N':
			actionType = N
		case 'S':
			actionType = S
		case 'E':
			actionType = E
		case 'W':
			actionType = W
		case 'L':
			actionType = L
		case 'R':
			actionType = R
		case 'F':
			actionType = F
		}
		actions = append(actions, Action{actionType, value})
	}

	result = calculateManhattanDistance(actions)
	return
}
