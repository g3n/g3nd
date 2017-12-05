package main

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
	"strings"
)

func init() {
	TestMap["gui.menu"] = &GuiMenu{}
}

type GuiMenu struct {
}

func (t *GuiMenu) Initialize(ctx *Context) {

	// Label
	mbText := "Selected: "
	mbOption := gui.NewLabel(mbText)
	mbOption.SetPosition(300, 10)
	mbOption.SetPaddings(2, 2, 2, 2)
	mbOption.SetBgColor(&math32.White)
	mbOption.SetBorders(1, 1, 1, 1)
	ctx.Gui.Add(mbOption)

	// Event handler for menu clicks
	onClick := func(evname string, ev interface{}) {
		path := strings.Join(ev.(*gui.MenuItem).IdPath(), "/")
		mbOption.SetText(mbText + path)
	}

	// Create menu bar
	mb := gui.NewMenuBar()
	mb.Subscribe(gui.OnClick, onClick)
	mb.SetPosition(10, 10)

	// Create Menu1 and adds it to the menu bar
	m1 := gui.NewMenu()
	m1.AddOption("Menu1/Option1").
		SetId("option1")
	m1.AddOption("Menu1/Option2").
		SetId("option2")
	m1.AddOption("Menu1/Option3").
		SetId("option3").
		SetEnabled(false)
	m1.AddSeparator()
	m1.AddOption("Menu1/Option4").
		SetId("option4")
	mb.AddMenu("Menu1", m1).
		SetId("menu1").
		SetShortcut(window.ModAlt, window.Key1)

	// Create Menu2 and adds it to the menu bar
	m2 := gui.NewMenu()
	m2.AddOption("Menu2/Option1").
		SetId("option1").
		SetIcon(icon.Build).
		SetShortcut(window.ModControl, window.KeyA)
	m2.AddOption("Menu2/Option two").
		SetId("option2").
		SetIcon(icon.Cached).
		SetShortcut(window.ModShift, window.KeyB)
	m2.AddSeparator()
	m2.AddOption("Menu2/Option three").
		SetId("option3").
		SetIcon(icon.Print).
		SetShortcut(window.ModAlt, window.KeyC)
	m2.AddOption("Menu2/Option four").
		SetId("option4").
		SetIcon(icon.Settings).
		SetShortcut(window.ModAlt|window.ModShift, window.KeyD)
	m2.AddOption("Menu2/Option five").
		SetId("option5").
		SetIcon(icon.Search).
		SetShortcut(window.ModAlt|window.ModShift|window.ModControl, window.KeyE)
	mb.AddMenu("Menu2", m2).
		SetId("menu2").
		SetShortcut(window.ModAlt, window.Key2)

	// Create Menu3 and adds it to the menu bar
	m3 := gui.NewMenu()
	m3.AddOption("Menu3 Option1").
		SetId("option1").
		SetIcon(icon.Star).
		SetShortcut(0, window.KeyF1)
	m3.AddOption("Menu3 Option2").
		SetId("option2").
		SetIcon(icon.StarBorder).
		SetShortcut(window.ModControl, window.KeyF2)
	// Creates Menu3/Menu1
	m3m1 := gui.NewMenu()
	m3m1.AddOption("Menu3/Menu1/Option1").
		SetId("option1").
		SetIcon(icon.StarHalf).
		SetShortcut(window.ModAlt, window.KeyF3)
	m3m1.AddOption("Menu3/Menu1/Option2").
		SetId("option2").
		SetIcon(icon.Opacity).
		SetShortcut(window.ModAlt|window.ModControl, window.KeyF4)
	m3m1.AddSeparator()
	// Creates Menu3/Menu1/Menu2
	m3m1m2 := gui.NewMenu()
	m3m1m2.AddOption("Menu3/Menu1/Menu2/Option1").
		SetId("option1").
		SetIcon(icon.HourglassFull).
		SetShortcut(window.ModAlt|window.ModControl|window.ModShift, window.KeyF5)
	m3m1m2.AddOption("Menu3/Menu1/Menu2/Option2").
		SetId("option2").
		SetIcon(icon.HourglassEmpty).
		SetShortcut(0, window.KeyF6)
	m3m1.AddMenu("Menu3/Menu1/Menu2", m3m1m2).
		SetId("menu2")
	m3.AddSeparator()
	m3.AddMenu("Menu3/Menu1", m3m1).
		SetId("menu1").
		SetIcon(icon.Home)
	m3.AddOption("Menu3/Option3").
		SetId("option3")
	mb.AddMenu("Menu3", m3).
		SetId("menu3").
		SetShortcut(window.ModAlt, window.Key3)

	// Add separators and options to the menu bar
	mb.AddSeparator()
	mb.AddOption("OptionA").
		SetId("optionA").
		SetShortcut(window.ModAlt, window.KeyA)
	mb.AddOption("OptionB").
		SetId("optionB").
		SetShortcut(window.ModAlt, window.KeyB)

	ctx.Gui.Add(mb)
	ctx.Gui.Root().SetKeyFocus(mb)

}

func (t *GuiMenu) Render(ctx *Context) {
}
