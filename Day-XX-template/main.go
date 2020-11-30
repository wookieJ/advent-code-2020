package main

import (
	"_DAY_NAME_/common"
	"fmt"
)

const dataPath = "data/input"

func main() {
	fmt.Println("--- _DAY_DESC_  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("Part 2 >> %d", resultPart2))
}

func firstPart(input string) int {
	return 0
}

func secondPart(input string) int {
	return 0
}