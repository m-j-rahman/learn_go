package piscine

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	for i := 0; i <= 4; i++ {
		for j := 1; j < len(array); j++ {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return array[2]
}
