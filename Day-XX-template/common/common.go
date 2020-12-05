package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetInputFromFile(path string) string {
	data, err := getFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func getFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func GetIntArrayFromStringInput(input, sep string) []int {
	lines := strings.Split(strings.TrimSpace(input), sep)
	var values []int
	for _, value := range lines {
		tmp, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unexpected input string %v", value))
		}
		values = append(values, tmp)
	}
	return values
}

func GetStringArrayFromStringInput(input, sep string) []string {
	return strings.Split(strings.TrimSpace(input), sep)
}

func GetArrayOfMapsFromString(input, arrayDelimiter string, mapDelimiters []string, fieldsDelimiter string) []map[string]string {
	var result []map[string]string
	mapStrings := GetStringArrayFromStringInput(input, arrayDelimiter)
	for _, mapString := range mapStrings {
		resultMap := make(map[string]string)
		fields := SplitAny(mapString, mapDelimiters)
		for _, fieldString := range fields {
			fieldSplit := GetStringArrayFromStringInput(fieldString, fieldsDelimiter)
			if _, ok := resultMap[fieldSplit[0]]; !ok {
				resultMap[fieldSplit[0]] = fieldSplit[1]
			}
		}
		result = append(result, resultMap)
	}
	return result
}

func SplitAny(s string, delimiters []string) []string { // TODO: tests
	seps := strings.Join(delimiters, "")
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func HaveSameElements(array1, array2 []string) bool { // TODO: tests
	if len(array1) != len(array2) {
		return false
	}
	for _, v := range array1 {
		if !ArrayContains(array2, v) {
			return false
		}
	}
	return true
}

func ArrayContains(array []string, element string) bool { // TODO: tests
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

func MapKeys(m map[string]string) []string { // TODO: tests
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}

func HaveAllKeys(m map[string]string, keys []string) bool {
	if len(m) != len(keys) {
		return false
	}
	for k, _ := range m {
		if !ArrayContains(keys, k) {
			return false
		}
	}
	return true
}
