package piscine

func NRune(s string, n int) rune {
	strArr := []rune(s)
	if n < 0 || n > (len(strArr)) {
		return 0
	} else {
		return strArr[n-1]
	}
}
