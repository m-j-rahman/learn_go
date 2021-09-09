package piscine

func Swap(a *int, b *int) {
	s := *a

	*a = *b
	*b = s
}
