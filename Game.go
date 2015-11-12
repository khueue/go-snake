package main

import (
	"time"
)

import (
	"github.com/nsf/termbox-go"
)

// Game is xxx
type Game struct {
}

// Run is
func (g *Game) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	world := NewWorld()
	quitChan := make(chan bool)
	eventChan := make(chan termbox.Event)

	// Poll for terminal events.
	go func() {
		for {
			event := termbox.PollEvent()
			eventChan <- event
		}
	}()

	// React to terminal events.
	go func() {
		for {
			select {
			case event := <-eventChan:
				switch event.Type {
				case termbox.EventKey:
					switch event.Key {
					case termbox.KeyEsc:
						quitChan <- true
					default:
						world.ProcessEvent(event)
					}
				}
			}
		}
	}()

	// Step the world.
	go func() {
		for {
			world.Step()
			termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
			world.Draw()
			termbox.Flush()
			time.Sleep(time.Duration(1000000/10) * time.Microsecond)
		}
	}()

	// s := entities.Snake{X: 1, Y: 2}
	// sp := entities.SnakePart{X: 10, Y: 20}
	// const color = termbox.ColorDefault
	// termbox.SetCell(s.GetX(), s.Y, 'S', color, color)
	// termbox.SetCell(sp.X, sp.Y, 'P', color, color)

	// Block until exit is requested.
	<-quitChan
}
