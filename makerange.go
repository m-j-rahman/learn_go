package main

import "fmt"

func MakeRange(min, max int) []int {
	var size int

	if min >= max {
		return nil
	} else {
		size = max - min
	}

	answer := make([]int, size)

	for i := 0; i < size; i++ {

		answer[i] = (max - min) + 1
		for i := range answer {
			answer[i] = min + i
		}
	}
	return answer
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
