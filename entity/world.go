package entity

import (
	"math/rand"
)

import (
	"github.com/nsf/termbox-go"
)

// World xxx
type World struct {
	player Snake
	foods  []*Food
}

// Init xxx
func (w *World) Init() {
	w.player = Snake{
		Position: Point{
			X: 1,
			Y: 1,
		},
		Direction: DirectionDown,
	}
	w.foods = []*Food{}
	w.spawnFood()
	w.spawnFood()
}

func (w *World) spawnFood() {
	food := &Food{
		Position: Point{
			X: rand.Intn(80),
			Y: rand.Intn(40),
		},
	}
	w.foods = append(w.foods, food)
}

func (w *World) removeFoodAtIndex(i int) {
	w.foods = append(w.foods[:i], w.foods[i+1:]...)
}

// ProcessEvent xxx
func (w *World) ProcessEvent(event *termbox.Event) {
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyArrowUp:
			if w.player.Direction != DirectionDown {
				w.player.Direction = DirectionUp
			}
		case termbox.KeyArrowDown:
			if w.player.Direction != DirectionUp {
				w.player.Direction = DirectionDown
			}
		case termbox.KeyArrowLeft:
			if w.player.Direction != DirectionRight {
				w.player.Direction = DirectionLeft
			}
		case termbox.KeyArrowRight:
			if w.player.Direction != DirectionLeft {
				w.player.Direction = DirectionRight
			}
		}
	}
}

// Step xxx
func (w *World) Step() {
	w.player.Step()

	for i, food := range w.foods {
		if food.At(w.player.Position) {
			w.player.EatFood(food)
			food.Kill()
			w.removeFoodAtIndex(i)
			w.spawnFood()
		}
	}
}

// Render xxx
func (w *World) Render() {
	w.player.Render()
	for _, food := range w.foods {
		food.Render()
	}
}
