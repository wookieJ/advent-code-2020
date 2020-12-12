package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "F10\nN3\nF7\nR90\nF11"
	input2 := "F10\nL180\nF9"

	// when
	result := firstPart(input)
	result2 := firstPart(input2)

	// then
	assert.Equal(t, 25, result)
	assert.Equal(t, 1, result2)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "F10\nN3\nF7\nR90\nF11"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 286, result)
}
