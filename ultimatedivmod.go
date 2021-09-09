package piscine

func UltimateDivMod(a *int, b *int) {
	c := 0

	c = (*a / *b)
	*b = (*a % *b)
	*a = c
}
