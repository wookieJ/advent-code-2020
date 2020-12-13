package main

import (
	"Day-12-Rain-Risk/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 12: Rain Risk  ---")
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
	x, y := 0, 0
	dx, dy := 1, 0
	for _, command := range lines {
		direction, value := command[0], common.StringToInt(command[1:])
		switch direction {
		case 'N':
			y += value
		case 'S':
			y -= value
		case 'W':
			x -= value
		case 'E':
			x += value
		case 'L':
			for value > 0 {
				value -= 90
				dx, dy = dy*(-1), dx
			}
		case 'R':
			for value > 0 {
				value -= 90
				dx, dy = dy, dx*(-1)
			}
		case 'F':
			x += value * dx
			y += value * dy
		}
	}
	return common.Abs(x) + common.Abs(y)
}

func secondPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	x, y := 0, 0
	wx, wy := 10, 1
	for _, command := range lines {
		direction, value := command[0], common.StringToInt(command[1:])
		switch direction {
		case 'N':
			wy += value
		case 'S':
			wy -= value
		case 'W':
			wx -= value
		case 'E':
			wx += value
		case 'L':
			for value > 0 {
				value -= 90
				wx, wy = wy*(-1), wx
			}
		case 'R':
			for value > 0 {
				value -= 90
				wx, wy = wy, wx*(-1)
			}
		case 'F':
			x += value * wx
			y += value * wy
		}
	}
	return common.Abs(x) + common.Abs(y)
}
