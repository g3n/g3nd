package main

import (
	"fmt"

	"github.com/g3n/engine/gui"
)

func init() {
	TestMap["gui.scrollbar"] = &GuiScrollBar{}
}

type GuiScrollBar struct{}

func (t *GuiScrollBar) Initialize(ctx *Context) {

	// Scroll bar 1
	sb1 := gui.NewHScrollBar(100, 16)
	sb1.SetPosition(10, 10)
	ctx.Gui.Add(sb1)
	// Position
	l1 := gui.NewLabel("Pos:")
	l1.SetPosition(sb1.Position().X+sb1.Width()+10, sb1.Position().Y)
	ctx.Gui.Add(l1)
	sb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l1.SetText(fmt.Sprintf("Pos:%1.2f", sb1.Value()))
	})

	// Scroll bar 2
	sb2 := gui.NewHScrollBar(300, 64)
	sb2.SetPosition(10, 40)
	ctx.Gui.Add(sb2)
	// Position
	l2 := gui.NewLabel("Pos:")
	l2.SetPosition(sb2.Position().X+sb2.Width()+10, sb2.Position().Y)
	ctx.Gui.Add(l2)
	sb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l2.SetText(fmt.Sprintf("Pos:%1.2f", sb2.Value()))
	})

	// Scroll bar 3
	sb3 := gui.NewVScrollBar(16, 100)
	sb3.SetPosition(10, 120)
	ctx.Gui.Add(sb3)
	// Position
	l3 := gui.NewLabel("Pos:")
	l3.SetPosition(sb3.Position().X+sb3.Width()+10, sb3.Position().Y)
	ctx.Gui.Add(l3)
	sb3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l3.SetText(fmt.Sprintf("Pos:%1.2f", sb3.Value()))
	})

	// Scroll bar 4
	sb4 := gui.NewVScrollBar(64, 300)
	sb4.SetPosition(10, 230)
	ctx.Gui.Add(sb4)
	// Position
	l4 := gui.NewLabel("Pos:")
	l4.SetPosition(sb4.Position().X+sb4.Width()+10, sb4.Position().Y)
	ctx.Gui.Add(l4)
	sb4.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l4.SetText(fmt.Sprintf("Pos:%1.2f", sb4.Value()))
	})
}

func (t *GuiScrollBar) Render(ctx *Context) {
}
