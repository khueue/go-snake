package entity

// Direction represents a direction: up, down, left, right.
type Direction int

// Values for Direction.
const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)
