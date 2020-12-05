package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "FBFBBFFRLR\nBFFFBBFRRR\nFFFBBBFRRR\nBBFFBBFRLL"

	// when
	result := firstPart(input)

	// then
	assert.Equal(t, 820, result)
}
