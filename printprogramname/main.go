package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := []rune(os.Args[0])
	var index int
	for i := len(argument) - 1; i > 0; i-- {
		if argument[i] == '/' {
			index = 2
			break
		}
	}
	for i := index - 1; i < len(argument); i++ {
		z01.PrintRune(argument[i])
	}
	z01.PrintRune('\n')
}
