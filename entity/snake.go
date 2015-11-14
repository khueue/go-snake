package entity

import (
	"github.com/nsf/termbox-go"
)

type Snake struct {
	Position  Point
	Direction Direction
}

func (s *Snake) Step() {
	switch s.Direction {
	case DirectionUp:
		s.Position.MoveUp(1)
	case DirectionDown:
		s.Position.MoveDown(1)
	case DirectionLeft:
		s.Position.MoveLeft(1)
	default:
		s.Position.MoveRight(1)
	}
}

func (s *Snake) EatFood(food *Food) {
	// Grow.
}

func (s *Snake) Render() {
	termbox.SetCell(s.Position.X, s.Position.Y, '*', termbox.ColorRed, termbox.ColorDefault)
}
