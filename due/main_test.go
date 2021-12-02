package main

import "testing"

func TestDirections(t *testing.T) {
	directions := "forward 5 down 5 forward 8 up 3 down 8 forward 2"

	t.Run("Assignment one example data", func(t *testing.T) {
		got := SumDirections(DirectionsStringToSlice(directions))
		want := 150

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, directions)
		}
	})

	t.Run("Assignment two example data", func(t *testing.T) {
		got := SumDirectionsWithAim(DirectionsStringToSlice(directions))
		want := 900
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, directions)
		}
	})
}
