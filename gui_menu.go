package main

import (
	"github.com/g3n/engine/gui"
)

func init() {
	TestMap["gui.menu"] = &GuiMenu{}
}

type GuiMenu struct {
}

func (t *GuiMenu) Initialize(ctx *Context) {

	m1 := gui.NewMenu()
	m1.SetPosition(10, 10)

	ctx.Gui.Add(m1)
}

func (t *GuiMenu) Render(ctx *Context) {
}
