package day13

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type Bus struct {
	id     int
	offset int
}

func calculateTimestamp(buses []Bus) int {
	value := 100000000000430
	multiplier := 631
	//buses = buses[1:]
	notFound := true
	for notFound {
		value += multiplier
		for _, bus := range buses {
			if (value+bus.offset)%bus.id == 0 {
				notFound = false
			} else {
				notFound = true
				break
			}
		}
		fmt.Println(value)
	}
	return value
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day13/data1")
	buses := []Bus{}
	for i, value := range strings.Split(lines[0], ",") {
		if value != "x" {
			busNumber, _ := strconv.Atoi(value)
			bus := Bus{busNumber, i}
			buses = append(buses, bus)
		}
	}
	result = calculateTimestamp(buses)
	return
}
