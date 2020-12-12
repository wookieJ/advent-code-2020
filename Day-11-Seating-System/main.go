package main

import (
	"Day-11-Seating-System/common"
	"Day-11-Seating-System/point"
	"fmt"
)

const dataPath = "data/input"

type SeatsMap map[int]map[int]string

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
	seatsMap := getSeatsMapFromInput(lines)
	return cellAutomataRun(seatsMap, areAllNeighboursFree, areFourNeighboursTaken)
}

func cellAutomataRun(seatsMap SeatsMap, aliveCellAlgorithm func(seats SeatsMap, row int, col int) bool,
	killCellAlgorithm func(seats SeatsMap, row int, col int) bool) int {
	end := true
	for end {
		end = false
		copyMap := common.CopyMapOfIntMap(seatsMap)
		for row, rowSeats := range copyMap {
			for col, _ := range rowSeats {
				if aliveCellAlgorithm(copyMap, row, col) {
					seatsMap[row][col] = "#"
					end = true
				} else if killCellAlgorithm(copyMap, row, col) {
					seatsMap[row][col] = "L"
					end = true
				}
			}
		}
	}
	return countTakenSeats(seatsMap)
}

func countTakenSeats(seatsMap SeatsMap) int {
	result := 0
	for _, line := range seatsMap {
		for _, seat := range line {
			if seat == "#" {
				result++
			}
		}
	}
	return result
}

func getSeatsMapFromInput(lines []string) SeatsMap {
	seatsMap := make(SeatsMap, 50)
	for row, line := range lines {
		rowMap := make(map[int]string, 50)
		seats := common.GetStringArrayFromStringInput(line, "")
		for col, seat := range seats {
			rowMap[col] = seat
		}
		seatsMap[row] = rowMap
	}
	return seatsMap
}

func areFourNeighboursTaken(seats SeatsMap, row int, col int) bool {
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

func areAllNeighboursFree(seats SeatsMap, row int, col int) bool {
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

func areAllEightDirectionVisibleSeatsFree(seats SeatsMap, row int, col int) bool {
	s := seats[row][col]
	if s == "L" {
		foundedInDirection := make(map[point.Point]bool)
		for i := 0; i < len(seats) || i < len(seats[0]); i++ {
			for r := row - i - 1; r < row+i+2; r += i + 1 {
				for c := col - i - 1; c < col+i+2; c += i + 1 {
					if (r != row || c != col) && r >= 0 && c >= 0 {
						seat := seats[r][c]
						if seat == "#" && foundedInDirection[computeDirection(row-r, col-c)] == false {
							return false
						} else if seat == "L" {
							direction := computeDirection(row-r, col-c)
							foundedInDirection[direction] = true
						}
					}
				}
			}
		}
		return true
	}
	return false
}

func fiveOccupiedSeatsVisible(seats SeatsMap, row int, col int) bool {
	s := seats[row][col]
	if s == "#" {
		foundedInDirection := make(map[point.Point]bool)
		cnt := 0
		for i := 0; i < len(seats) || i < len(seats[0]); i++ {
			for r := row - i - 1; r < row+i+2; r += i + 1 {
				for c := col - i - 1; c < col+i+2; c += i + 1 {
					if r != row || c != col && r >= 0 && c >= 0 {
						seat := seats[r][c]
						if seat == "#" && foundedInDirection[computeDirection(row-r, col-c)] == false {
							cnt++
							foundedInDirection[computeDirection(row-r, col-c)] = true
						} else if seat == "L" {
							direction := computeDirection(row-r, col-c)
							foundedInDirection[direction] = true
						}
						if cnt == 5 {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func computeDirection(x, y int) point.Point {
	p := point.Point{X: x, Y: y}
	return p.DirectionVector(0.001)
}

func secondPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	seatsMap := getSeatsMapFromInput(lines)
	return cellAutomataRun(seatsMap, areAllEightDirectionVisibleSeatsFree, fiveOccupiedSeatsVisible)
}
