package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 306, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 291, result)
}

func TestShouldGetSecondExampleWithLoop(t *testing.T) {
	// given
	input := "Player 1:\n43\n19\n\nPlayer 2:\n2\n29\n14"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 0, result)
}
