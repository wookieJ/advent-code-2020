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

func TestShouldCheckIfFloatPointAreTheSame(t *testing.T) {
	// given
	point1 := PointF{}
	point2 := PointF{2, 3}
	point3 := PointF{3, 2}
	point4 := PointF{0, 0}

	// when
	same := point1.Same(point4, 0.001)
	notSame := point1.Same(point2, 0.001)
	notSame2 := point2.Same(point3, 0.001)

	// then
	assert.True(t, same)
	assert.False(t, notSame)
	assert.False(t, notSame2)
}

func TestShouldComputeModule(t *testing.T) {
	// when
	point0 := Point{0, 0}
	point1 := Point{1, 0}
	point2 := Point{0, 1}
	point3 := Point{5, 0}
	point4 := Point{0, -1}

	// then
	assert.Equal(t, 0.0, point0.Module())
	assert.Equal(t, 1.0, point1.Module())
	assert.Equal(t, 1.0, point2.Module())
	assert.Equal(t, 5.0, point3.Module())
	assert.Equal(t, 1.0, point4.Module())
}

func TestShouldComputeUnitVector(t *testing.T) {
	// when
	point0 := Point{0, 0}
	point1 := Point{1, 0}
	point2 := Point{0, 1}
	point3 := Point{5, 0}
	point4 := Point{-5, -5}
	point5 := Point{1, 2}
	point6 := Point{2, 223}
	point7 := Point{-2, 6}
	point8 := Point{-2, -1}
	point9 := Point{0, -1}

	// then
	assert.Equal(t, Point{0, 0}, point0.DirectionVector())
	assert.Equal(t, Point{1, 0}, point1.DirectionVector())
	assert.Equal(t, Point{0, 1}, point2.DirectionVector())
	assert.Equal(t, Point{1, 0}, point3.DirectionVector())
	assert.Equal(t, Point{-1, -1}, point4.DirectionVector())
	assert.Equal(t, Point{1, 1}, point5.DirectionVector())
	assert.Equal(t, Point{1, 1}, point6.DirectionVector())
	assert.Equal(t, Point{-1, 1}, point7.DirectionVector())
	assert.Equal(t, Point{-1, -1}, point8.DirectionVector())
	assert.Equal(t, Point{0, -1}, point9.DirectionVector())
}
