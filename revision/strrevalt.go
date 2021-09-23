package main

import "fmt"

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println(Reverse("Hello World"))
}
