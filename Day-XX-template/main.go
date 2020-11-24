package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const dataPath = "data/input"

func main() {
	fmt.Println("--- _DAY_DESC_  ---")
	input := getInputFromFile()

	resultPart1 := firstPart(input)
	resultPart2 := secondPart(input)

	fmt.Println(fmt.Sprintf("Part 1 >> %d", resultPart1))
	fmt.Println(fmt.Sprintf("Part 2 >> %d", resultPart2))
}

func firstPart(input string) int {
	return 0
}

func secondPart(input string) int {
	return 0
}

func getInputFromFile() string {
	data, err := getFile(dataPath)
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
