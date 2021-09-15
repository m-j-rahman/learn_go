package piscine

func TrimAtoi(s string) int {
	neg := false
	str := []rune(s)
	trim := 0

	for i := 0; i < len(str); i++ {
		if !neg && trim == 0 && str[i] == '-' {
			neg = true
		}
		if str[i] >= '0' && str[i] <= '9' {
			trim *= 10
			trim += int(str[i] - 48)
		}
	}
	if neg {
		return trim * -1
	}
	return trim
}
