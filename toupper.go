package piscine

func ToUpper(s string) string {
	a := []rune(s)
	for i := 0; i < len(s); i++ {
		if a[i] >= 97 && a[i] <= 122 {
			a[i] = a[i] - 32
		}
	}
	return string(a)
}
