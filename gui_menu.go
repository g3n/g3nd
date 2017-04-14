package main

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets"
	"github.com/g3n/engine/window"
)

func init() {
	TestMap["gui.menu"] = &GuiMenu{}
}

type GuiMenu struct {
}

func (t *GuiMenu) Initialize(ctx *Context) {

	m1 := gui.NewMenu()
	m1.SetPosition(10, 10)

	m1.AddOption("Menu1 Option1")
	m1.AddOption("Menu1 Option2 (longer)")
	m1i1 := m1.AddOption("Menu1 Option3")
	m1i1.SetEnabled(false)
	m1.AddSeparator()
	m1.AddOption("Menu1 Option4")
	ctx.Gui.Add(m1)

	m2 := gui.NewMenu()
	m2.SetPosition(m1.Position().X+m1.Width()+10, 10)
	m2i1 := m2.AddOption("Menu2 Option1")
	m2i1.SetIcon(assets.Add)
	m2i2 := m2.AddOption("Menu2 Option2 (longer)")
	m2i2.SetIcon(assets.ArrowBack)
	m2i2.SetShortcut(window.ModShift, window.KeyA)
	m2.AddOption("Menu2 Option3")
	m2.AddSeparator()
	m2i4 := m2.AddOption("Menu2 Option4")
	m2i4.SetShortcut(window.ModControl, window.KeyB)
	ctx.Gui.Add(m2)

	m3 := gui.NewMenu()
	m3.SetPosition(m1.Position().X, m1.Position().Y+m1.Height()+10)
	m3.AddOption("Menu3 Option1")
	m3.AddOption("Menu3 Option2")

	m3sm1 := gui.NewMenu()
	m3sm1.AddOption("Menu3/SubMenu1/Option1")
	m3sm1.AddOption("Menu3/SubMenu1/Option2")
	m3sm1.AddSeparator()
	m3sm1sm2 := gui.NewMenu()
	m3sm1sm2.AddOption("Menu3/SubMenu1/SubMenu2/Option1")
	m3sm1sm2.AddOption("Menu3/SubMenu1/SubMenu2/Option2")
	m3sm1.AddMenu("Menu3/SubMenu1/SubMenu2", m3sm1sm2)

	m3.AddSeparator()
	m3.AddMenu("Menu3/SubMenu1", m3sm1)
	m3.AddOption("Menu3/Option3")

	ctx.Gui.Add(m3)
}

func (t *GuiMenu) Render(ctx *Context) {
}
