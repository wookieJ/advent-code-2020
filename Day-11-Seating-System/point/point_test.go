package point

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldAddPointValues(t *testing.T) {
	// given
	point := Point{}

	// when
	point.AddX(5)
	point.AddY(1)
	point.AddX(2)
	point.Add(-1, -4)

	// then
	assert.Equal(t, Point{6, -3}, point)
}

func TestShouldCheckIfPointAreTheSame(t *testing.T) {
	// given
	point1 := Point{}
	point2 := Point{2, 3}
	point3 := Point{3, 2}
	point4 := Point{0, 0}

	// when
	same := point1.Same(point4)
	notSame := point1.Same(point2)
	notSame2 := point2.Same(point3)

	// then
	assert.True(t, same)
	assert.False(t, notSame)
	assert.False(t, notSame2)
}
