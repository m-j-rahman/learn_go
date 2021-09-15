package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var a []rune
	if n == 0 {
		a = append(a, 48)
	}

	for i := 0; n > 0; i++ {
		a = append(a, rune(n%10)+48)
		n = n / 10
	}

	l := len(a)
	for j := range a {
		if j == l-1 {
		} else {
			for i := 0; i < l-1; i++ {
				if a[i] > a[i+1] {
					a[i], a[i+1] = a[i+1], a[i]
				}
			}
		}
	}
	for _, r := range a {
		z01.PrintRune(r)
	}
}
