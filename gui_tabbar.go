package main

import (
	"fmt"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["gui.tabbar"] = &GuiTabBar{}
}

type GuiTabBar struct {
	tb *gui.TabBar
}

func (t *GuiTabBar) Initialize(ctx *Context) {

	// Button for adding tabs
	tabCounter := 1
	colors := []string{
		"LightSteelBlue", "PowderBlue", "LightBlue", "SkyBlue", "LightSkyBlue", "DeepSkyBlue",
	}
	b1 := gui.NewButton("Add Tab")
	b1.SetPosition(10, 10)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		cname := colors[tabCounter%len(colors)]
		tabText := fmt.Sprintf("Tab: %d (%s)", tabCounter, cname)
		tab := t.tb.AddTab(tabText)
		tabCounter++
		tab.Content().SetLayout(gui.NewFillLayout(true, true))
		label := gui.NewImageLabel(tabText)
		label.SetFontSize(20)
		tab.Content().Add(label)
		tab.Content().SetColor(math32.NewColor(cname))
	})
	ctx.Gui.Add(b1)

	// Button for removing all tabs
	b2 := gui.NewButton("Clear All")
	b2.SetPosition(b1.Position().X+b1.Width()+10, b1.Position().Y)
	b2.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		for t.tb.TabCount() > 0 {
			t.tb.RemoveTab(0)
		}
		tabCounter = 1
	})
	ctx.Gui.Add(b2)

	// Creates TabBar
	t.tb = gui.NewTabBar(0, 0)
	tby := b1.Position().Y + b1.Height() + 10
	t.tb.SetPosition(b1.Position().X, tby)
	ctx.Gui.Add(t.tb)

	// Resizes TabBar when main window resizes
	ctx.Gui.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		t.tb.SetSize(ctx.Gui.ContentWidth()-t.tb.Position().X-10, ctx.Gui.ContentHeight()-tby-10)
	})
	ctx.Gui.Dispatch(gui.OnResize, nil)
}

func (t *GuiTabBar) Render(ctx *Context) {
}
