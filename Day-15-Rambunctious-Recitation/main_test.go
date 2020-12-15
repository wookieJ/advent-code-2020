package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "0,3,6"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 436, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "0,3,6"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 175594, result)
}
