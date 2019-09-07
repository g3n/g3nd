package gui

import (
	"fmt"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.scrollbar"] = &GuiScrollBar{}
}

type GuiScrollBar struct{}

// Start is called once at the start of the demo.
func (t *GuiScrollBar) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// Scroll bar 1
	sb1 := gui.NewHScrollBar(100, 16)
	sb1.SetPosition(10, 10)
	a.DemoPanel().Add(sb1)
	// Position
	l1 := gui.NewLabel("Pos:")
	l1.SetPosition(sb1.Position().X+sb1.Width()+10, sb1.Position().Y)
	a.DemoPanel().Add(l1)
	sb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l1.SetText(fmt.Sprintf("Pos:%1.2f", sb1.Value()))
	})

	// Scroll bar 2
	sb2 := gui.NewHScrollBar(300, 64)
	sb2.SetPosition(10, 40)
	a.DemoPanel().Add(sb2)
	// Position
	l2 := gui.NewLabel("Pos:")
	l2.SetPosition(sb2.Position().X+sb2.Width()+10, sb2.Position().Y)
	a.DemoPanel().Add(l2)
	sb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l2.SetText(fmt.Sprintf("Pos:%1.2f", sb2.Value()))
	})

	// Scroll bar 3
	sb3 := gui.NewVScrollBar(16, 100)
	sb3.SetPosition(10, 120)
	a.DemoPanel().Add(sb3)
	// Position
	l3 := gui.NewLabel("Pos:")
	l3.SetPosition(sb3.Position().X+sb3.Width()+10, sb3.Position().Y)
	a.DemoPanel().Add(l3)
	sb3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l3.SetText(fmt.Sprintf("Pos:%1.2f", sb3.Value()))
	})

	// Scroll bar 4
	sb4 := gui.NewVScrollBar(64, 300)
	sb4.SetPosition(10, 230)
	a.DemoPanel().Add(sb4)
	// Position
	l4 := gui.NewLabel("Pos:")
	l4.SetPosition(sb4.Position().X+sb4.Width()+10, sb4.Position().Y)
	a.DemoPanel().Add(l4)
	sb4.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		l4.SetText(fmt.Sprintf("Pos:%1.2f", sb4.Value()))
	})
}

// Update is called every frame.
func (t *GuiScrollBar) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiScrollBar) Cleanup(a *app.App) {}
