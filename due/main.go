package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DirectionsStringToSlice(directions string) []string {
	return strings.Fields(directions)
}

func SumDirections(directions []string) int {
	forward := 0
	depth := 0

	for index, direction := range directions {
		if len(directions) <= index+1 {
			break
		}
		value, _ := strconv.Atoi(directions[index+1])

		switch direction {
		case "forward":
			forward += value
		case "up":
			depth -= value
		case "down":
			depth += value
		default:
			continue
		}
	}
	return forward * depth
}

func SumDirectionsWithAim(directions []string) int {
	forward := 0
	depth := 0
	aim := 0

	for index, direction := range directions {
		if len(directions) <= index+1 {
			break
		}
		value, _ := strconv.Atoi(directions[index+1])

		switch direction {
		case "forward":
			forward += value
			depth += (aim * value)
		case "up":
			aim -= value
		case "down":
			aim += value
		default:
			continue
		}
	}
	return forward * depth
}

func main() {
	directionsSlice := DirectionsStringToSlice(testString())
	fmt.Println(SumDirectionsWithAim(directionsSlice))
}

func testString() string {
	return "assignment data"
}
