package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 11, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 6, result)
}
