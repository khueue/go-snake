package entity

// Direction represents a direction: up, down, left, right.
type Direction int

// Values for Direction.
const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)
