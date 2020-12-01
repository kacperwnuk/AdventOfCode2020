package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func init() {
	//fmt.Println(loadFromFile("day1\\data1"))
}

func loadFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func convertToInt(strings []string) []int {
	intValues := []int{}
	for _, stringValue := range strings {
		intValue, _ := strconv.Atoi(stringValue)
		intValues = append(intValues, intValue)
	}
	return intValues
}

var YEAR = 2020

func RunFirstPuzzle() int {
	stringValues := loadFromFile("day1\\data1")
	intValues := convertToInt(stringValues)

	missingValue := make(map[int]int)
	result := 0
	for _, intValue := range intValues {
		if val, ok := missingValue[intValue]; ok {
			fmt.Println(val, " ", intValue)
			result = val * intValue
		} else {
			missingValue[YEAR-intValue] = intValue
		}
	}
	return result
}
