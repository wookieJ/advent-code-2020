package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "939\n7,13,x,x,59,x,31,19"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 295, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "1\n17,x,13,19"
	input2 := "1\n67,7,59,61"
	input3 := "1\n67,x,7,59,61"
	input4 := "1\n67,7,x,59,61"
	input5 := "1\n1789,37,47,1889"
	input6 := "939\n7,13,x,x,59,x,31,19"

	// when
	result := secondPart(input)
	result2 := secondPart(input2)
	result3 := secondPart(input3)
	result4 := secondPart(input4)
	result5 := secondPart(input5)
	result6 := secondPart(input6)

	// then
	assert.Equal(t, 3417, result)
	assert.Equal(t, 754018, result2)
	assert.Equal(t, 779210, result3)
	assert.Equal(t, 1261476, result4)
	assert.Equal(t, 1202161486, result5)
	assert.Equal(t, 1068781, result6)
}

func TestShouldGetSecondExample2(t *testing.T) {
	// given
	input := "1\n2,3,4,5"
	input2 := "1\n2,x,4,5"
	input3 := "1\n2,x,4,x"
	input4 := "1\n2,x,x,5,6"
	input5 := "1\n2,x,4,x,6"

	// when
	result := secondPart(input)
	result2 := secondPart(input2)
	result3 := secondPart(input3)
	result4 := secondPart(input4)
	result5 := secondPart(input5)

	// then
	assert.Equal(t, 2, result)
	assert.Equal(t, 2, result2)
	assert.Equal(t, 2, result3)
	assert.Equal(t, 2, result4)
	assert.Equal(t, 2, result5)
}
