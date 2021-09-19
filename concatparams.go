package piscine

func ConcatParams(args []string) string {
	emptyString := ""
	count := 0

	for range args {
		count++
	}

	for i := range args {
		if i == count-1 {
			emptyString += args[i]
		} else {
			emptyString += args[i] + "\n"
		}
	}
	return emptyString
}
