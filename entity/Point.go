package entity

type Point struct {
	X, Y int
}

func (p *Point) At(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *Point) MoveRight(steps int) {
	p.X += steps
}

func (p *Point) MoveLeft(steps int) {
	p.X -= steps
}

func (p *Point) MoveUp(steps int) {
	p.Y -= steps
}

func (p *Point) MoveDown(steps int) {
	p.Y += steps
}
