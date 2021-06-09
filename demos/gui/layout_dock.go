package gui

import (
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.layout_dock"] = &GuiLayoutDock{}
}

type GuiLayoutDock struct{}

// Start is called once at the start of the demo.
func (t *GuiLayoutDock) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	dl := gui.NewDockLayout()
	a.DemoPanel().SetLayout(dl)

	// First top panel
	top1 := gui.NewPanel(0, 50)
	top1.SetBorders(1, 1, 1, 1)
	top1.SetPaddings(4, 4, 4, 4)
	top1.SetColor(math32.NewColor("green"))
	top1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockTop})
	a.DemoPanel().Add(top1)

	// Second top panel
	top2 := gui.NewPanel(0, 50)
	top2.SetBorders(1, 1, 1, 1)
	top2.SetPaddings(4, 4, 4, 4)
	top2.SetColor(math32.NewColor("blue"))
	top2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockTop})
	a.DemoPanel().Add(top2)

	// First bottom panel
	bot1 := gui.NewPanel(0, 32)
	bot1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockBottom})
	bot1.SetColor(math32.NewColor("red"))
	bot1.SetBorders(1, 1, 1, 1)
	a.DemoPanel().Add(bot1)

	// Second bottom panel
	bot2 := gui.NewPanel(0, 32)
	bot2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockBottom})
	bot2.SetColor(math32.NewColor("green"))
	bot2.SetBorders(1, 1, 1, 1)
	a.DemoPanel().Add(bot2)

	// First left panel
	left1 := gui.NewPanel(40, 0)
	left1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockLeft})
	left1.SetColor(math32.NewColor("black"))
	left1.SetBorders(1, 1, 1, 1)
	a.DemoPanel().Add(left1)

	// Second left panel
	left2 := gui.NewPanel(40, 0)
	left2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockLeft})
	left2.SetColor(math32.NewColor("red"))
	left2.SetBorders(1, 1, 1, 1)
	a.DemoPanel().Add(left2)

	// First right panel
	right1 := gui.NewPanel(40, 0)
	right1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockRight})
	right1.SetColor(math32.NewColor("black"))
	right1.SetBorders(1, 1, 1, 1)
	a.DemoPanel().Add(right1)

	// Second right panel
	right2 := gui.NewPanel(40, 0)
	right2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockRight})
	right2.SetColor(math32.NewColor("green"))
	right2.SetBorders(1, 1, 1, 1)
	a.DemoPanel().Add(right2)

	// Center panel
	center := gui.NewPanel(0, 0)
	center.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockCenter})
	a.DemoPanel().Add(center)
}

// Update is called every frame.
func (t *GuiLayoutDock) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiLayoutDock) Cleanup(a *app.App) {}
