package entity

// Point is a position in X, Y.
type Point struct {
	X, Y float64
}

// At xxx
func (p *Point) At(p2 Point) bool {
	return int(p.X) == int(p2.X) && int(p.Y) == int(p2.Y)
}

// MoveRight xxx
func (p *Point) MoveRight(steps int) {
	p.X += float64(steps)
}

// MoveLeft xxx
func (p *Point) MoveLeft(steps int) {
	p.X -= float64(steps)
}

// MoveUp xxx
func (p *Point) MoveUp(steps int) {
	p.Y -= float64(steps)
}

// MoveDown xxx
func (p *Point) MoveDown(steps int) {
	p.Y += float64(steps)
}
