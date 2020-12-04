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
	passports := common.GetStringArrayFromStringInput(input, "\n\n")
	result := 0
	for _, passport := range passports {
		maps := getPassportMapFromString(passport)
		if len(maps) == len(requiredFields) {
			result++
		}
	}
	return result
}

func secondPart(input string) int {
	passports := common.GetStringArrayFromStringInput(input, "\n\n")
	result := 0
	for _, passportString := range passports {
		passport := getPassportMapFromString(passportString)
		passportValidFields := 0
		if len(passport) == len(requiredFields) {
			for key, value := range passport {
				switch key {
				case "byr":
					if isInRange(value, 1920, 2002) {
						passportValidFields++
					}
				case "iyr":
					if isInRange(value, 2010, 2020) {
						passportValidFields++
					}
				case "eyr":
					if isInRange(value, 2020, 2030) {
						passportValidFields++
					}
				case "hgt":
					if strings.Contains(value, "cm") {
						if isInRange(strings.Replace(value, "cm", "", 1), 150, 193) {
							passportValidFields++
						}
					} else if strings.Contains(value, "in") {
						if isInRange(strings.Replace(value, "in", "", 1), 59, 76) {
							passportValidFields++
						}
					}
				case "hcl":
					r := regex.NewRegex("#[0-9a-f]{6}")
					if r.Matches(value) {
						passportValidFields++
					}
				case "ecl":
					r := regex.NewRegex("amb|blu|brn|gry|grn|hzl|oth")
					if r.Matches(value) {
						passportValidFields++
					}
				case "pid":
					r := regex.NewRegex("\\d{9}")
					if r.Matches(value) {
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

func getPassportMapFromString(passportString string) map[string]string {
	passport := make(map[string]string)
	newLineFields := common.GetStringArrayFromStringInput(passportString, "\n")
	var fields []string
	for _, newLineField := range newLineFields {
		spaceFields := common.GetStringArrayFromStringInput(newLineField, " ")
		for _, spaceField := range spaceFields {
			fields = append(fields, spaceField)
		}
	}
	for _, field := range fields {
		keyValueField := common.GetStringArrayFromStringInput(field, ":")
		for _, req := range requiredFields {
			if req == keyValueField[0] {
				if val, ok := passport[keyValueField[0]]; ok {
					println("Already exists: ", keyValueField[0], ":", val)
				} else {
					passport[keyValueField[0]] = keyValueField[1]
				}
			}
		}
	}
	return passport
}

func isInRange(value string, low, high int) bool {
	if v, err := strconv.Atoi(value); err == nil {
		if v >= low && v <= high {
			return true
		}
	}
	return false
}
