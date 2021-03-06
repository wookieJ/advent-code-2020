package main

import (
	"Day-14-Docking-Data/common"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dataPath = "data/input"

type Memory struct {
	address int
	value   int
}

type Program struct {
	mask     string
	commands []Memory
}

func main() {
	fmt.Println("\n--- Day 14: Docking Data  ---")
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
	programs := getProgramList(input)
	memory := make(map[int]int)
	for _, program := range programs {
		for _, command := range program.commands {
			intV, _ := maskValue(command.value, program.mask, false)
			memory[command.address] = intV
		}
	}
	return sumValues(memory)
}

func getProgramList(input string) []Program {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	var programs = make([]Program, 0)
	cnt := 0
	var mask string
	programList := make([]Memory, 0)
	for _, v := range lines {
		if strings.Contains(v, "mask") {
			if mask != "" {
				programs = append(programs, Program{mask, programList})
				programList = make([]Memory, 0)
				cnt++
			}
			mask = strings.Split(v, " = ")[1]
		} else if strings.Contains(v, "mem") {
			r := strings.TrimPrefix(v, "mem[")
			value := strings.Split(r, "] = ")
			mem := Memory{common.StringToInt(value[0]), common.StringToInt(value[1])}
			programList = append(programList, mem)
		}
	}
	return append(programs, Program{mask, programList})
}

func sumValues(m map[int]int) int {
	result := 0
	for _, v := range m {
		result += v
	}
	return result
}

func maskValue(value int, mask string, floatingX bool) (int, []int) {
	binary := strconv.FormatInt(int64(value), 2)
	bin := strings.Repeat("0", 36-len(binary))
	binary = bin + binary
	var floatingIndices []int
	for i, bit := range mask {
		if bit == 'X' {
			if floatingX {
				floatingIndices = append(floatingIndices, 35-i)
			}
			continue
		} else if !floatingX || bit == '1' {
			binary = replaceAtIndex(binary, bit, i)
		}
	}
	intV64, _ := strconv.ParseInt(binary, 2, 64)
	return int(intV64), floatingIndices
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func secondPart(input string) int {
	programs := getProgramList(input)
	memory := make(map[int]int)
	for _, program := range programs {
		for _, command := range program.commands {
			intV, floatingIndices := maskValue(command.address, program.mask, true)
			perms := []int{intV}
			for _, index := range floatingIndices {
				for _, perm := range perms {
					with1 := perm | 1<<index
					with0 := with1 ^ 1<<index
					perms = append(perms, with1, with0)
				}
			}
			for _, index := range perms {
				memory[index] = command.value
			}
		}
	}
	return sumValues(memory)
}
