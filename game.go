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

// Game is the admin of everything.
type Game struct {
	eventChan chan *termbox.Event
	quitChan  chan bool
	world     entity.World
}

func (g *Game) init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	rand.Seed(time.Now().UnixNano())

	g.quitChan = make(chan bool)
	g.eventChan = make(chan *termbox.Event)

	g.world = entity.World{}
	g.world.Init()
}

func (g *Game) destroy() {
	termbox.Close()
}

// Run creates and runs the game. Runs until user quits.
func (g *Game) Run() {
	g.init()
	defer g.destroy()

	go g.pollForEvents()
	go g.handleEvents()
	go g.runGameLoop()

	g.waitForQuit()

}

func (g *Game) pollForEvents() {
	for {
		event := termbox.PollEvent()
		g.eventChan <- &event
	}
}

func (g *Game) handleEvents() {
	for {
		select {
		case event := <-g.eventChan:
			switch event.Type {
			case termbox.EventKey:
				switch event.Key {
				case termbox.KeyEsc:
					g.quitChan <- true
				default:
					g.world.ProcessEvent(event)
				}
			}
		}
	}
}

func (g *Game) runGameLoop() {
	stepRate := time.Duration(1000/20) * time.Millisecond
	for {
		g.world.Step()
		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
		g.world.Render()
		termbox.Flush()
		time.Sleep(stepRate)
	}
}

func (g *Game) waitForQuit() {
	<-g.quitChan
}
