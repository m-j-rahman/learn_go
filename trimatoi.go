package main

import "fmt"

func TrimAtoi(s string) int {
	neg := false
	str := []rune(s)
	trim := 0

	for i := 0; i < len(str); i++ {
		if !neg && trim == 0 && str[i] == '-' {
			neg = true
		}
		if str[i] >= '0' && str[i] <= '9' {
			trim *= 10
			trim += int(str[i] - 48)
		}
	}
	if neg {
		return trim * -1
	}
	return trim
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
