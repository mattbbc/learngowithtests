package resistors

import "testing"

func TestGetValue(t *testing.T) {
	cases := []struct {
		Colour        string
		ExpectedValue int
	}{
		{"black", 0},
		{"brown", 1},
		{"red", 2},
	}

	for _, tc := range cases {
		got := GetValue(tc.Colour)

		if got != tc.ExpectedValue {
			t.Errorf("got %d, wanted %d", got, tc.ExpectedValue)
		}
	}
}

func BenchmarkGetValue(b *testing.B) {
	// This is a benchmark test
	// N is determined by the compiler somehow
	// Run with go test -bench=.
	for i := 0; i < b.N; i++ {
		GetValue("black")
	}
}

func TestGetValueSlices(t *testing.T) {
	cases := []struct {
		Colour        string
		ExpectedValue int
	}{
		{"black", 0},
		{"brown", 1},
		{"red", 2},
	}

	for _, tc := range cases {
		got := GetValueSlices(tc.Colour)

		if got != tc.ExpectedValue {
			t.Errorf("got %d, wanted %d", got, tc.ExpectedValue)
		}
	}
}

func BenchmarkGetValueSlices(b *testing.B) {
	// This is a benchmark test
	// N is determined by the compiler somehow
	// Run with go test -bench=.
	for i := 0; i < b.N; i++ {
		GetValueSlices("black")
	}
}
