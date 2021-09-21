package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, num := range s {
		if num < '0' || num > '9' {
			return 0
		} else {
			conv := int(num) - 48
			result = (result * 10) + conv
		}
	}
	return result
}
