package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 2, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 1, result)
}
