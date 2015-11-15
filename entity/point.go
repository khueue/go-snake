package entity

// Point is a position in X, Y.
type Point struct {
	X, Y int
}

// At xxx
func (p *Point) At(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

// MoveRight xxx
func (p *Point) MoveRight(steps int) {
	p.X += steps
}

// MoveLeft xxx
func (p *Point) MoveLeft(steps int) {
	p.X -= steps
}

// MoveUp xxx
func (p *Point) MoveUp(steps int) {
	p.Y -= steps
}

// MoveDown xxx
func (p *Point) MoveDown(steps int) {
	p.Y += steps
}
