package main

import (
	"math/rand"
	"time"
)

import (
	"github.com/nsf/termbox-go"
)

import (
	"github.com/khueue/go-snake/entity"
)

// Game is the book keeper of everything.
type Game struct {
}

// Run creates and runs the game. Runs until user quits.
func (g *Game) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	rand.Seed(time.Now().UnixNano())

	world := entity.World{}
	world.Init()

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

	// Advance the world.
	go func() {
		for {
			world.Step()
			termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
			world.Render()
			termbox.Flush()
			time.Sleep(time.Duration(1000000/20) * time.Microsecond)
		}
	}()

	// Block until exit is requested.
	<-quitChan
}
