package point

type Point struct {
	X int
	Y int
}

func (p *Point) IncX() {
	p.X += 1
}

func (p *Point) IncY() {
	p.X += 1
}

func (p *Point) AddX(value int) {
	p.X += value
}

func (p *Point) AddY(value int) {
	p.Y += value
}

func (p *Point) Add(x, y int) {
	p.AddX(x)
	p.AddY(y)
}

func (p *Point) Same(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}
