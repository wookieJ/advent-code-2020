package point

import "math"

type Point struct {
	X int
	Y int
}

type PointF struct {
	X float64
	Y float64
}

func (p *Point) IncX() {
	p.X += 1
}

func (p *PointF) IncX() {
	p.X += 1
}

func (p *Point) IncY() {
	p.X += 1
}

func (p *PointF) IncY() {
	p.X += 1
}

func (p *Point) AddX(value int) {
	p.X += value
}

func (p *PointF) AddX(value float64) {
	p.X += value
}

func (p *Point) AddY(value int) {
	p.Y += value
}

func (p *PointF) AddY(value float64) {
	p.Y += value
}

func (p *Point) Add(x, y int) {
	p.AddX(x)
	p.AddY(y)
}

func (p *PointF) Add(x, y float64) {
	p.AddX(x)
	p.AddY(y)
}

func (p *Point) Same(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *PointF) Same(p2 PointF, precision float64) bool {
	return p.X >= p2.X-precision && p.X <= p2.X+precision && p.Y >= p2.Y-precision && p.Y <= p2.Y+precision
}

func (p *Point) DirectionVector(precision float64) Point {
	module := p.Module()
	if module == 0 {
		return *p
	}
	dx := sign(float64(p.X)/module, precision)
	dy := sign(float64(p.Y)/module, precision)
	return Point{int(dx), int(dy)}
}

func (p *PointF) UnitVector() PointF {
	module := p.Module()
	if module == 0 {
		return *p
	}
	return PointF{p.X / module, p.Y / module}
}

func sign(dx, precision float64) float64 {
	if dx > -1*precision && dx <= precision {
		dx = 0
	} else if dx > 0 {
		dx = 1
	} else {
		dx = -1
	}
	return dx
}

func (p *Point) Module() float64 {
	return math.Sqrt(math.Pow(float64(p.X), 2) + math.Pow(float64(p.Y), 2))
}

func (p *PointF) Module() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}
