package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	numbers := "1721\n979\n366\n299\n675\n1456"

	// when
	result := firstPart(numbers)

	// then
	assert.Equal(t, 514579, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	numbers := "1721\n979\n366\n299\n675\n1456"

	// when
	result := secondPart(numbers)

	// then
	assert.Equal(t, 241861950, result)
}
