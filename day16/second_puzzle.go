package day16

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Criteria struct {
	first  map[int]bool
	second map[int]bool
}

func NewCriteria() Criteria {
	return Criteria{
		first:  map[int]bool{},
		second: map[int]bool{},
	}
}

var criteria = map[string]Criteria{}

func copyMap(oldMap map[string]Criteria) map[string]Criteria {
	newMap := map[string]Criteria{}
	for key, value := range oldMap {
		newMap[key] = value
	}
	return newMap
}

func checkCriteria(possibleCriterias map[string]Criteria, value int) map[string]Criteria {
	validCriterias := copyMap(possibleCriterias)
	for key, c := range possibleCriterias {
		if _, ok := c.first[value]; ok {
			continue
		}
		if _, ok := c.second[value]; ok {
			continue
		}
		delete(validCriterias, key)
	}
	return validCriterias
}

type Result struct {
	id            int
	possibilities int
	keys          []string
}

func findOrder(tickets [][]int) []Result {
	results := []Result{}
	for i := 0; i < len(tickets[0]); i++ {
		possibleCriterias := copyMap(criteria)
		for j := 0; j < len(tickets); j++ {
			possibleCriterias = checkCriteria(possibleCriterias, tickets[j][i])
		}
		fmt.Println(len(possibleCriterias))
		keys := []string{}
		for key := range possibleCriterias {
			keys = append(keys, key)
		}
		results = append(results, Result{id: i, possibilities: len(possibleCriterias), keys: keys})
	}
	return results
}

func isInvalidTicket(ticket []int) bool {
	for _, val := range ticket {
		if _, ok := validNumbers[val]; !ok {
			return true
		}
	}
	return false
}

func RunSecondPuzzle() (result int) {
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
			c := NewCriteria()
			for j := val; j <= val2; j++ {
				validNumbers[j] = name
				c.first[j] = true
			}
			for j := val3; j <= val4; j++ {
				validNumbers[j] = name
				c.second[j] = true
			}
			criteria[name] = c
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
		if !isInvalidTicket(values) {
			nearbyTickets = append(nearbyTickets, values)
		}
		i++
	}
	results := findOrder(nearbyTickets)
	order := generateKeyOrder(results)
	fmt.Println(order)
	return
}

func generateKeyOrder(results []Result) []string {
	keys := make([]string, len(results))
	keysUsed := map[string]struct{}{}
	sort.Slice(results, func(i, j int) bool {
		return results[i].possibilities < results[j].possibilities
	})
	for _, result := range results {
		for _, key := range result.keys {
			if _, ok := keysUsed[key]; !ok {
				keysUsed[key] = struct{}{}
				keys[result.id] = key
				fmt.Println(key, result.id)
			}
		}
	}
	return keys
}
