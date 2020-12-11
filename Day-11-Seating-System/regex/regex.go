package regex

import (
	"index/suffixarray"
	"regexp"
)

type Regex struct {
	pattern string
	regexp  regexp.Regexp
}

func NewRegex(pattern string) Regex {
	regex := Regex{pattern: pattern}
	regex.regexp = *regexp.MustCompile(pattern)
	return regex
}

func (r Regex) Matches(word string) bool {
	if matches := r.regexp.FindAllStringSubmatch(word, -1); matches != nil {
		return len(matches) > 0
	}
	return false
}

func (r Regex) AllMatches(word string) [][]string {
	return r.regexp.FindAllStringSubmatch(word, -1)
}

func (r Regex) AllFlattenMatches(word string) []string {
	result := make([]string, 0)
	allMatches := r.AllMatches(word)
	for _, match := range allMatches {
		for _, subMatch := range match {
			result = append(result, subMatch)
		}
	}
	return result
}

func (r Regex) NumberOfOccurrence(word string) int {
	index := suffixarray.New([]byte(word))
	return len(index.FindAllIndex(&r.regexp, -1))
}

func (r Regex) AllOccurrencesIndexes(word string) [][]int {
	index := suffixarray.New([]byte(word))
	return index.FindAllIndex(&r.regexp, -1)
}

func (r Regex) GetGroups(word string) []string {
	return r.regexp.FindAllStringSubmatch(word, -1)[0]
}
