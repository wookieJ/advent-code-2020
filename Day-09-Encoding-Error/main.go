package main

import (
	"Day-09-Encoding-Error/common"
	"fmt"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 9: Encoding Error  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input, 25)
	resultPart2 := secondPart(input, 25)

	fmt.Println(fmt.Sprintf("  Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("  Part 2 >> %d", resultPart2))
}

func firstPart(input string, preambleLength int) int {
	numbers := common.GetIntArrayFromStringInput(input, "\n")
	var sumElements = numbers[:preambleLength]
	for i, number := range numbers[preambleLength:] {
		if checkIfSumOfTwoExists(sumElements[i:i+preambleLength], number) {
			sumElements = append(sumElements, number)
		} else {
			return number
		}
	}
	return 0
}

func checkIfSumOfTwoExists(array []int, sum int) bool {
	for i, v1 := range array {
		for j, v2 := range array {
			if i != j && v1+v2 == sum {
				return true
			}
		}
	}
	return false
}

func secondPart(input string, l int) int {
	findSum := firstPart(input, l)
	numbers := common.GetIntArrayFromStringInput(input, "\n")
	var summingSetMinMaxSum int
	for i := range numbers {
		var summingSet []int
		movingSum := 0
		for _, v := range numbers[i+1:] {
			movingSum += v
			if movingSum < findSum {
				summingSet = append(summingSet, v)
			} else if movingSum == findSum {
				if min, max, err := common.MinMax(summingSet); err == nil {
					if summingSetMinMaxSum < min+max {
						summingSetMinMaxSum = min + max
					}
				}
			}
		}
	}
	return summingSetMinMaxSum
}
