package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	// given
	number := 2

	// when
	result := number * number

	// then
	assert.Equal(t, 4, result)
}
