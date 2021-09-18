package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(os.Args) - 1; i > 0; i-- {
		for _, a := range arg[i] {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
