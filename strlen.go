package piscine

func StrLen(s string) int {
	c := []rune(s)
	length := 0

	for index, letter := range c {
		letter += 0
		length = index
	}

	length += 1
	return length
}
