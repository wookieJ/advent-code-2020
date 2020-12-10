package main

import (
	"Day-10-Adapter-Array/common"
	"fmt"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 10: Adapter Array  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("  Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("  Part 2 >> %d", resultPart2))
}

func firstPart(input string) int {
	lines := common.GetIntArrayFromStringInput(input, "\n")
	var voltage, v1, v3 = 0, 0, 0
	for range lines {
		ok, v := common.HaveAnyElement(lines, []int{voltage + 1, voltage + 2, voltage + 3})
		if ok { //&& !common.IntArrayContains(used, v) {
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

func secondPart2(input string) int {
	lines := common.GetIntArrayFromStringInput(input, "\n")
	var voltage = 0
	_, max, _ := common.MinMax(lines)
	result := 1
	checkVoltage2(lines, voltage, []int{1,2,3}, &result, max)
	return result
}

func secondPart(input string) int {
	lines := common.GetIntArrayFromStringInput(input, "\n")
	var voltage = 0
	_, max, _ := common.MinMax(lines)
	result := 0
	checkVoltage(lines, voltage, &result, max)
	return result
}

func checkVoltage(lines []int, voltage int, result *int, max int) int {
	if voltage == max {
		return 1
	}
	numbers := numberOfElements(lines, []int{voltage + 1, voltage + 2, voltage + 3})
	for _, num := range numbers {
		r := checkVoltage(lines, num, result, max)
		if r > 0 {
			*result += r
		}
	}
	return 0
}

func checkVoltage2(lines []int, voltage int, v []int, result *int, max int) {
	if voltage == max {
		return
	}
	numbers := numberOfElements(lines, v)
	*result *= RevSum(len(numbers)-1) + 1
	var vv []int
	if common.IntArrayContains(v,voltage+1) {
		vv = append(vv, voltage+2,voltage+3)
	}
	if common.IntArrayContains(v,voltage+2) {
		vv = append(vv, voltage+1,voltage+3)
	}
	for _, num := range numbers {
		checkVoltage2(lines, num, []int{num + 1, num + 2, num + 3}, result, max)
	}
	return
}

func RevSum(max int) int {
	res := 0
	for i := max; i > 0; i-- {
		res += RevSum(i-1) + 1
	}
	return res
}

func numberOfElements(array1, array2 []int) []int {
	var res []int
	for _, v := range array1 {
		if common.IntArrayContains(array2, v) {
			res = append(res, v)
		}
	}
	return res
}
