package main

import (
	"Day-03-Toboggan-Trajectory/common"
	"fmt"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 3: Toboggan Trajectory  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("  Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("  Part 2 >> %d", resultPart2))
}

func firstPart(input string) int {
	return checkSlope(input, 3, 1)
}

func checkSlope(input string, dx, dy int) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	result, x, y := 0, 0, 0
	for range lines {
		if y >= len(lines) {
			break
		}
		slope := common.GetStringArrayFromStringInput(lines[y], "")
		if len(slope) <= x {
			x -= len(slope)
		}
		if slope[x] == "#" {
			result++
		}
		x += dx
		y += dy
	}
	return result
}

func secondPart(input string) int {
	return checkSlope(input, 1, 1) *
		checkSlope(input, 3, 1) *
		checkSlope(input, 5, 1) *
		checkSlope(input, 7, 1) *
		checkSlope(input, 1, 2)
}
