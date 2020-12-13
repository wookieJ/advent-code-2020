package main

import (
	"Day-08-Handheld-Halting/common"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dataPath = "data/input"

type Operation struct {
	name  string
	value int
}

func main() {
	fmt.Println("\n--- Day 8: Handheld Halting  ---")
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
	var accumulator int
	lines := common.GetStringArrayFromStringInput(input, "\n")
	instructions := getInstructionsMap(lines)
	i := 0
	var used []int
	for true {
		if common.IntArrayContains(used, i) {
			break
		}
		if instructions[i].name == "acc" {
			accumulator += instructions[i].value
		} else if instructions[i].name == "jmp" {
			i += instructions[i].value
			continue
		}
		used = append(used, i)
		i++
	}
	return accumulator
}

func getInstructionsMap(lines []string) map[int]Operation {
	instructions := make(map[int]Operation)
	for i, line := range lines {
		l := strings.Split(line, " ")
		if strings.Contains(l[1], "+") {
			opValue, _ := strconv.Atoi(l[1][1:])
			instructions[i] = Operation{l[0], opValue}
		} else {
			opValue, _ := strconv.Atoi(l[1][1:])
			instructions[i] = Operation{l[0], (-1) * opValue}
		}
	}
	return instructions
}

func secondPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	instructions := getInstructionsMap(lines)
	result := 0
	var used []int
	for true {
		r, err := runProgramWithTerminationAtTheEnd(instructions)
		if !err {
			result = r
			break
		}
		if len(used) > 0 {
			if instructions[used[len(used)-1]].name == "jmp" {
				instructions[used[len(used)-1]] = Operation{"nop", instructions[used[len(used)-1]].value}
			} else if instructions[used[len(used)-1]].name == "nop" {
				instructions[used[len(used)-1]] = Operation{"jmp", instructions[used[len(used)-1]].value}
			}
		}
		for k, v := range instructions {
			if !common.IntArrayContains(used, k) {
				if v.name == "jmp" {
					instructions[k] = Operation{"nop", v.value}
					used = append(used, k)
					break
				} else if v.name == "nop" {
					instructions[k] = Operation{"jmp", v.value}
					used = append(used, k)
					break
				}
			}
		}
	}
	return result
}

func runProgramWithTerminationAtTheEnd(instructions map[int]Operation) (int, bool) {
	var accumulator int
	i := 0
	var used []int
	for true {
		if common.IntArrayContains(used, i) {
			return 0, true
		}
		if i == len(instructions)-1 {
			return accumulator, false
		}
		if instructions[i].name == "jmp" {
			used = append(used, i)
			i += instructions[i].value
		} else {
			if instructions[i].name == "acc" {
				accumulator += instructions[i].value
			}
			used = append(used, i)
			i++
		}
	}
	return accumulator, false
}
