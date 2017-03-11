package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["gui.layout_dock"] = &GuiLayoutDock{}
}

type GuiLayoutDock struct{}

func (t *GuiLayoutDock) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	dl := gui.NewDockLayout()
	ctx.Gui.SetLayout(dl)

	// First top panel
	top1 := gui.NewPanel(0, 50)
	top1.SetBorders(1, 1, 1, 1)
	top1.SetPaddings(4, 4, 4, 4)
	top1.SetColor(&math32.Green)
	top1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockTop})
	ctx.Gui.Add(top1)

	// Second top panel
	top2 := gui.NewPanel(0, 50)
	top2.SetBorders(1, 1, 1, 1)
	top2.SetPaddings(4, 4, 4, 4)
	top2.SetColor(&math32.Blue)
	top2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockTop})
	ctx.Gui.Add(top2)

	// First bottom panel
	bot1 := gui.NewPanel(0, 32)
	bot1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockBottom})
	bot1.SetColor(&math32.Red)
	bot1.SetBorders(1, 1, 1, 1)
	ctx.Gui.Add(bot1)

	// Second bottom panel
	bot2 := gui.NewPanel(0, 32)
	bot2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockBottom})
	bot2.SetColor(&math32.Green)
	bot2.SetBorders(1, 1, 1, 1)
	ctx.Gui.Add(bot2)

	// First left panel
	left1 := gui.NewPanel(40, 0)
	left1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockLeft})
	left1.SetColor(&math32.Black)
	left1.SetBorders(1, 1, 1, 1)
	ctx.Gui.Add(left1)

	// Second left panel
	left2 := gui.NewPanel(40, 0)
	left2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockLeft})
	left2.SetColor(&math32.Red)
	left2.SetBorders(1, 1, 1, 1)
	ctx.Gui.Add(left2)

	// First right panel
	right1 := gui.NewPanel(40, 0)
	right1.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockRight})
	right1.SetColor(&math32.Black)
	right1.SetBorders(1, 1, 1, 1)
	ctx.Gui.Add(right1)

	// Second right panel
	right2 := gui.NewPanel(40, 0)
	right2.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockRight})
	right2.SetColor(&math32.Green)
	right2.SetBorders(1, 1, 1, 1)
	ctx.Gui.Add(right2)

	// Center panel
	center := gui.NewPanel(0, 0)
	center.SetLayoutParams(&gui.DockLayoutParams{Edge: gui.DockCenter})
	ctx.Gui.Add(center)
}

func (t *GuiLayoutDock) Render(ctx *Context) {
}
