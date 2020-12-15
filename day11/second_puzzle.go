package day11

import "advent-of-code/utils"

func newOccupiedCounter(layout [][]rune, y int, x int) (totalOccupied int) {
	for _, coord := range coords {
		xPos := x
		yPos := y
	loop:
		for {
			xPos = xPos + coord.x
			yPos = yPos + coord.y
			if xPos < 0 || yPos < 0 || xPos >= len(layout[0]) || yPos >= len(layout) {
				break loop
			}
			switch layout[yPos][xPos] {
			case OCCUPIED:
				totalOccupied++
				break loop
			case EMPTY:
				break loop
			case FLOOR:
				continue
			}
		}
	}
	return
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day11/data1")

	layout := [][]rune{}
	for i, line := range lines {
		layout = append(layout, []rune{})
		for _, letter := range line {
			switch letter {
			case FLOOR:
				layout[i] = append(layout[i], FLOOR)
			case EMPTY:
				layout[i] = append(layout[i], EMPTY)
			case OCCUPIED:
				layout[i] = append(layout[i], OCCUPIED)
			}
		}
	}

	result = countSeats(layout, newOccupiedCounter)

	return
}
