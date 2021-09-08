package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	aRune := "0123456789"

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if i < j && j < k {

					z01.PrintRune(rune(aRune[i]))
					z01.PrintRune(rune(aRune[j]))
					z01.PrintRune(rune(aRune[k]))

					if i == 7 && k == 9 {
						z01.PrintRune(rune(10))
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
}
