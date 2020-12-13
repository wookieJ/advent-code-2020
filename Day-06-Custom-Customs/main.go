package main

import (
	"Day-06-Custom-Customs/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 6: Custom Customs  ---")
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
	tests := common.GetArrayOfArrays(input, "\n\n", "\n")
	result := 0
	for _, value := range tests {
		letters := common.GetAllLetters(value)
		result += len(letters)
	}
	return result
}

func secondPart(input string) int {
	tests := common.GetArrayOfArrays(input, "\n\n", "\n")
	result := 0
	for _, value := range tests {
		letters := common.GetCommonLetters(value)
		result += len(letters)
	}
	return result
}
