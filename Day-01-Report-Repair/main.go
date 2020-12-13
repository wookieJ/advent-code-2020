package main

import (
	"Day-01-Report-Repair/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 1: Report Repair  ---")
	input := common.GetInputFromFile(dataPath)

	start := time.Now()
	resultPart1 := firstPart(input)
	firstPartDuration := time.Since(start)
	fmt.Println(fmt.Sprintf("  Part 1 >> %d [after: %v]", resultPart1, firstPartDuration))

	start = time.Now()
	resultPart2 := secondPart(input)
	secondPartDuration := time.Since(start)
	fmt.Println(fmt.Sprintf("  Part 2 >> %d [after: %v]", resultPart2, secondPartDuration))
}

func firstPart(input string) int {
	intValues := common.GetIntArrayFromStringInput(input, "\n")
	v1, v2 := findPairs(intValues)
	return v1 * v2
}

func secondPart(input string) int {
	intValues := common.GetIntArrayFromStringInput(input, "\n")
	v1, v2, v3 := findTriples(intValues)
	return v1 * v2 * v3
}

func findPairs(values []int) (int, int) {
	for i, v1 := range values {
		for _, v2 := range values[i:] {
			if v1+v2 == 2020 {
				return v1, v2
			}
		}
	}
	return 0, 0
}

func findTriples(values []int) (int, int, int) {
	for i, v1 := range values {
		for j, v2 := range values[i:] {
			for _, v3 := range values[j:] {
				if v1+v2+v3 == 2020 {
					return v1, v2, v3
				}
			}
		}
	}
	return 0, 0, 0
}
