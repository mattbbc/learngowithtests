package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

// Spy / Mock sleeping thing that records how many times it's called
// by appending each call type to the slice
type SpyCountdownOperations struct {
	Calls []string
}

// This method makes SpyCountdownOperations implement our Sleeper interface
// in the other file
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// This method  makes SpyCountdownOperations implement the io.Writer interface
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints the right thing", func(t *testing.T) {
		buf := &bytes.Buffer{}
		sleeper := &SpyCountdownOperations{}
		Countdown(buf, sleeper)

		got := buf.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("sleeps between each print operation", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", spySleepPrinter.Calls, want)
		}
	})
}

// This will monitor how many times the sleep fn is called
type SpyTime struct {
	durationSlept time.Duration
}

// Mock sleep function
// Adhere to Sleeper interface
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
