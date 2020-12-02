package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCheckIfStringMatches(t *testing.T) {
	// given
	pattern := "test\\d+"
	regex := NewRegex(pattern)

	// when
	matches := regex.Matches("test01")

	// then
	assert.True(t, matches)
}
