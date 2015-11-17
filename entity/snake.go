package entity

import (
	"github.com/nsf/termbox-go"
)

// Snake xxx
type Snake struct {
	Parts []*SnakePart
	Speed int
}

// type SnakeParts []*SnakePart
//
// func (b *SnakeParts) First() *SnakePart {
// 	return (*b)[0]
// }
//
// func (b *SnakeParts) Tail() SnakeParts {
// 	return (*b)[1:]
// }
//
// func (b *SnakeParts) Last() *SnakePart {
// 	return (*b)[len(*b)-1]
// }
//
// func (b *SnakeParts) Attach(part *SnakePart) {
// 	*b = append(*b, part)
// }

// SnakePart xxx
type SnakePart struct {
	Position      Point
	PrevDirection Direction
	NextDirection Direction
}

// Step xxx
func (s *Snake) Step() {
	head := s.Parts[0]
	head.step(s.Speed, nil)

	next := head
	for _, part := range s.Parts[1:] {
		part.step(s.Speed, next)
		next = part
	}
}

func (part *SnakePart) step(speed int, ahead *SnakePart) {
	if part.NextDirection != DirectionNone {
		switch part.NextDirection {
		case DirectionUp:
			part.Position.MoveUp(speed)
		case DirectionDown:
			part.Position.MoveDown(speed)
		case DirectionLeft:
			part.Position.MoveLeft(speed)
		case DirectionRight:
			part.Position.MoveRight(speed)
		default:
			panic("Should never happen")
		}
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
	for i := 0; i < food.Energy; i++ {
		lastPart := s.Parts[len(s.Parts)-1]
		newPart := &SnakePart{
			Position: Point{
				X: lastPart.Position.X,
				Y: lastPart.Position.Y,
			},
			PrevDirection: DirectionNone,
			NextDirection: DirectionNone,
		}
		s.Attach(newPart)
	}
}

// Attach xxx
func (s *Snake) Attach(part *SnakePart) {
	s.Parts = append(s.Parts, part)
}

// Render xxx
func (s *Snake) Render() {
	var face rune

	for i := len(s.Parts) - 1; i >= 1; i-- {
		part := s.Parts[i]
		face = 'o'
		color := termbox.ColorYellow
		termbox.SetCell(part.Position.X, part.Position.Y, face, color, termbox.ColorDefault)
	}

	color := termbox.ColorYellow
	head := s.Parts[0]
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
}
