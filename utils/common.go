package utils

import (
	"bufio"
	"os"
)

func LoadFromFile(fileName string) []string {
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
