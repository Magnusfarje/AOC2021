package main

import "fmt"

func NumOfIncreasesInSlice(numbers []int) int {
	increases := 0
	for i, _ := range numbers {
		if i == 0 {
			continue
		}
		if numbers[i] > numbers[i-1] {
			increases += 1
		}
	}
	return increases
}

func SumSliceByRange(numbers []int, _range int) []int {
	sumSlice := []int{}
	for i, _ := range numbers {
		if len(numbers) <= i+_range {
			break
		}

		sum := 0
		for j := 0; j <= _range; j++ {
			sum += numbers[i+j]
		}
		sumSlice = append(sumSlice, sum)
	}
	return sumSlice
}

func main() {
	assignmentData := []int{}

	// Assignment 1
	fmt.Println(NumOfIncreasesInSlice(assignmentData))
	// Assignment 2
	fmt.Println(NumOfIncreasesInSlice(SumSliceByRange(assignmentData, 2)))
}
