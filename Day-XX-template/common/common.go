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
