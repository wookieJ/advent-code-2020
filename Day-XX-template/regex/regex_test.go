package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCheckIfStringMatches(t *testing.T) {
	// given
	pattern := "test\\d+"
	regex := NewRegex(pattern)

	// when
	matches := regex.Matches("test01")

	// then
	assert.True(t, matches)
}

func TestShouldGetAllMatches(t *testing.T) {
	// given
	pattern := "a(x*)b"
	regex := NewRegex(pattern)

	// when
	allMatches1 := regex.AllMatches("ab")
	allMatches2 := regex.AllMatches("axb")
	allMatches3 := regex.AllMatches("axxxb")
	allMatches4 := regex.AllMatches("axb-ab-xx")

	// then
	assert.Equal(t, [][]string{{"ab", ""}}, allMatches1)
	assert.Equal(t, [][]string{{"axb", "x"}}, allMatches2)
	assert.Equal(t, [][]string{{"axxxb", "xxx"}}, allMatches3)
	assert.Equal(t, [][]string{{"axb", "x"}, {"ab", ""}}, allMatches4)
}

func TestShouldGetAllFlattenMatches(t *testing.T) {
	// given
	pattern := "a(x*)b"
	pattern2 := "a+"
	regex := NewRegex(pattern)
	regex2 := NewRegex(pattern2)

	// when
	allMatches1 := regex.AllFlattenMatches("ab")
	allMatches2 := regex.AllFlattenMatches("axb")
	allMatches3 := regex.AllFlattenMatches("axxxb")
	allMatches4 := regex.AllFlattenMatches("axb-ab-xx")
	allMatches5 := regex2.AllFlattenMatches("ala ma kota")

	// then
	assert.Equal(t, []string{"ab", ""}, allMatches1)
	assert.Equal(t, []string{"axb", "x"}, allMatches2)
	assert.Equal(t, []string{"axxxb", "xxx"}, allMatches3)
	assert.Equal(t, []string{"axb", "x", "ab", ""}, allMatches4)
	assert.Equal(t, []string{"a", "a", "a", "a"}, allMatches5)
}

func TestShouldGetNumberOfOccurrence(t *testing.T) {
	// given
	pattern := "a"
	regex := NewRegex(pattern)

	// when
	number := regex.NumberOfOccurrence("ab")
	number2 := regex.NumberOfOccurrence("aa")
	number3 := regex.NumberOfOccurrence("aba")
	number4 := regex.NumberOfOccurrence("abaaa awd   -?a")

	// then
	assert.Equal(t, 1, number)
	assert.Equal(t, 2, number2)
	assert.Equal(t, 2, number3)
	assert.Equal(t, 6, number4)
}

func TestShouldGetAllOccurrencesIndexes(t *testing.T) {
	// given
	pattern := "a"
	regex := NewRegex(pattern)

	// when
	number := regex.AllOccurrencesIndexes("ab")
	number2 := regex.AllOccurrencesIndexes("aa")
	number3 := regex.AllOccurrencesIndexes("aba")
	number4 := regex.AllOccurrencesIndexes("abaaa awd   -?a")

	// then
	assert.Equal(t, [][]int{{0, 1}}, number)
	assert.Equal(t, [][]int{{0, 1}, {1, 2}}, number2)
	assert.Equal(t, [][]int{{0, 1}, {2, 3}}, number3)
	assert.Equal(t, [][]int{{0, 1}, {2, 3}, {3, 4}, {4, 5}, {6, 7}, {14, 15}}, number4)
}

func TestShouldGetGroups(t *testing.T) {
	// given
	pattern := "(\\d+)-(\\d+)-(\\d+)"
	regex := NewRegex(pattern)

	// when
	number := regex.GetGroups("01-01-1999")
	number2 := regex.GetGroups("1-2-3")

	// then
	assert.Equal(t, []string{"01-01-1999", "01", "01", "1999"}, number)
	assert.Equal(t, []string{"1-2-3", "1", "2", "3"}, number2)
}
