package main

import (
	"Day-17-Conway-Cubes/common"
	"fmt"
	"time"
)

const dataPath = "data/input"

type Grid map[int]map[int]map[int]bool

func main() {
	fmt.Println("\n--- Day 17: Conway Cubes  ---")
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
	grid := getGridMapFromInput(lines)
	return cellAutomataRun(grid, shouldActivate, 6)
}

func cellAutomataRun(grid Grid, activateFunction func(seats Grid, z, y, x int) bool, nCycles int) int {
	size := len(grid[0])
	for nCycles > 0 {
		copyMap := copyMap(grid)
		for z, slice := range copyMap {
			for y, row := range slice {
				for x := range row {
					grid[z][y][x] = activateFunction(copyMap, z, y, x)
				}
			}
		}
		cnt := 6 - nCycles
		grid[cnt-1] = newEmptyLayer(size+1)
		grid[cnt+1] = newEmptyLayer(size+1)
		nCycles -= 1
	}
	return countActiveStates(grid)
}

func newEmptyLayer(size int) map[int]map[int]bool {
	layer := make(map[int]map[int]bool, size)
	for i := 0; i < size; i++ {
		layer[i] = make(map[int]bool, size)
		for j := 0; j < size; j++ {
			layer[i][j] = false
		}
	}
	return layer
}

func copyMap(originalMap Grid) Grid {
	copyMap := make(Grid, len(originalMap))
	for k1, v1 := range originalMap {
		subMap1 := make(map[int]map[int]bool, len(originalMap))
		for k2, v2 := range v1 {
			subMap2 := make(map[int]bool, len(v2))
			for k2, v2 := range v2 {
				subMap2[k2] = v2
			}
			subMap1[k2] = subMap2
		}
		copyMap[k1] = subMap1
	}
	return copyMap
}

func countActiveStates(grid Grid) int {
	result := 0
	for _, z := range grid {
		for _, y := range z {
			for _, x := range y {
				if x {
					result++
				}
			}
		}
	}
	return result
}

func getGridMapFromInput(lines []string) Grid {
	size := len(lines)
	grid := make(Grid, size)
	zMap := make(map[int]map[int]bool)
	for y, line := range lines {
		yMap := make(map[int]bool, size)
		rowArray := common.GetStringArrayFromStringInput(line, "")
		for x, cubeState := range rowArray {
			if cubeState == "#" {
				yMap[x] = true
			} else {
				yMap[x] = false
			}
		}
		zMap[y] = yMap
	}
	grid[0] = zMap
	return grid
}

func shouldActivate(grid Grid, z, y, x int) bool {
	active := grid[z][y][x]
	activeNeighboursNumber := 0
	for zn := z - 1; zn < z+2; zn++ {
		for yn := y - 1; yn < y+2; yn++ {
			for xn := x - 1; xn < x+2; xn++ {
				if zn != z || yn != y || xn != x {
					state := grid[zn][yn][xn]
					if state {
						activeNeighboursNumber++
					}
					if activeNeighboursNumber > 3 {
						return false
					}
				}
			}
		}
	}
	if (active && (activeNeighboursNumber == 2 || activeNeighboursNumber == 3)) || !active && activeNeighboursNumber == 3 {
		return true
	}
	return false
}

func secondPart(input string) int {
	return 0
}
