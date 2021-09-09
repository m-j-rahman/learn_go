package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	r := []rune(s)

	for index, letter := range r {
		index += 0
		z01.PrintRune(letter)
	}
}
