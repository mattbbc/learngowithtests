package racing

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(first, second string, timeout time.Duration) (winner string, err error) {
	// select allows you to wait on multiple channels
	// Both ping functions return a channel
	// Whichever one writes to its channel first will get its code executed in its case
	select {
	case <-ping(first):
		return first, nil
	case <-ping(second):
		return second, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", first, second)
	}
}

func ping(url string) chan struct{} {
	// Always make() channels
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
