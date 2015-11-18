package entity

import (
	"github.com/nsf/termbox-go"
)

// Food is an object that a snake can eat.
type Food struct {
	Position Point
	Energy   int
	Eaten    bool
}

// Kill marks the food as removed.
func (f *Food) Kill() {
	f.Eaten = true
}

// At returns true if it's in the same position as point.
func (f *Food) At(point Point) bool {
	return f.Position.At(point)
}

// Render draws the food on screen.
func (f *Food) Render() {
	if !f.Eaten {
		termbox.SetCell(int(f.Position.X), int(f.Position.Y), '$', termbox.ColorGreen, termbox.ColorDefault)
	}
}
