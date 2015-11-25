package framework

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/khueue/go-snake/config"
	"github.com/khueue/go-snake/entity"
	"github.com/nsf/termbox-go"
)

// Game is the admin of everything.
type Game struct {
	eventChan chan termbox.Event
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
	g.eventChan = make(chan termbox.Event)

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

	fmt.Println("Welcome!")

	go g.pollForEvents()
	go g.handleEvents()
	go g.runGameLoop()

	g.waitForQuit()
}

func (g *Game) pollForEvents() {
	for {
		event := termbox.PollEvent()
		g.eventChan <- event
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

func nowSeconds() float64 {
	return float64(time.Now().UnixNano()) / (1000 * 1000 * 1000)
}

func (g *Game) runGameLoop() {
	for {
		currentTime := nowSeconds()

		g.world.Step()

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
		g.world.Render()
		termbox.Flush()

		// Cap FPS.
		frameTime := (nowSeconds() - currentTime) * 1000000
		// fmt.Println("stepping took microseconds", frameTime)
		sleepTime := time.Duration(1000000.0/float64(config.TargetFPS)-frameTime) * time.Microsecond
		// fmt.Println("sleeping for", sleepTime)
		time.Sleep(sleepTime)
	}
}

func (g *Game) waitForQuit() {
	<-g.quitChan
	fmt.Println("Bye! ESC pressed.")
}
