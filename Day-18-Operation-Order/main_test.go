package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "1 + 2 * 3 + 4 * 5 + 6"
	input2 := "2 * 3 + (4 * 5)"
	input3 := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	input4 := "1 + 2 * 3 + 4 * 5 + 6\n((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	input5 := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"

	// when
	result := firstPart(input)
	result2 := firstPart(input2)
	result3 := firstPart(input3)
	result4 := firstPart(input4)
	result5 := firstPart(input5)

	// then
	assert.Equal(t, 71, result)
	assert.Equal(t, 26, result2)
	assert.Equal(t, 13632, result3)
	assert.Equal(t, 13703, result4)
	assert.Equal(t, 12240, result5)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "1 + 2 * 3 + 4 * 5 + 6"
	input2 := "2 * 3 + (4 * 5)"
	input3 := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	input4 := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"

	// when
	result := secondPart(input)
	result2 := secondPart(input2)
	result3 := secondPart(input3)
	result4 := secondPart(input4)

	// then
	assert.Equal(t, 231, result)
	assert.Equal(t, 46, result2)
	assert.Equal(t, 23340, result3)
	assert.Equal(t, 669060, result4)
}
