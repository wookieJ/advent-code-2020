package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "1\n1\n1"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 0, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "1\n1\n1"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 0, result)
}
