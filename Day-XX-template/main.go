package main

import (
	"_DAY_NAME_/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- _DAY_DESC_  ---")
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
	//lines := common.GetStringArrayFromStringInput(input, "\n")
	//lines := common.GetIntArrayFromStringInput(input, "\n")
	//for _, line := range lines {
	//
	//}
	return 0
}

func secondPart(input string) int {
	return 0
}