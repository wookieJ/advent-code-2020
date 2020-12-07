package main

import (
	"Day-07-Handy-Haversacks/common"
	"Day-07-Handy-Haversacks/regex"
	"fmt"
	"strconv"
	"strings"
)

const dataPath = "data/input"
const bagsPattern = "^([\\w\\s]*) bags contain ([\\w\\s\\d,]*) bags?\\.$"
const noBagsPattern = "^([\\w\\s]*) bags contain no other bags.$"

type Pair struct {
	name  string
	value int
}

func main() {
	fmt.Println("\n--- Day 7: Handy Haversacks  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("  Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("  Part 2 >> %d", resultPart2))
}

func firstPart(input string) int {
	rules := common.GetStringArrayFromStringInput(input, "\n")
	ruleTree := getResultMap(rules)
	result := 0
	for key := range ruleTree {
		result += getPaths(ruleTree, key, []string{})
	}
	return result
}

func getResultMap(rules []string) map[string][]Pair {
	ruleTree := make(map[string][]Pair)
	for _, rule := range rules {
		noBagsRegex := regex.NewRegex(noBagsPattern)
		bagsRegex := regex.NewRegex(bagsPattern)
		if noBagsRegex.Matches(rule) {
			groups := noBagsRegex.GetGroups(rule)
			ruleTree[groups[1]] = []Pair{}
		} else if bagsRegex.Matches(rule) {
			groups := bagsRegex.GetGroups(rule)
			children := strings.Split(groups[2], ",")
			var ch []Pair
			for _, value := range children {
				value = strings.TrimSpace(value)
				childGroups := regex.NewRegex("(\\d)+ (.*)").GetGroups(value)
				childValue := strings.TrimSpace(childGroups[1])
				childName := strings.TrimSuffix(childGroups[2], "bags")
				childName = strings.TrimSuffix(childName, "bag")
				childName = strings.TrimSpace(childName)
				childValueInt, _ := strconv.Atoi(childValue)
				ch = append(ch, Pair{childName, childValueInt})
			}
			ruleTree[groups[1]] = ch
		}
	}
	return ruleTree
}

func getPaths(m map[string][]Pair, key string, path []string) int {
	children := m[key]
	if len(children) == 0 {
		return 0
	}
	for _, child := range children {
		if child.name == "shiny gold" {
			return 1
		}
		if getPaths(m, child.name, path) == 1 {
			return 1
		}
	}
	return 0
}

func countBags(m map[string][]Pair, key string) int {
	children := m[key]
	if len(children) == 0 {
		return 0
	}
	res := 0
	for _, child := range children {
		res += child.value * (countBags(m, child.name) + 1)
	}
	return res
}

func secondPart(input string) int {
	rules := common.GetStringArrayFromStringInput(input, "\n")
	ruleTree := getResultMap(rules)
	return countBags(ruleTree, "shiny gold")
}
