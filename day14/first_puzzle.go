package day14

import (
	"advent-of-code/utils"
	"strconv"
	"strings"
)

var memory = map[string]int{}

func calculateValue(value string, mask string) int {
	prefix := strings.Repeat("0", len(mask)-len(value))
	value = prefix + value
	var out string
	for i, val := range mask {

		if val == 'X' {
			out += string(value[i])
		} else {
			out += string(val)
		}

	}
	result, _ := strconv.ParseInt(out, 2, 64)
	return int(result)
}

func RunFirstPuzzle() (result int) {
	lines := utils.LoadFromFile("day14/data1")

	var mask string
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, " = ")[1]
		} else {
			strs := strings.Split(line, " = ")
			key := strs[0][4 : len(strs[0])-1]
			value, _ := strconv.Atoi(strs[1])
			strValue := strconv.FormatInt(int64(value), 2)
			memory[key] = calculateValue(strValue, mask)
		}

	}
	for _, value := range memory {
		result += value
	}
	return
}
