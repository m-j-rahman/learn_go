package main

import "fmt"

func AlphaCount(s string) int {
	result := []rune(s)
	count := 0

	for i := 0; i < len(result); i++ {
		if (result[i] >= 'A' && result[i] <= 'Z') || (result[i] >= 'a' && result[i] <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
