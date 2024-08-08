package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Got %d expected %d", sum, expected)
	}
}

func ExampleAdd() {
	// This is a testable example that will appear in godoc
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
