package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"

	// when
	result := firstPart(input, 5)

	// then
	assert.Equal(t, 127, result)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"

	// when
	result := secondPart(input, 5)

	// then
	assert.Equal(t, 62, result)
}
