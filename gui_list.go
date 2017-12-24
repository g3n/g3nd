package main

import (
	"fmt"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
)

func init() {
	TestMap["gui.list"] = &GuiList{}
}

type GuiList struct{}

func (t *GuiList) Initialize(ctx *Context) {

	//axis := graphic.NewAxisHelper(1)
	//ctx.Scene.Add(axis)

	icons := []string{
		icon.ArrowBack,
		icon.ArrowDownward,
		icon.ArrowDropDown,
		icon.ArrowDropDownCircle,
		icon.ArrowDropUp,
		icon.ArrowForward,
		icon.ArrowUpward,
	}

	// List 1 vertical/single selection
	li1 := gui.NewVList(100, 200)
	li1.SetPosition(10, 10)
	ctx.Gui.Add(li1)
	// List 1 - add button
	b1 := gui.NewButton("Add")
	b1.SetPosition(li1.Position().X+li1.Width()+10, li1.Position().Y)
	b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		pos := li1.Len()
		item := gui.NewImageLabel(fmt.Sprintf("label %d", pos))
		item.SetIcon(string(icons[pos%len(icons)]))
		li1.Add(item)
	})
	ctx.Gui.Add(b1)
	// List 1 - remove button
	b2 := gui.NewButton("Del")
	b2.SetPosition(li1.Position().X+li1.Width()+10, li1.Position().Y+30)
	b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if li1.Len() > 0 {
			li1.RemoveAt(0)
		}
	})
	ctx.Gui.Add(b2)
	// List 1 - clear button
	b3 := gui.NewButton("Clear")
	b3.SetPosition(li1.Position().X+li1.Width()+10, li1.Position().Y+60)
	b3.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		li1.Clear()
	})
	ctx.Gui.Add(b3)

	// List 2 vertical/multi selection
	li2 := gui.NewVList(100, 200)
	li2.SetSingle(false)
	li2.SetPosition(b1.Position().X+b1.Width()+50, 10)
	ctx.Gui.Add(li2)
	// List 2 - add button
	b4 := gui.NewButton("Add")
	b4.SetPosition(li2.Position().X+li2.Width()+10, li2.Position().Y)
	b4.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		item := gui.NewImageLabel(fmt.Sprintf("label %d", li2.Len()))
		li2.Add(item)
	})
	ctx.Gui.Add(b4)
	// List 2 - remove button
	b5 := gui.NewButton("Del")
	b5.SetPosition(li2.Position().X+li2.Width()+10, li2.Position().Y+30)
	b5.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if li2.Len() > 0 {
			li2.RemoveAt(0)
		}
	})
	ctx.Gui.Add(b5)
	// List 2 - clear button
	b6 := gui.NewButton("Clear")
	b6.SetPosition(li2.Position().X+li2.Width()+10, li2.Position().Y+60)
	b6.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		li2.Clear()
	})
	ctx.Gui.Add(b6)

	// List 3 horizontal/single selection
	li3 := gui.NewHList(400, 64)
	li3.SetPosition(10, 250)
	ctx.Gui.Add(li3)
	// List 3 - add button
	b7 := gui.NewButton("Add")
	b7.SetPosition(li3.Position().X, li3.Position().Y+li3.Height()+10)
	b7.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		pos := li3.Len()
		item := gui.NewImageLabel(fmt.Sprintf("label %d", pos))
		item.SetIcon(string(icons[pos%len(icons)]))
		li3.Add(item)
	})
	ctx.Gui.Add(b7)
	// List 3 - remove button
	b8 := gui.NewButton("Del")
	b8.SetPosition(b7.Position().X+b7.Width()+40, li3.Position().Y+li3.Height()+10)
	b8.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if li3.Len() > 0 {
			li3.RemoveAt(0)
		}
	})
	ctx.Gui.Add(b8)
	// List 3 - clear button
	b9 := gui.NewButton("Clear")
	b9.SetPosition(b8.Position().X+b8.Width()+40, li3.Position().Y+li3.Height()+10)
	b9.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		li3.Clear()
	})
	ctx.Gui.Add(b9)
}

func (t *GuiList) Render(ctx *Context) {
}
