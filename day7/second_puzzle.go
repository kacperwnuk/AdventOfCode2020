package day7

import (
	"advent-of-code/utils"
	"strconv"
	"strings"
)

type Bag struct {
	key   string
	value int
}

func countTotalBags(currentBag Bag, bagInfo map[string][]Bag) int {
	total := 0

	for _, bag := range bagInfo[currentBag.key] {
		if bag.value == 0 {
			break
		}
		total += bag.value + bag.value*countTotalBags(bag, bagInfo)
	}

	return total
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day7/data1")

	bagInfo := make(map[string][]Bag)

	for _, line := range lines {
		strs := strings.Split(line, " contain ")
		key := strs[0]
		key = strings.TrimSuffix(key, " bags")
		bagsStr := strs[1]
		bags := strings.Split(bagsStr, ", ")
		bags[len(bags)-1] = strings.TrimSuffix(bags[len(bags)-1], ".")

		bagList := []Bag{}
		for i, value := range bags {
			bags[i] = strings.TrimSuffix(value, " bag")
			bags[i] = strings.TrimSuffix(bags[i], " bags")
			val, _ := strconv.Atoi(string(bags[i][0]))
			bags[i] = bags[i][2:]
			bag := Bag{bags[i], val}
			bagList = append(bagList, bag)
		}
		bagInfo[key] = bagList
	}
	result = countTotalBags(Bag{"shiny gold", 1}, bagInfo)
	return
}
