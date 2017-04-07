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

	m1.AddItem("Menu1 Option1")
	m1.AddItem("Menu1 Option2 (longer)")
	m1.AddItem("Menu1 Option3")
	m1.AddSeparator()
	m1.AddItem("Menu1 Option4")

	ctx.Gui.Add(m1)
}

func (t *GuiMenu) Render(ctx *Context) {
}
