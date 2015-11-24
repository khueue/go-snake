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
func (p *Point) MoveRight(distance float64) {
	p.X += distance
}

// MoveLeft xxx
func (p *Point) MoveLeft(distance float64) {
	p.X -= distance
}

// MoveUp xxx
func (p *Point) MoveUp(distance float64) {
	p.Y -= distance
}

// MoveDown xxx
func (p *Point) MoveDown(distance float64) {
	p.Y += distance
}
