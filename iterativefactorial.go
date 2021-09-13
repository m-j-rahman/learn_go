package main

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 {
		return 0
	}

	for i := 1; i < nb+1; i++ {
		result = result * i
	}

	return result
}

func main() {
	arg := 0
	fmt.Println(IterativeFactorial(arg))
}
