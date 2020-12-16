package day16

import (
	"advent-of-code/utils"
	"strconv"
	"strings"
)

var validNumbers = map[int]string{}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day16/data1")

	i := 0
	for lines[i] != "your ticket:" {
		if lines[i] != "" {
			values := strings.Split(lines[i], ": ")
			name := values[0]
			criterias := strings.Split(values[1], " or ")
			numbers := strings.Split(criterias[0], "-")
			numbers = append(numbers, strings.Split(criterias[1], "-")...)
			val, _ := strconv.Atoi(numbers[0])
			val2, _ := strconv.Atoi(numbers[1])
			val3, _ := strconv.Atoi(numbers[2])
			val4, _ := strconv.Atoi(numbers[3])
			for i := val; i <= val2; i++ {
				validNumbers[i] = name
			}
			for i := val3; i <= val4; i++ {
				validNumbers[i] = name
			}

		}
		i++
	}
	i++
	yourTicket := []int{}
	strValues := strings.Split(lines[i], ",")
	for _, strValue := range strValues {
		val, _ := strconv.Atoi(strValue)
		yourTicket = append(yourTicket, val)
	}
	i += 3
	nearbyTickets := [][]int{}
	for i < len(lines) {
		strValues := strings.Split(lines[i], ",")
		values := []int{}
		for _, strValue := range strValues {
			val, _ := strconv.Atoi(strValue)
			values = append(values, val)
		}
		nearbyTickets = append(nearbyTickets, values)
		i++
	}
	result = countInvalidTickets(nearbyTickets)

	return
}

func countInvalidTickets(tickets [][]int) (invalid int) {

	for _, ticket := range tickets {
		for _, val := range ticket {
			if _, ok := validNumbers[val]; !ok {
				invalid += val
			}
		}
	}
	return
}
