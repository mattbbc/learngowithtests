package iteration

// Repeats a provided string a given number of times
func Repeat(letter string, number int) string {
	result := ""

	for i := 0; i < number; i++ {
		result += letter
	}

	return result
}
