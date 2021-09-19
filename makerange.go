package piscine

func MakeRange(min, max int) []int {
	var size int

	if min >= max {
		return nil
	} else {
		size = max - min
	}

	answer := make([]int, size)

	for i := 0; i < size; i++ {

		answer[i] = (max - min) + 1
		for i := range answer {
			answer[i] = min + i
		}
	}
	return answer
}
