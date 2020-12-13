package main

import (
	"Day-13-Shuttle-Search/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 13: Shuttle Search  ---")
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
	lines := common.GetStringArrayFromStringInput(input, "\n")
	earliest, tmp := common.StringToInt(lines[0]), lines[1]
	var ids []int
	for _, v := range common.GetStringArrayFromStringInput(tmp, ",") {
		if v != "x" {
			ids = append(ids, common.StringToInt(v))
		}
	}
	value := earliest
	for true {
		for _, bus := range ids {
			if bus != -1 && value%bus == 0 {
				return bus * (value - earliest)
			}
		}
		value++
	}
	return 0
}

func secondPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	_, tmp := common.StringToInt(lines[0]), lines[1]
	bussIds := make([]int, 0)
	indexes := make([]int, 0)
	n := 0
	t := common.GetStringArrayFromStringInput(tmp, ",")
	for _, v := range t {
		if v != "x" {
			bussIds = append(bussIds, common.StringToInt(v))
			indexes = append(indexes, n)
		}
		n++
	}
	value := 0
	diff := bussIds[0]
	for i, id := range bussIds[1:] {
		for (value+indexes[i+1])%id != 0 {
			value += diff
		}
		diff *= id
	}
	return value
}
