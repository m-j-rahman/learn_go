package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if char >= 0 && char <= 47 || char >= 58 && char <= 64 || char >= 91 && char <= 96 || char >= 123 {
			return false
		}
	}
	return true
}
