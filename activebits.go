package piscine

func ActiveBits(n int) #int {
	count := 0
	if n < 0 {
		n = -n
	}
	re := 0
	div := 0
	DivMod(n, 2, &div, &re)
	for div > 0 {
		DivMod(n, 2, &div, &re)
		n = div
		if re == 1 {
			count++
		}
	}
	return int(count)
}
