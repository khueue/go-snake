package entity

import (
	"github.com/nsf/termbox-go"
)

// Snake is a thingy
type Snake struct {
	X, Y int
	Direction int
}

// GetX does cool stuff
func (s *Snake) GetX() int {
	return s.X
}

// Step does cool stuff
func (s *Snake) Step() {
	switch s.Direction {
	case DirectionUp: s.Y -= 1
	case DirectionDown: s.Y += 1
	case DirectionLeft: s.X -= 2
	default: s.X += 2
	}
}

// Draw does cool stuff
func (s *Snake) Draw() {
	termbox.SetCell(s.X, s.Y, '*', termbox.ColorRed, termbox.ColorDefault)
}
