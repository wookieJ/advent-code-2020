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

func firstPart(input string, preambuleLength int) int {
	numbers := common.GetIntArrayFromStringInput(input, "\n")
	var used []int
	var preambule = numbers[:preambuleLength]
	for i, number := range numbers[preambuleLength:] {
		if v1, v2, ok := checkIfItIsSum(preambule[i:i+preambuleLength], used, number); ok {
			preambule = append(preambule, number)
			used = append(used, v1, v2)
		} else {
			return number
		}
	}
	return 0
}

func checkIfItIsSum(preambule, used []int, number int) (int, int, bool) {
	for i, v1 := range preambule {
		for j, v2 := range preambule {
			if i != j {
				if v1+v2 == number {
					return v1, v2, true
				}
			}
		}
	}

	return 0, 0, false
}

func secondPart(input string, l int) int {
	find := firstPart(input, l)
	numbers := common.GetIntArrayFromStringInput(input, "\n")
	var maxL int
	for i, _ := range numbers {
		var result []int
		f := 0
		for _, v2 := range numbers[i+1:] {
			f += v2
			if f < find {
				result = append(result, v2)
			} else if f == find {
				min, max := MinMax(result)
				if maxL < min + max {
					maxL = min + max
				}
			}
		}
	}
	return maxL
}

func MinMax(array []int) (int, int) {
	if len(array) == 0 {
		return 0,0
	}
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
