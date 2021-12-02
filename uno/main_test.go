package main

import (
	"testing"
)

func TestSonarSweep(t *testing.T) {
	numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	t.Run("Assignment one example data", func(t *testing.T) {
		got := NumOfIncreasesInSlice(numbers)
		want := 7

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Assignment two example data", func(t *testing.T) {
		nums := SumSliceByRange(numbers, 2)
		got := NumOfIncreasesInSlice(nums)
		want := 5

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
