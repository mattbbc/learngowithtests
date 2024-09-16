package resistors

import "slices"

var COLOURS = []string{"black", "brown", "red"}

func GetValue(colour string) int {
	for idx := range COLOURS {
		if COLOURS[idx] == colour {
			return idx
		}
	}

	return -1
}

func GetValueSlices(colour string) int {
	return slices.Index(COLOURS, colour)
}
