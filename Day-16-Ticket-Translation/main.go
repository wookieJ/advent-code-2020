package main

import (
	"Day-16-Ticket-Translation/common"
	"Day-16-Ticket-Translation/point"
	"Day-16-Ticket-Translation/regex"
	"fmt"
	"strings"
	"time"
)

const dataPath = "data/input"

func main() {
	fmt.Println("\n--- Day 16: Ticket Translation  ---")
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
	nearbyTickets, rules, _, _ := parseInput(input)
	invalids := getInvalidFields(nearbyTickets, rules)
	return sumArray(invalids)
}

func sumArray(invalids map[int]int) int {
	res := 0
	for _, v := range invalids {
		res += v
	}
	return res
}

func getInvalidFields(tickets [][]int, rules map[string]point.PointP) map[int]int {
	var invalids = make(map[int]int)
	for i, ticket := range tickets {
		for _, field := range ticket {
			valid := false
			for _, rule := range rules {
				if isValidField(field, rule) {
					valid = true
					break
				}
			}
			if !valid {
				invalids[i] = field
			}
		}
	}
	return invalids
}

func isValidField(field int, rule point.PointP) bool {
	return field >= rule.Lower.X && field <= rule.Lower.Y || field >= rule.Upper.X && field <= rule.Upper.Y
}

func secondPart(input string) int {
	nearbyTickets, rules, yourTicket, rulesNames := parseInput(input)
	invalids := getInvalidFields(nearbyTickets, rules)
	var validTickets = make([][]int, 0)
	for i, ticket := range nearbyTickets {
		if _, ok := invalids[i]; !ok {
			validTickets = append(validTickets, ticket)
		}
	}
	var rulesOrder = make(map[int][]int)
	for ruleIndex, name := range rulesNames {
		rule := rules[name]
		for fieldIndex := range validTickets[0] {
			if ruleValid(rule, validTickets, fieldIndex) {
				if value, ok := rulesOrder[ruleIndex]; ok {
					rulesOrder[ruleIndex] = append(value, fieldIndex)
				} else {
					rulesOrder[ruleIndex] = []int{fieldIndex}
				}
			}
		}
	}
	cnt := 1
	var finalFields = make(map[int]int)
	var used []int
	for range rulesOrder {
		for ruleIndex, fieldIndexes := range rulesOrder {
			if len(fieldIndexes) == cnt {
				for _, fieldIndex := range fieldIndexes {
					if !common.IntArrayContains(used, fieldIndex) {
						finalFields[ruleIndex] = fieldIndex
						used = append(used, fieldIndex)
						cnt++
					}
				}
			}
		}
	}
	result := 1
	for ruleIndex, fieldIndex := range finalFields {
		if strings.Contains(rulesNames[ruleIndex], "departure") {
			result *= yourTicket[fieldIndex]
		}
	}
	return result
}

func parseInput(input string) ([][]int, map[string]point.PointP, []int, []string) {
	var nearbyTickets [][]int
	var rules = make(map[string]point.PointP)
	var yourTicket = make([]int, 0)
	var rulesNames = make([]string, 0)
	inputSplit := common.GetStringArrayFromStringInput(input, "\n\n")
	rulesLines := common.GetStringArrayFromStringInput(inputSplit[0], "\n")
	yourTicketLines := common.GetStringArrayFromStringInput(strings.TrimLeft(inputSplit[1], "your ticket:\n"), "\n")
	odderTicketsLines := common.GetStringArrayFromStringInput(strings.TrimLeft(inputSplit[2], "nearby tickets:\n"), "\n")
	for _, line := range rulesLines {
		r := regex.NewRegex("([a-zA-Z0-9 ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
		groups := r.GetGroups(line)
		if len(groups) > 0 {
			rulesNames = append(rulesNames, groups[0][1])
			rules[groups[0][1]] = point.PointP{Lower: point.Point{X: common.StringToInt(groups[0][2]), Y: common.StringToInt(groups[0][3])},
				Upper: point.Point{X: common.StringToInt(groups[0][4]), Y: common.StringToInt(groups[0][5])}}
		}
	}
	for _, line := range yourTicketLines {
		yourTicket = common.GetIntArrayFromStringInput(line, ",")
	}
	for _, line := range odderTicketsLines {
		ticket := common.GetIntArrayFromStringInput(line, ",")
		nearbyTickets = append(nearbyTickets, ticket)
	}
	return nearbyTickets, rules, yourTicket, rulesNames
}

func ruleValid(rule point.PointP, tickets [][]int, fieldIdx int) bool {
	for _, ticket := range tickets {
		if !isValidField(ticket[fieldIdx], rule) {
			return false
		}
	}
	return true
}
