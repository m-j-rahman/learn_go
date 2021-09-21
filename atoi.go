package piscine

func Atoi(str string) int {
	result := 0
	val := 1
	for i, num := range str {
		conv := int(num) - 48
		if conv <= 9 && conv >= 0 {
			result = (result * 10) + conv
		} else if conv == -3 && i == 0 {
			val = -1
		} else if conv == -5 && i == 0 {
			val = 1
		} else {
			return 0
		}
	}

	result *= val
	return result
}
