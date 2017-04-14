package main

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets"
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

	m2 := gui.NewMenu()
	m2.SetPosition(m1.Position().X+m1.Width()+10, 10)

	m2i1 := m2.AddItem("Menu2 Option1")
	m2i1.SetIcon(assets.Add)

	m2i2 := m2.AddItem("Menu2 Option2 (longer)")
	m2i2.SetIcon(assets.ArrowBack)

	m2.AddItem("Menu2 Option3")
	m2.AddSeparator()
	m2.AddItem("Menu2 Option4")
	ctx.Gui.Add(m2)
}

func (t *GuiMenu) Render(ctx *Context) {
}
