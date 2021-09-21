package piscine

func FindNextPrime(nb int) int {
	if IsPrime2(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}

func IsPrime2(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
