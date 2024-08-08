package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("an array of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d, given %v expected %d", got, numbers, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("two slices", func(t *testing.T) {
		first := []int{1, 2}
		second := []int{0, 9}

		got := SumAll(first, second)
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	t.Run("two slices length 2", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("zero length slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSums(t, got, want)
	})

	t.Run("one length slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSums(t, got, want)
	})
}
