package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["gui.window"] = &GuiWindow{}
}

type GuiWindow struct{}

func (t *GuiWindow) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	w1 := gui.NewWindow(200, 100)
	w1.SetPosition(10, 10)
	w1.SetResizable(gui.ResizeAll)
	w1.SetLayout(gui.NewFillLayout(true, true))
	c1 := gui.NewImageLabel(
		"This is the client area of the Window 1\n" +
			"This window does not have a title\n" +
			"It is resizable on all borders\n",
	)
	c1.SetBgColor(math32.NewColor("white"))
	w1.Add(c1)
	ctx.Gui.Add(w1)

	w2 := gui.NewWindow(200, 100)
	w2.SetTitle("Window2")
	w2.SetPosition(w1.Position().X+w1.Width()+50, 10)
	w2.SetResizable(gui.ResizeRight | gui.ResizeBottom)
	w2.SetLayout(gui.NewFillLayout(true, true))
	c2 := gui.NewImageLabel(
		"This is the client area of the Window 2\n" +
			"This window has a title\n" +
			"It is resizable on the bottom and right borders\n" +
			"It is movable by dragging its title",
	)
	c2.SetBgColor(&math32.Color{0.7, 0.8, 0.9})
	w2.Add(c2)
	ctx.Gui.Add(w2)
}

func (t *GuiWindow) Render(ctx *Context) {
}
