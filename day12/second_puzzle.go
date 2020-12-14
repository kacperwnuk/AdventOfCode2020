package day12

import (
	"advent-of-code/utils"
	"math"
	"strconv"
)

//Action N means to move the waypoint north by the given value.
//Action S means to move the waypoint south by the given value.
//Action E means to move the waypoint east by the given value.
//Action W means to move the waypoint west by the given value.
//Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
//Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
//Action F means to move forward to the waypoint a number of times equal to the given value.

func calculateManhattanDistanceUsingWaypoint(actions []Action) int {
	x := 0
	y := 0
	wx := 10
	wy := 1

	for _, action := range actions {
		switch action.name {
		case N:
			wy += action.value
		case S:
			wy -= action.value
		case E:
			wx += action.value
		case W:
			wx -= action.value
		case L:
			turn := action.value / 90
			switch turn {
			case 3:
				oldWX := wx
				wx = wy
				wy = -oldWX
			case 2:
				wx = -wx
				wy = -wy
			case 1:
				oldWX := wx
				wx = -wy
				wy = oldWX
			}
		case R:
			turn := action.value / 90
			switch turn {
			case 1:
				oldWX := wx
				wx = wy
				wy = -oldWX
			case 2:
				wx = -wx
				wy = -wy
			case 3:
				oldWX := wx
				wx = -wy
				wy = oldWX
			}
		//	90 -> 1, -10
		//	180 -> -10, -1
		//	270 -> -1, 10
		case F:
			x += wx * action.value
			y += wy * action.value
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func RunSecondPuzzle() (result int) {
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

	result = calculateManhattanDistanceUsingWaypoint(actions)

	return
}
