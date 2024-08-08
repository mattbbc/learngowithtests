package main

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers) // Must use make to create slice with preallocated length

	// Other option, create empty slice and use append()
	var sums []int

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers) // Add to slice with given length
		sums = append(sums, Sum(numbers)) // Append to slice
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
