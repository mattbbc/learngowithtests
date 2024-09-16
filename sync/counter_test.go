package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := &Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := &Counter{}

		// Simulate multiple goroutines trying to update the same Counter instance
		// (race condition)
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

func assertCount(t testing.TB, c *Counter, want int) {
	t.Helper()
	if c.Value() != want {
		// We could say c.value instead of the function here
		// as we are in the same package as the counter
		// but we can't do that from another package.
		t.Errorf("got %d, wanted %d", c.Value(), want)
	}
}
