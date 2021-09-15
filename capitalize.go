package piscine

func Capitalize(s string) string {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	char := []rune(string(str))
	for j := 0; j <= len(s)-1; j++ {
		if j == 0 {
			if char[j] >= 97 && char[j] <= 122 {
				char[j] = char[j] - 32
			}
		} else if char[j-1] >= 97 && char[j-1] <= 122 || char[j-1] >= 65 && char[j-1] <= 90 || char[j-1] >= 48 && char[j-1] <= 57 {
		} else if char[j] >= 97 && char[j] <= 122 || char[j] >= 65 && char[j] <= 90 {
			char[j] = char[j] - 32
		}
	}
	return string(char)
}
