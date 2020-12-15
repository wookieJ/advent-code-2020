package main

import (
	"Day-15-Rambunctious-Recitation/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 15: Rambunctious Recitation  ---")
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
	numbers := common.GetIntArrayFromStringInput(input, ",")
	return checkNthNumber(numbers, 2020)
}

func secondPart(input string) int {
	numbers := common.GetIntArrayFromStringInput(input, ",")
	return checkNthNumber(numbers, 30000000)
}

func checkNthNumber(numbers []int, n int) int {
	lastUsedMap := make(map[int][]int)
	for i, v := range numbers {
		lastUsedMap[v] = []int{i}
	}
	i := len(numbers) - 1
	for i < n-1 {
		value := numbers[i]
		indexes, containsValue := lastUsedMap[value]
		if containsValue && len(indexes) > 1 {
			id := indexes[len(indexes)-1]
			id -= indexes[len(indexes)-2]
			numbers = append(numbers, id)
			lastUsedMap[id] = append(lastUsedMap[id], i+1)
		} else {
			numbers = append(numbers, 0)
			lastUsedMap[0] = append(lastUsedMap[0], i+1)
		}
		i++
	}
	return numbers[i]
}

// todo - check algorithm (https://codegolf.stackexchange.com/questions/186654/nth-term-of-van-eck-sequence)
// todo - check https://github.com/mnml/aoc/blob/master/2020/15/2.go
