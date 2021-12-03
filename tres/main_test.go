package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestAssignmentOne(t *testing.T) {
	diagnosticReport := "00100 11110 10110 10111 10101 01111 00111 11100 10000 11001 00010 01010"
	diagnosticReportsSlice := strings.Fields(diagnosticReport)

	t.Run("Get first bits from diagnostics", func(t *testing.T) {
		got := ConcatElementsFromSliceAtIndex(&diagnosticReportsSlice, 0)
		want := "011110011100"

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("Get gamma and epsilon from concated diagnostic bits", func(t *testing.T) {
		gotGamma, gotEpsilon := MostAndLeastOccurences("011110011100")
		wantGamma, wantEpsilon := "1", "0"
		if gotGamma != wantGamma || gotEpsilon != wantEpsilon {
			t.Errorf("Got gamma %s and epsilon %s want gamma %s and epsilon %s", gotGamma, gotEpsilon, wantGamma, wantEpsilon)
		}
	})

	t.Run("Get gamma and epsilon bit values", func(t *testing.T) {
		gotGamma, gotEpsilon := GammaAndEpsilonBitValues(diagnosticReport)
		wantGamma, wantEpsilon := "10110", "01001"

		if gotGamma != wantGamma && gotEpsilon != wantEpsilon {
			t.Errorf("Got gamma %s and epsilon %s want gamma %s and epsilon %s", gotGamma, gotEpsilon, wantGamma, wantEpsilon)
		}
	})

	t.Run("Calculate power consumption", func(t *testing.T) {
		gamma, elipson := GammaAndEpsilonBitValues(diagnosticReport)
		got := CalculatePowerConsumption(gamma, elipson)
		var want int64 = 198

		if got != want {
			t.Errorf("Got %d want %d", got, want)
		}
	})

	t.Run("Get most occurences on index", func(t *testing.T) {
		got, _ := MostAndLeastOcurrencesOnIndex(&diagnosticReportsSlice, 0)
		want := "1"

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("New slice with only value x in position n", func(t *testing.T) {
		got := FilterSliceByIndexValue(&diagnosticReportsSlice, 0, "1")
		want := []string{"11110", "10110", "10111", "10101", "11100", "10000", "11001"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}

	})

	t.Run("Calculate oxygenRating", func(t *testing.T) {
		got := CalculateOxygenRating(diagnosticReport, false)
		var want int64 = 23

		if got != want {
			t.Errorf("Got %d want %d", got, want)
		}
	})

	t.Run("Calculate C with scrubber", func(t *testing.T) {
		got := CalculateOxygenRating(diagnosticReport, true)
		var want int64 = 10

		if got != want {
			t.Errorf("Got %d want %d", got, want)
		}
	})

}
