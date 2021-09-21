package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, s := range arr {
		if f(s) {
			return true
		}
	}
	return false
}
