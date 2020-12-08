package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 5, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"

	// when
	result := secondPart(input)

	// then
	assert.Equal(t, 8, result)
}
