package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("got %q expected %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// This is a benchmark test
	// N is determined by the compiler somehow
	// Run with go test -bench=.
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
