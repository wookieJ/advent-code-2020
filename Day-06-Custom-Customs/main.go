package main

import (
	"Day-06-Custom-Customs/common"
	"fmt"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 6: Custom Customs  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("  Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("  Part 2 >> %d", resultPart2))
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
