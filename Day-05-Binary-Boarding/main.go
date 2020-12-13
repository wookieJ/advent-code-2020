package main

import (
	"Day-05-Binary-Boarding/common"
	"fmt"
	"github.com/jucardi/go-streams/streams"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 5: Binary Boarding  ---")
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

func seatToId(v interface{}) interface{} {
	seat := v.(string)
	row := binaryFinding(seat, 'F', 0, 127, 0, 7)
	column := binaryFinding(seat, 'L', 0, 7, 7, 10)
	return seatId(row, column)
}

func binaryFinding(seat string, lowPartChar byte, lowerIdx, higherIdx, lowSeatIdx, highSeatIdx int) int {
	low := lowerIdx
	high := higherIdx
	for i := lowSeatIdx; i < highSeatIdx-1; i++ {
		diff := high - low
		if seat[i] == lowPartChar {
			high -= diff / 2
			high -= 1
		} else {
			low += (diff / 2) + 1
		}
	}
	if seat[highSeatIdx-1] == lowPartChar {
		return low
	} else {
		return high
	}
}

func firstPart(input string) int {
	seats := common.GetStringArrayFromStringInput(input, "\n")
	return streams.FromArray(seats).
		Map(seatToId).
		OrderBy(common.CompareInt).
		Last().(int)
}

func secondPart(input string) int {
	seats := common.GetStringArrayFromStringInput(input, "\n")
	seatIds := streams.FromArray(seats).
		Map(seatToId).
		ToArray().([]int)
	return streams.FromArray(allSeatsIds()).
		Filter(isMissingSeat(seatIds)).
		First().(int)
}

func isMissingSeat(ids []int) func(v interface{}) bool {
	return func(v interface{}) bool {
		input := v.(int)
		if !common.IntArrayContains(ids, input) &&
			common.IntArrayContains(ids, input+1) &&
			common.IntArrayContains(ids, input-1) {
			return true
		}
		return false
	}
}

func allSeatsIds() []int {
	var ids []int
	for i := 0; i < 127; i++ {
		for j := 0; j < 8; j++ {
			ids = append(ids, seatId(i, j))
		}
	}
	return ids
}

func seatId(i int, j int) int {
	return i*8 + j
}
