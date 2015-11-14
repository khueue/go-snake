// go:generate $GOPATH/bin/stringer -type Direction

package entity

type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)
