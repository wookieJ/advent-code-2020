package main

import (
	"Day-18-Operation-Order/common"
	"fmt"
	"strings"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 18: Operation Order  ---")
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
	result := 0
	for _, line := range lines {
		result += evalMath(line, false)
	}
	return result
}

func evalMath(line string, precedence bool) int {
	if strings.Contains(line, "(") {
		bracketEquations := getBracketSubEquation(line)
		for _, bracketEquation := range bracketEquations {
			subEquationResult := evalMath(bracketEquation[1:][:len(bracketEquation)-2], precedence)
			line = strings.Replace(line, bracketEquation, fmt.Sprint(subEquationResult), 1)
		}
	}
	products := strings.Split(line, " ")
	result := toInt(products[0])
	if !precedence {
		result = noPrecedenceMath(products)
	} else {
		result = precedenceMath(products)
	}
	return result
}

func precedenceMath(products []string) int {
	for i := 1; i < len(products)-1; i += 2 {
		if products[i] == "+" {
			products[i-1] = fmt.Sprint(toInt(products[i-1]) + toInt(products[i+1]))
			products = append(products[:i], products[i+2:]...)
			i -= 2
		}
	}
	result := toInt(products[0])
	for i := 1; i < len(products)-1; i += 2 {
		if products[i] == "*" {
			result *= toInt(products[i+1])
		}
	}
	return result
}

func noPrecedenceMath(products []string) int {
	result := toInt(products[0])
	for i := 1; i < len(products)-1; i++ {
		if products[i] == "+" {
			result += toInt(products[i+1])
		} else if products[i] == "*" {
			result *= toInt(products[i+1])
		}
	}
	return result
}

func getBracketSubEquation(line string) []string {
	var result []string
	start := 0
	cnt := 0
	for i, v := range line {
		if v == '(' {
			if cnt == 0 {
				start = i
			}
			cnt++
		} else if v == ')' {
			cnt--
			if cnt == 0 {
				result = append(result, line[start:i+1])
			}
		}
	}
	return result
}

func toInt(value string) int {
	return common.StringToInt(value)
}

func secondPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	result := 0
	for _, line := range lines {
		result += evalMath(line, true)
	}
	return result
}
