package main

import (
	"Day-11-Seating-System/common"
	"fmt"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 11: Seating System  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("  Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("  Part 2 >> %d", resultPart2))
}

func firstPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	seatsMap := make(map[int]map[int]string, 50)
	for row, line := range lines {
		rowMap := make(map[int]string, 50)
		seats := common.GetStringArrayFromStringInput(line, "")
		for col, seat := range seats {
			rowMap[col] = seat
		}
		seatsMap[row] = rowMap
	}
	end := true
	result := 0
	for end {
		end = false
		copyMap := make(map[int]map[int]string, 50)
		for row, line := range seatsMap {
			rowMap := make(map[int]string, 50)
			for col, seat := range line {
				rowMap[col] = seat
			}
			copyMap[row] = rowMap
		}
		for row, rowSeats := range copyMap {
			for col, _ := range rowSeats {
				if shouldAliveCell(copyMap, row, col) {
					seatsMap[row][col] = "#"
					end = true
				} else if shouldKillCell(copyMap, row, col) {
					seatsMap[row][col] = "L"
					end = true
				}
			}
		}
	}
	for _, line := range seatsMap {
		for _, seat := range line {
			if seat == "#" {
				result++
			}
		}
	}
	return result
}

func shouldKillCell(seats map[int]map[int]string, row int, col int) bool {
	if seats[row][col] == "#" {
		cnt := 0
		for r := row - 1; r < row+2; r++ {
			for c := col - 1; c < col+2; c++ {
				if r != row || c != col {
					seat := seats[r][c]
					if seat == "#" {
						cnt++
					}
					if cnt == 4 {
						return true
					}
				}
			}
		}
	}
	return false
}

func shouldAliveCell(seats map[int]map[int]string, row int, col int) bool {
	s := seats[row][col]
	if s == "L" {
		cnt := 0
		for r := row - 1; r < row+2; r++ {
			for c := col - 1; c < col+2; c++ {
				if r != row || c != col {
					seat, ok := seats[r][c]
					if seat == "L" || seat == "." || !ok {
						cnt++
					}
					if cnt == 8 {
						return true
					}
				}
			}
		}
	}
	return false
}

func secondPart(input string) int {
	return 0
}
