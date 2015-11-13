package entity

import (
	"github.com/nsf/termbox-go"
)

type Food struct {
	Position Point
	Eaten    bool
}

func (f *Food) Kill() {
	f.Eaten = true
}

func (f *Food) At(point Point) bool {
	return f.Position.At(point)
}

func (f *Food) Render() {
	if !f.Eaten {
		termbox.SetCell(f.Position.X, f.Position.Y, '$', termbox.ColorGreen, termbox.ColorDefault)
	}
}
