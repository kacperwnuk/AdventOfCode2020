package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	fmt.Println(loadFromFile("day2\\data1"))
}

type Problem struct {
	minValue int
	maxValue int
	letter   uint8
	text     string
}

func loadFromFile(fileName string) []Problem {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}

	problems := []Problem{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		valueRange := scanner.Text()
		scanner.Scan()
		l := scanner.Text()
		scanner.Scan()
		text := scanner.Text()
		minValue, _ := strconv.Atoi(strings.Split(valueRange, "-")[0])
		maxValue, _ := strconv.Atoi(strings.Split(valueRange, "-")[1])
		letter := l[0]
		problems = append(problems, Problem{minValue: minValue, maxValue: maxValue, letter: letter, text: text})

	}
	return problems
}

func countLetter(text string, letter uint8) int {
	total := 0
	for i := range text {
		if text[i] == letter {
			total += 1
		}
	}
	return total
}

func RunFirstPuzzle() int {
	valid := 0
	problems := loadFromFile("day2\\data1")

	for _, problem := range problems {
		value := countLetter(problem.text, problem.letter)
		if value >= problem.minValue && value <= problem.maxValue {
			valid++
		}
	}

	return valid
}
