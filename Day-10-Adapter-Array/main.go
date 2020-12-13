package main

import (
	"Day-10-Adapter-Array/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 10: Adapter Array  ---")
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
	lines := common.GetIntArrayFromStringInput(input, "\n")
	var voltage, v1, v3 = 0, 0, 0
	for range lines {
		ok, v := common.HaveAnyElement(lines, []int{voltage + 1, voltage + 2, voltage + 3})
		if ok {
			var diff = v - voltage
			voltage = v
			if diff >= 0 && diff <= 3 {
				if diff == 1 {
					v1++
				} else if diff == 3 {
					v3++
				}
			}
		}
	}
	return v1 * (v3 + 1)
}

func secondPart(input string) int {
	lines := common.GetIntArrayFromStringInput(input, "\n")
	_, max, _ := common.MinMax(lines)
	lines = append(lines, 0)
	return checkVoltage(lines, 0, max, make(map[int]int))
}

func checkVoltage(values []int, voltage int, max int, cache map[int]int) int {
	if cachedValue, ok := cache[voltage]; ok {
		return cachedValue
	}
	if voltage == max {
		cache[voltage] = 1
		return 1
	}
	if !common.IntArrayContains(values, voltage) {
		cache[voltage] = 0
		return 0
	}
	result := checkVoltage(values, voltage+3, max, cache)
	result += checkVoltage(values, voltage+2, max, cache)
	result += checkVoltage(values, voltage+1, max, cache)
	cache[voltage] = result
	return result
}
