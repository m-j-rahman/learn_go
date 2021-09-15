package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char >= 0 && char <= 31 || char >= 127 {
			return false
		}
	}
	return true
}
