package entity

import (
	"github.com/nsf/termbox-go"
)

// Snake xxx
type Snake struct {
	Parts []*SnakePart
}

// SnakePart xxx
type SnakePart struct {
	Position      Point
	PrevDirection Direction
	NextDirection Direction
}

// Step xxx
func (s *Snake) Step() {
	head := s.Parts[0]
	head.step(nil)

	next := head
	for _, part := range s.Parts[1:] {
		part.step(next)
		next = part
	}
}

func (part *SnakePart) step(ahead *SnakePart) {
	switch part.NextDirection {
	case DirectionUp:
		part.Position.MoveUp(1)
	case DirectionDown:
		part.Position.MoveDown(1)
	case DirectionLeft:
		part.Position.MoveLeft(1)
	case DirectionRight:
		part.Position.MoveRight(1)
	default:
		panic("Should never happen")
	}

	part.PrevDirection = part.NextDirection
	if ahead != nil {
		part.NextDirection = ahead.PrevDirection
	}
}

// ChangeDirection xxx
func (s *Snake) ChangeDirection(nextDirection Direction) {
	head := s.Parts[0]
	switch nextDirection {
	case DirectionUp:
		if head.PrevDirection != DirectionDown {
			head.NextDirection = DirectionUp
		}
	case DirectionDown:
		if head.PrevDirection != DirectionUp {
			head.NextDirection = DirectionDown
		}
	case DirectionLeft:
		if head.PrevDirection != DirectionRight {
			head.NextDirection = DirectionLeft
		}
	case DirectionRight:
		if head.PrevDirection != DirectionLeft {
			head.NextDirection = DirectionRight
		}
	}
}

// Position xxx
func (s *Snake) Position() Point {
	return s.Parts[0].Position
}

// EatFood xxx
func (s *Snake) EatFood(food *Food) {
	lastPart := s.Parts[len(s.Parts)-1]
	var x, y int
	switch lastPart.PrevDirection {
	case DirectionUp:
		x = lastPart.Position.X
		y = lastPart.Position.Y + 1
	case DirectionDown:
		x = lastPart.Position.X
		y = lastPart.Position.Y - 1
	case DirectionLeft:
		x = lastPart.Position.X + 1
		y = lastPart.Position.Y
	case DirectionRight:
		x = lastPart.Position.X - 1
		y = lastPart.Position.Y
	}
	newPart := &SnakePart{
		Position: Point{
			X: x,
			Y: y,
		},
		NextDirection: lastPart.PrevDirection,
	}
	s.Attach(newPart)
}

// Attach xxx
func (s *Snake) Attach(part *SnakePart) {
	s.Parts = append(s.Parts, part)
}

// Render xxx
func (s *Snake) Render() {
	head := s.Parts[0]
	color := termbox.ColorYellow
	var face rune
	switch head.NextDirection {
	case DirectionUp:
		face = '^'
	case DirectionDown:
		face = 'v'
	case DirectionLeft:
		face = '<'
	case DirectionRight:
		face = '>'
	}
	termbox.SetCell(head.Position.X, head.Position.Y, face, color, termbox.ColorDefault)

	for _, part := range s.Parts[1:] {
		face = 'o'
		color = termbox.ColorYellow
		termbox.SetCell(part.Position.X, part.Position.Y, face, color, termbox.ColorDefault)
	}
}
