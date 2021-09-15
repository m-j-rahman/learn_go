package piscine

func ToLower(s string) string {
	a := []rune(s)
	for i := 0; i < len(s); i++ {
		if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] = a[i] + 32
		}
	}
	return string(a)
}
