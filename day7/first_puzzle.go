package day7

import (
	"advent-of-code/utils"
	"strings"
)

type Bags []string

func (b Bags) contains(bagName string) bool {
	for _, value := range b {
		if value == bagName {
			return true
		}
	}
	return false
}

func countAllBags(bagInfo map[string]Bags) int {
	total := 0
	bagList := Bags{"shiny gold"}
	for len(bagList) > 0 {
		newBagList := Bags{}
		for key, bags := range bagInfo {
			for _, bag := range bagList {
				if bags.contains(bag) {
					newBagList = append(newBagList, key)
					delete(bagInfo, key)
					total++
					break
				}
			}
		}
		bagList = newBagList
	}
	return total
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day7/data1")

	bagInfo := make(map[string]Bags)

	for _, line := range lines {
		strs := strings.Split(line, " contain ")
		key := strs[0]
		key = strings.TrimSuffix(key, " bags")
		bagsStr := strs[1]
		bags := strings.Split(bagsStr, ", ")
		bags[len(bags)-1] = strings.TrimSuffix(bags[len(bags)-1], ".")
		for i, value := range bags {
			bags[i] = strings.TrimSuffix(value, " bag")
			bags[i] = strings.TrimSuffix(bags[i], " bags")
			bags[i] = bags[i][2:]
		}
		bagInfo[key] = bags
	}

	result = countAllBags(bagInfo)

	return
}
