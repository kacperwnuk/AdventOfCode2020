package day15

func RunFirstPuzzle() (result int) {
	values := []int{17, 1, 3, 16, 19, 0}
	//values := []int{2,1,3}
	//values := []int{0,3,6}
	history := map[int]int{}

	for j, value := range values[:len(values)-1] {
		history[value] = j + 1
	}

	i := len(values)
	for i <= 30000000 {
		result = values[i-1]
		if place, ok := history[result]; ok {
			values = append(values, i-place)
		} else {
			values = append(values, 0)
		}
		history[result] = i
		i++
	}
	return
}
