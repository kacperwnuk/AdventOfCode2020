package day2

func checkExistence(text string, pos1 int32, pos2 int32, letter uint8) bool {
	return (text[pos1] == letter && text[pos2] != letter) || (text[pos1] != letter && text[pos2] == letter)
}

func RunSecondPuzzle() int {
	valid := 0
	problems := loadFromFile("day2\\data1")

	for _, problem := range problems {
		exist := checkExistence(problem.text, int32(problem.minValue-1), int32(problem.maxValue-1), problem.letter)
		if exist {
			valid++
		}
	}

	return valid
}
