package main

import "fmt"

func CountIncreasesBetweenIndexes(numbers []int, indexStep int) int {
	increases := 0
	for i, _ := range numbers {
		if len(numbers) <= i+indexStep {
			break
		}
		if numbers[i] < numbers[i+indexStep] {
			increases += 1
		}
	}
	return increases
}

func main() {
	assignmentData := []int{}

	// Assignment 1
	fmt.Println(CountIncreasesBetweenIndexes(assignmentData, 1))
	// Assignment 2
	fmt.Println(CountIncreasesBetweenIndexes(assignmentData, 3))
}
