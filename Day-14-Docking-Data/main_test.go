package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0\n"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 165, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "mask = 000000000000000000000000000000X1001X\nmem[42] = 100\nmask = 00000000000000000000000000000000X0XX\nmem[26] = 1"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 208, result)
}
