package main

import (
	"fmt"
	"strconv"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["gui.layout_grid"] = &GuiLayoutGrid{}
}

type GuiLayoutGrid struct {
	bwidth  float32
	bheight float32
	layout  *gui.GridLayout
	colspan *int
}

func (t *GuiLayoutGrid) Initialize(ctx *Context) {

	var p *gui.Panel
	// Add button
	b1 := gui.NewButton("Add")
	b1.SetPosition(10, 10)
	b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		button := gui.NewButton(fmt.Sprintf("child %d", len(p.Children())))
		button.SetWidth(button.Width() + t.bwidth)
		button.SetHeight(button.Height() + t.bheight)
		var params gui.GridLayoutParams
		if t.colspan != nil {
			params.ColSpan = *t.colspan
			button.SetLayoutParams(&params)
			log.Info("LayoutParams: %v", params)
		}
		p.Add(button)
	})
	ctx.Gui.Add(b1)

	// Clear button
	b2 := gui.NewButton("Clear")
	b2.SetPosition(b1.Position().X+b1.Width()+10, 10)
	b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		p.DisposeChildren(true)
		t.colspan = nil
	})
	ctx.Gui.Add(b2)

	// Child button width slider
	s1 := gui.NewHSlider(42, 22)
	s1.SetPosition(b2.Position().X+b2.Width()+10, 10)
	s1.SetText("width")
	s1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.bwidth = s1.Value() * 20
	})
	ctx.Gui.Add(s1)

	// Slider to set the child width
	s2 := gui.NewHSlider(42, 22)
	s2.SetPosition(s1.Position().X+s1.Width()+10, 10)
	s2.SetText("height")
	s2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.bheight = s2.Value() * 20
	})
	ctx.Gui.Add(s2)

	// Edit field for child colspan
	e1 := gui.NewEdit(60, "Colspan")
	e1.SetPosition(s2.Position().X+s2.Width()+10, 10)
	e1.MaxLength = 2
	e1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		text := e1.Text()
		v, err := strconv.Atoi(text)
		if err == nil {
			if t.colspan == nil {
				t.colspan = new(int)
				*t.colspan = v
			}
		} else {
			t.colspan = nil
		}
	})
	ctx.Gui.Add(e1)

	// Grid Layout horizontal alignment
	dd1 := gui.NewDropDown(100, gui.NewImageLabel("horizontal"))
	dd1.SetPosition(10, b1.Position().Y+b1.Width()+6)
	dd1o1 := gui.NewImageLabel("None")
	dd1o1.SetUserData(gui.AlignNone)
	dd1.Add(dd1o1)
	dd1o2 := gui.NewImageLabel("Left")
	dd1o2.SetUserData(gui.AlignLeft)
	dd1.Add(dd1o2)
	dd1o3 := gui.NewImageLabel("Center")
	dd1o3.SetUserData(gui.AlignCenter)
	dd1.Add(dd1o3)
	dd1o4 := gui.NewImageLabel("Right")
	dd1o4.SetUserData(gui.AlignRight)
	dd1.Add(dd1o4)
	dd1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		sel := dd1.Selected()
		t.layout.SetAlignH(sel.UserData().(gui.Align))
	})
	ctx.Gui.Add(dd1)

	// Grid Layout vertical alignment
	dd2 := gui.NewDropDown(100, gui.NewImageLabel("vertical"))
	dd2.SetPosition(dd1.Position().X+dd1.Width()+10, dd1.Position().Y)
	dd2o1 := gui.NewImageLabel("None")
	dd2o1.SetUserData(gui.AlignNone)
	dd2.Add(dd2o1)
	dd2o2 := gui.NewImageLabel("Top")
	dd2o2.SetUserData(gui.AlignTop)
	dd2.Add(dd2o2)
	dd2o3 := gui.NewImageLabel("Center")
	dd2o3.SetUserData(gui.AlignCenter)
	dd2.Add(dd2o3)
	dd2o4 := gui.NewImageLabel("Bottom")
	dd2o4.SetUserData(gui.AlignBottom)
	dd2.Add(dd2o4)
	dd2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		sel := dd2.Selected()
		t.layout.SetAlignV(sel.UserData().(gui.Align))
	})
	ctx.Gui.Add(dd2)

	// Creates panel with grid layout
	p = gui.NewPanel(600, 400)
	p.SetPosition(10, dd1.Position().Y+dd1.Height()+10)
	p.SetColor(&math32.White)
	p.SetBorders(1, 1, 1, 1)
	t.layout = gui.NewGridLayout(6)
	p.SetLayout(t.layout)
	ctx.Gui.Add(p)
}

func (t *GuiLayoutGrid) Render(ctx *Context) {
}
