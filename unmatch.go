package piscine

func Unmatch(arr []int) int {
	var count int

	for _, a := range arr {
		count = 0
		for _, v := range arr {
			if v == a {
				count++
			}
		}
		if count%2 != 0 {
			return a
		}
	}
	return -1
}
