package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GammaAndEpsilonBitValues(diagnosticReport string) (gamma string, epsilon string) {
	diagnosticReportsSlice := strings.Fields(diagnosticReport)

	for i := 0; len(diagnosticReportsSlice[0]) > i; i++ {
		_gamma, _epsilon := MostAndLeastOcurrencesOnIndex(&diagnosticReportsSlice, i)
		gamma += _gamma
		epsilon += _epsilon
	}

	return
}

func MostAndLeastOccurences(concatedDiagnostics string) (gamma string, epsilon string) {
	gammaMatch := regexp.MustCompile("1")
	matches := gammaMatch.FindAllStringIndex(concatedDiagnostics, -1)
	l := len(matches)

	if float64(l) >= (float64(len(concatedDiagnostics)) / 2) {
		return "1", "0"
	}

	return "0", "1"
}

func MostAndLeastOcurrencesOnIndex(diagnosticReportSlice *[]string, onIndex int) (most string, least string) {
	concatedDiagnostics := ConcatElementsFromSliceAtIndex(diagnosticReportSlice, onIndex)
	most, least = MostAndLeastOccurences(concatedDiagnostics)

	return
}

func CalculatePowerConsumption(gamma string, epsilon string) (powerConsumption int64) {
	gammaInt, _ := BitToInt(gamma)
	epsilonInt, _ := BitToInt(epsilon)

	powerConsumption = gammaInt * epsilonInt

	return
}

func ConcatElementsFromSliceAtIndex(diagnosticReportsSlice *[]string, index int) (concatedDiagnostics string) {
	for _, diagnostic := range *diagnosticReportsSlice {
		concatedDiagnostics += diagnostic[index : index+1]
	}

	return
}

func FilterSliceByIndexValue(diagnosticReportsSlice *[]string, filterIndex int, filter string) (filteredSlice []string) {
	for _, item := range *diagnosticReportsSlice {
		if item[filterIndex:filterIndex+1] == filter {
			filteredSlice = append(filteredSlice, item)
		}
	}

	return
}

func CalculateOxygenRating(diagnosticReport string, calculateCo2 bool) int64 {
	diagnosticReportsSlice := strings.Fields(diagnosticReport)
	newSlice := ScrubSliceData(diagnosticReportsSlice, calculateCo2)

	val, _ := BitToInt(newSlice[0])
	return val
}

func ScrubSliceData(diagnosticReportsSlice []string, calculateCo2 bool) (sliceAfterScrubbing []string) {
	sliceAfterScrubbing = diagnosticReportsSlice

	for i := 0; len(diagnosticReportsSlice[0]) > i; i++ {
		most, least := MostAndLeastOcurrencesOnIndex(&sliceAfterScrubbing, i)
		occurenceToUse := most
		if calculateCo2 {
			occurenceToUse = least
		}

		sliceAfterScrubbing = FilterSliceByIndexValue(&sliceAfterScrubbing, i, occurenceToUse)

		if len(sliceAfterScrubbing) == 1 {
			break
		}
	}

	return
}

func BitToInt(bits string) (int64, error) {
	return strconv.ParseInt(bits, 2, 64)
}

func main() {
	testInput := "inputdata"

	//Assignment 1
	gamma, elipson := GammaAndEpsilonBitValues(testInput)

	fmt.Printf("Assignment 1 answer: %d\n", CalculatePowerConsumption(gamma, elipson))
	//Assignment 2
	oxygenRating := CalculateOxygenRating(testInput, false)
	co2ScrubbingRating := CalculateOxygenRating(testInput, true)

	fmt.Printf("Assignment 2 answer: %d", oxygenRating*co2ScrubbingRating)
}
