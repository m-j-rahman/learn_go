package piscine

func IsNumeric(s string) bool {
	for _, char := range s {
		if char >= 0 && char <= 47 || char >= 58 {
			return false
		}
	}
	return true
}
