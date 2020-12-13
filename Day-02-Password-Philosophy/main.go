package main

import (
	"Day-02-Password-Philosophy/common"
	"fmt"
	"github.com/jucardi/go-streams/streams"
	"index/suffixarray"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const dataPath = "data/input"
const passPattern = "(\\d+)-(\\d+) (\\w): (\\w+)"

var passRegex = regexp.MustCompile(passPattern)

func main() {
	fmt.Println("\n--- Day 2: Password Philosophy  ---")
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

// TODO: Regexp module
func firstPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	result := 0
	for _, passPolicy := range lines {
		groups := passRegex.FindAllStringSubmatch(passPolicy, -1)[0]
		lowerRange, _ := strconv.Atoi(groups[1])
		higherRange, _ := strconv.Atoi(groups[2])
		searchingCharacter := groups[3]
		password := groups[4]
		if strings.Contains(password, searchingCharacter) {
			characterRegex := regexp.MustCompile(searchingCharacter)
			index := suffixarray.New([]byte(password))
			numberOfOccurrence := len(index.FindAllIndex(characterRegex, -1))
			if numberOfOccurrence >= lowerRange && numberOfOccurrence <= higherRange {
				result++
			}
		}
	}
	return result
}

func secondPart(input string) int {
	lines := common.GetStringArrayFromStringInput(input, "\n")
	result := 0
	for _, passPolicy := range lines {
		groups := passRegex.FindAllStringSubmatch(passPolicy, -1)[0]
		lowerRange, _ := strconv.Atoi(groups[1])
		higherRange, _ := strconv.Atoi(groups[2])
		searchingCharacter := groups[3]
		password := groups[4]
		if strings.Contains(password, searchingCharacter) {
			characterRegex := regexp.MustCompile(searchingCharacter)
			index := suffixarray.New([]byte(password))
			characterPositions := index.FindAllIndex(characterRegex, -1)
			if charactersInPositions := streams.
				FromArray(characterPositions).
				Filter(passWithCharsInPositions(lowerRange, higherRange)).
				Count(); charactersInPositions == 1 {
				result++
			}
		}
	}
	return result
}

func passWithCharsInPositions(lowerRange int, higherRange int) func(v interface{}) bool {
	return func(v interface{}) bool {
		return v.([]int)[0] == lowerRange-1 || v.([]int)[0] == higherRange-1
	}
}
