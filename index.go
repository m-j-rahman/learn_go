package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	find := len([]rune(toFind))

	if len(str) < find {
		return -1
	}
	for i := 0; 1 < len(str); i++ {
		if len(str[i:]) >= find {
			if s[i:i+find] == toFind {
				return i
			}
		} else {
			return -1
		}
	}
	return -1
}
