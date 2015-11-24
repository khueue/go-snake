package entity

import "github.com/nsf/termbox-go"

// Snake xxx
type Snake struct {
	Parts     []*SnakePart
	VelocityX float64
	VelocityY float64
}

// SnakePart xxx
type SnakePart struct {
	Position      Point
	CurrDirection Direction
	NextDirection Direction
}

// Step xxx
func (s *Snake) Step(dt float64) {
	head := s.Parts[0]
	head.step(s.VelocityX, dt, nil)

	next := head
	for _, part := range s.Parts[1:] {
		part.step(s.VelocityX, dt, next)
		next = part
	}
}
func (part *SnakePart) step(velocity float64, dt float64, ahead *SnakePart) {
	distance := velocity * dt
	if ahead == nil || ahead != nil && !part.Position.At(ahead.Position) {
		if part.NextDirection != DirectionNone {
			switch part.NextDirection {
			case DirectionUp:
				part.Position.MoveUp(distance)
			case DirectionDown:
				part.Position.MoveDown(distance)
			case DirectionLeft:
				part.Position.MoveLeft(distance)
			case DirectionRight:
				part.Position.MoveRight(distance)
			default:
				panic("should never happen")
			}
		}
	}

	part.CurrDirection = part.NextDirection
	if ahead != nil {
		part.NextDirection = ahead.CurrDirection
	}
}

// ChangeDirection xxx
func (s *Snake) ChangeDirection(nextDirection Direction) {
	head := s.Parts[0]
	switch nextDirection {
	case DirectionUp:
		if head.CurrDirection != DirectionDown {
			head.NextDirection = DirectionUp
		}
	case DirectionDown:
		if head.CurrDirection != DirectionUp {
			head.NextDirection = DirectionDown
		}
	case DirectionLeft:
		if head.CurrDirection != DirectionRight {
			head.NextDirection = DirectionLeft
		}
	case DirectionRight:
		if head.CurrDirection != DirectionLeft {
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
			CurrDirection: DirectionRight,
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
		termbox.SetCell(int(part.Position.X), int(part.Position.Y), face, color, termbox.ColorDefault)
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
	termbox.SetCell(int(head.Position.X), int(head.Position.Y), face, color, termbox.ColorDefault)
}
