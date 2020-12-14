package day14

import (
	"advent-of-code/utils"
	"strconv"
	"strings"
)

var mem = map[int]int{}

func setValue(bits string, value int) {
	for i, val := range bits {
		if val == 'X' {
			firstOption := bits[:i] + string('0') + bits[i+1:]
			secondOption := bits[:i] + string('1') + bits[i+1:]
			setValue(firstOption, value)
			setValue(secondOption, value)
			return
		}
	}
	result, _ := strconv.ParseInt(bits, 2, 64)
	mem[int(result)] = value
}

func getOut(key string, mask string, value int) {
	prefix := strings.Repeat("0", len(mask)-len(key))
	key = prefix + key
	var out string
	for i, val := range mask {
		if val == 'X' || val == '1' {
			out += string(val)
		} else {
			out += string(key[i])
		}
	}
	setValue(out, value)
}

func RunSecondPuzzle() (result int) {
	lines := utils.LoadFromFile("day14/data1")

	var mask string
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, " = ")[1]
		} else {
			strs := strings.Split(line, " = ")
			key, _ := strconv.Atoi(strs[0][4 : len(strs[0])-1])
			strKey := strconv.FormatInt(int64(key), 2)
			value, _ := strconv.Atoi(strs[1])
			getOut(strKey, mask, value)
		}

	}
	for _, value := range mem {
		result += value
	}
	return
}
