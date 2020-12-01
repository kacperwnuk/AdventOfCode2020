package day1

import "fmt"

func RunSecondPuzzle() int {
	stringValues := loadFromFile("day1\\data1")
	intValues := convertToInt(stringValues)

	result := 0
	for i, intValue := range intValues {
		res, err := findMissingTwo(YEAR-intValue, intValues[i+1:])
		if err == nil {
			fmt.Println(intValue)
			result = res * intValue
		}
	}
	return result
}

func findMissingTwo(totalValue int, intValues []int) (int, error) {
	missingValue := make(map[int]int)
	result := 0
	for _, intValue := range intValues {
		if val, ok := missingValue[intValue]; ok {
			fmt.Println(val, " ", intValue)
			result = val * intValue
		} else {
			missingValue[totalValue-intValue] = intValue
		}
	}
	if result != 0 {
		return result, nil
	} else {
		return 0, fmt.Errorf("Cannot find 2 numbers for given criteria")
	}
}
