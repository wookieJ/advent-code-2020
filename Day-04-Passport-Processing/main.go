package main

import (
	"Day-04-Passport-Processing/common"
	"Day-04-Passport-Processing/regex"
	"fmt"
	"strconv"
	"strings"
)

const dataPath = "data/input"

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	fmt.Println("\n--- Day 4: Passport Processing  ---")
	input := common.GetInputFromFile(dataPath)

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("  Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("  Part 2 >> %d", resultPart2))
}

func firstPart(input string) int {
	passports := common.GetArrayOfMapsFromString(input, "\n\n", []string{" ", "\n"}, ":")
	result := 0
	for _, passport := range passports {
		delete(passport, "cid")
		if common.HaveAllKeys(passport, requiredFields) {
			result++
		}
	}
	return result
}

func secondPart(input string) int {
	passports := common.GetArrayOfMapsFromString(input, "\n\n", []string{" ", "\n"}, ":")
	result := 0
	for _, passport := range passports {
		delete(passport, "cid")
		passportValidFields := 0
		if common.HaveAllKeys(passport, requiredFields) {
			for key, value := range passport {
				switch key {
				case "byr":
					if isStringNumberInRange(value, 1920, 2002) {
						passportValidFields++
					}
				case "iyr":
					if isStringNumberInRange(value, 2010, 2020) {
						passportValidFields++
					}
				case "eyr":
					if isStringNumberInRange(value, 2020, 2030) {
						passportValidFields++
					}
				case "hgt":
					if strings.Contains(value, "cm") {
						if isStringNumberInRange(strings.Replace(value, "cm", "", 1), 150, 193) {
							passportValidFields++
						}
					} else if strings.Contains(value, "in") {
						if isStringNumberInRange(strings.Replace(value, "in", "", 1), 59, 76) {
							passportValidFields++
						}
					}
				case "hcl":
					if regex.NewRegex("^#[0-9a-f]{6}$").Matches(value) {
						passportValidFields++
					}
				case "ecl":
					if regex.NewRegex("amb|blu|brn|gry|grn|hzl|oth").Matches(value) {
						passportValidFields++
					}
				case "pid":
					if regex.NewRegex("^\\d{9}$").Matches(value) {
						passportValidFields++
					}
				}
			}
			if passportValidFields == len(requiredFields) {
				result++
			}
		}
	}
	return result
}

func isStringNumberInRange(value string, low, high int) bool {
	if v, err := strconv.Atoi(value); err == nil {
		if v >= low && v <= high {
			return true
		}
	}
	return false
}
