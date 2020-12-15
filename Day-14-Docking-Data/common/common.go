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

func SplitAny(s string, delimiters []string) []string {
	seps := strings.Join(delimiters, "")
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func HaveSameElements(array1, array2 []string) bool {
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

func ArrayContains(array []string, element string) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

func IntArrayContains(array []int, element int) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

func MapKeys(m map[string]string) []string {
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

func CompareInt(a interface{}, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) < b.(int) {
		return -1
	} else {
		return 0
	}
}

func CompareStrings(a interface{}, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}

func GetArrayOfArrays(input, firstDelimiter, secondDelimiter string) [][]string {
	var result [][]string
	arrays := GetStringArrayFromStringInput(input, firstDelimiter)
	for _, array := range arrays {
		a := GetStringArrayFromStringInput(array, secondDelimiter)
		var tmp []string
		for _, value := range a {
			tmp = append(tmp, value)
		}
		result = append(result, tmp)
	}
	return result
}

func SplitAndGetAll(array []string, delimiter string) []string {
	set := make(map[string]string)
	for _, value := range array {
		for _, letter := range strings.Split(value, delimiter) {
			set[letter] = ""
		}
	}
	return MapKeys(set)
}

func ToInterfaceArray(array []string) []interface{} {
	result := make([]interface{}, len(array))
	for i, v := range array {
		result[i] = v
	}
	return result
}

func GetArraysIntersection(a1, a2 []string) []string {
	var result = make([]string, 0)
	for _, v := range a1 {
		if ArrayContains(a2, v) {
			result = append(result, v)
		}
	}
	return result
}

func GetIntArraysIntersection(a1, a2 []int) []int {
	var result = make([]int, 0)
	for _, v := range a1 {
		if IntArrayContains(a2, v) {
			result = append(result, v)
		}
	}
	return result
}

func SplitAndGetCommon(array []string, delimiter string) []string {
	var result []string
	for i, value := range array {
		letters := strings.Split(value, delimiter)
		if i == 0 {
			result = letters
		} else {
			result = GetArraysIntersection(result, letters)
		}
	}
	return result
}

func MinMax(array []int) (int, int) {
	if array == nil || len(array) == 0 {
		log.Panic("Empty or nil array")
	}
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func CopyMap(originalMap map[string]string) map[string]string {
	copyMap := make(map[string]string, len(originalMap))
	for key, value := range originalMap {
		copyMap[key] = value
	}
	return copyMap
}

func CopyMapOfMap(originalMap map[string]map[string]string) map[string]map[string]string {
	copyMap := make(map[string]map[string]string, len(originalMap))
	for k1, v1 := range originalMap {
		subMap := make(map[string]string, len(v1))
		for k2, v2 := range v1 {
			subMap[k2] = v2
		}
		copyMap[k1] = subMap
	}
	return copyMap
}

func CopyMapOfIntMap(originalMap map[int]map[int]string) map[int]map[int]string {
	copyMap := make(map[int]map[int]string, len(originalMap))
	for k1, v1 := range originalMap {
		subMap := make(map[int]string, len(v1))
		for k2, v2 := range v1 {
			subMap[k2] = v2
		}
		copyMap[k1] = subMap
	}
	return copyMap
}

func Abs(value int) int {
	if value < 0 {
		return -1 * value
	}
	return value
}

func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Panicf("Cannot convert %s value to int", value)
	}
	return intValue
}

func GCD(a, b int) int { // todo test it
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers []int) int { // todo test it
	a, b := integers[0], integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM([]int{result, integers[i]})
	}

	return result
}
