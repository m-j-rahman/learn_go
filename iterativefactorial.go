package main

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 || nb > 24 {
		return 0
	} else if nb >= 0 && nb <= 24 {
		for i := 1; i < nb+1; i++ {
			result = result * i
		}
	}
	return result
}

func main() {
	arg := 25
	fmt.Println(IterativeFactorial(arg))
}
