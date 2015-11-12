package main

import (
	"github.com/nsf/termbox-go"
)

import (
	"github.com/khueue/go-snake/entity"
)

// World is
type World struct {
	player entity.Snake
}

// NewWorld is
func NewWorld() *World {
	return &World{
		player: entity.Snake{
			X:         1,
			Y:         1,
			Direction: entity.DirectionDown,
		},
	}
}

// ProcessEvent is
func (w *World) ProcessEvent(event termbox.Event) {
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyArrowUp:
			w.player.Direction = entity.DirectionUp
		case termbox.KeyArrowDown:
			w.player.Direction = entity.DirectionDown
		case termbox.KeyArrowLeft:
			w.player.Direction = entity.DirectionLeft
		case termbox.KeyArrowRight:
			w.player.Direction = entity.DirectionRight
		}
	}
}

// Step aoeu
func (w *World) Step() {
	w.player.Step()
}

// Draw aoeu
func (w *World) Draw() {
	w.player.Draw()
}
