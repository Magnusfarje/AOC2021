package main

import (
	"testing"
)

func TestSonarSweep(t *testing.T) {
	numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	t.Run("Count increases between adjecant values in array", func(t *testing.T) {
		got := CountIncreasesBetweenIndexes(numbers, 1)
		want := 7

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Count increases between adjecant+2 values in array", func(t *testing.T) {
		got := CountIncreasesBetweenIndexes(numbers, 3)
		want := 5

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
