package main

import (
	"fmt"
	"math/rand"
	"path/filepath"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
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
	icons := []string{
		"Build", "Cached", "Done", "ChangeHistory", "FavoriteBorder", "Check",
	}
	b1 := gui.NewButton("Add Tab")
	b1.SetPosition(10, 10)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		// Adds a new Tab
		cname := colors[tabCounter%len(colors)]
		tabText := fmt.Sprintf("Tab: %d (%s)", tabCounter, cname)
		tab := t.tb.AddTab(tabText)
		tab.SetIcon(icon.Codepoint(icons[tabCounter%len(icons)]))
		tabCounter++
		// Sets the Tab content panel
		label := gui.NewImageLabel(tabText)
		label.SetFontSize(20)
		label.SetBgColor(math32.NewColor(cname))
		tab.SetContent(label)
		// Sets Tab menu
		t.setTabMenu(ctx, tab)

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

func (t *GuiTabBar) setTabMenu(ctx *Context, tab *gui.Tab) {

	// Creates menu and adds options
	menu := gui.NewMenu()
	menu.SetBounded(false)
	menu.SetVisible(false)
	const (
		closeTab       = "Close Tab"
		closeOthers    = "Close other tabs"
		closeTabsRight = "Close tabs to the right"
		setImage       = "Set image"
		pin            = "Pin tab"
		unpin          = "Unpin tab"
		moveLeft       = "Move tab left"
		moveRight      = "Move tab right"
	)
	options := []string{
		closeTab, closeOthers, closeTabsRight, "", setImage, "", pin, unpin, "",
		moveLeft, moveRight,
	}
	for _, op := range options {
		if op != "" {
			menu.AddOption(op).SetId(op)
		} else {
			menu.AddSeparator()
		}
	}

	images := []string{
		"add.png", "call-start.png", "document-open.png", "document-revert.png",
	}

	// Subscribe to menu click
	menu.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		menu.SetVisible(false)
		opid := ev.(*gui.MenuItem).Id()
		switch opid {
		case closeTab:
			t.tb.RemoveTab(t.tb.TabPosition(tab))
		case closeOthers:
			for t.tb.TabCount() > 1 {
				if tab != t.tb.TabAt(0) {
					t.tb.RemoveTab(0)
					continue
				}
				if tab != t.tb.TabAt(1) {
					t.tb.RemoveTab(1)
					continue
				}
			}
		case closeTabsRight:
			pos := t.tb.TabPosition(tab)
			for t.tb.TabCount() > pos+1 {
				t.tb.RemoveTab(pos + 1)
			}
		case setImage:
			idx := rand.Int31n(int32(len(images)))
			tab.SetImage(filepath.Join(ctx.DirData, "/images", images[idx]))
		case pin:
			tab.SetPinned(true)
		case unpin:
			tab.SetPinned(false)
		case moveLeft:
			pos := t.tb.TabPosition(tab)
			if pos == 0 {
				return
			}
			t.tb.MoveTab(pos, pos-1)
		case moveRight:
			pos := t.tb.TabPosition(tab)
			if pos >= t.tb.TabCount()-1 {
				return
			}
			t.tb.MoveTab(pos, pos+1)
		}
	})
	menu.Subscribe(gui.OnMouseOut, func(evname string, ev interface{}) {
		menu.SetVisible(false)
	})

	// Adds menu to the Tab header and sets it visible when
	// the Tab header is right clicked
	header := tab.Header()
	header.Add(menu)
	header.Subscribe(gui.OnRightClick, func(evname string, ev interface{}) {
		mev := ev.(*window.MouseEvent)
		px, py := header.ContentCoords(mev.Xpos, mev.Ypos)
		menu.SetPosition(px, py)
		menu.SetVisible(true)
	})
}

func (t *GuiTabBar) Render(ctx *Context) {
}
