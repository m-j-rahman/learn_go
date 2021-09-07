package main

import (
	"github.com/01-edu/z01"
)

func main() {
	asciiStart := 48
	asciiEnd := 57

	for asciiStart <= asciiEnd {
		z01.PrintRune(rune(asciiStart))

		asciiStart++

	}

	z01.PrintRune('\n')
}
