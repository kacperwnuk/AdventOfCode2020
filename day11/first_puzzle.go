package day11

import (
	"advent-of-code/utils"
)

const (
	FLOOR    = '.'
	EMPTY    = 'L'
	OCCUPIED = '#'
)

func countSeats(layout [][]rune, occupiedCounter func([][]rune, int, int) int) (totalOccupied int) {
	seatsChanged, newLayout := generateNewLayout(layout, occupiedCounter)
	for seatsChanged != 0 {
		seatsChanged, newLayout = generateNewLayout(newLayout, occupiedCounter)
	}

	for i := range newLayout {
		for j := range newLayout[i] {
			if newLayout[i][j] == OCCUPIED {
				totalOccupied++
			}
		}
	}
	return
}

func generateNewLayout(layout [][]rune, occupiedCounter func([][]rune, int, int) int) (int, [][]rune) {
	seatsChanged := 0
	newLayout := [][]rune{}

	for i := range layout {
		newLayout = append(newLayout, []rune{})
		for j := range layout[i] {
			occupiedSeats := occupiedCounter(layout, i, j)
			switch layout[i][j] {
			case EMPTY:
				if occupiedSeats == 0 {
					newLayout[i] = append(newLayout[i], OCCUPIED)
					seatsChanged++
				} else {
					newLayout[i] = append(newLayout[i], EMPTY)
				}
			case OCCUPIED:
				if occupiedSeats >= 5 {
					newLayout[i] = append(newLayout[i], EMPTY)
					seatsChanged++
				} else {
					newLayout[i] = append(newLayout[i], OCCUPIED)
				}
			case FLOOR:
				newLayout[i] = append(newLayout[i], FLOOR)
			}
		}
	}
	return seatsChanged, newLayout
}

type Pos struct {
	x int
	y int
}

var coords = []Pos{
	{1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1},
}

func countOccupied(layout [][]rune, y int, x int) (totalOccupied int) {
	for _, coord := range coords {
		xPos := x + coord.x
		yPos := y + coord.y
		if xPos < 0 || yPos < 0 || xPos >= len(layout[0]) || yPos >= len(layout) {
			continue
		}
		if layout[yPos][xPos] == OCCUPIED {
			totalOccupied++
		}
	}
	return
}

func RunFirstPuzzle() (result int) {
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

	result = countSeats(layout, countOccupied)

	return
}
