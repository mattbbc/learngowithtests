package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	// We pass in a pointer here as buffer.Write() is only implemented
	// on pointer receiver methods. Fprintf looks like this
	// func fmt.Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// To adhere to the io.Writer interface you need to implement .Write() i.e.
	// type Writer interface {
	//   Write(p []byte) (n int, err error)
	// }
	// Therefore &buffer implements io.Writer
	Greet(&buffer, "Potato")

	got := buffer.String()
	want := "Hello, Potato"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
