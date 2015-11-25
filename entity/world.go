package entity

import (
	"math/rand"

	"github.com/khueue/go-snake/config"
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
		Velocity: 1,
	}
	w.player.Attach(&SnakePart{
		Position: Point{
			X: 10,
			Y: 10,
		},
		PrevDirection: DirectionNone,
		NextDirection: DirectionRight,
	})
	w.foods = []*Food{}
	for i := 0; i < config.StartingFood; i++ {
		w.spawnFood()
	}
}

func (w *World) spawnFood() {
	food := &Food{
		Position: Point{
			X: rand.Intn(100),
			Y: rand.Intn(40),
		},
		Energy: config.FoodEnergy,
	}
	w.foods = append(w.foods, food)
}

func (w *World) removeFoodAtIndex(i int) {
	w.foods = append(w.foods[:i], w.foods[i+1:]...)
}

// ProcessEvent xxx
func (w *World) ProcessEvent(event termbox.Event) {
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyArrowUp:
			w.player.ChangeDirection(DirectionUp)
		case termbox.KeyArrowDown:
			w.player.ChangeDirection(DirectionDown)
		case termbox.KeyArrowLeft:
			w.player.ChangeDirection(DirectionLeft)
		case termbox.KeyArrowRight:
			w.player.ChangeDirection(DirectionRight)
		}
	}
}

// Step xxx
func (w *World) Step() {
	w.player.Step()

	for i, food := range w.foods {
		if food.At(w.player.Position()) {
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
