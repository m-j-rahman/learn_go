package piscine

func LastRune(s string) rune {
	strArr := []rune(s)
	return rune(strArr[len(strArr)-1])
}
