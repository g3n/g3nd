package main

import (
	"fmt"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/graphic"
)

func init() {
	TestMap["gui.layout_vbox"] = &GuiLayoutVBox{}
}

type GuiLayoutVBox struct{}

func (t *GuiLayoutVBox) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	// Vertical panel
	params := gui.VBoxLayoutParams{Expand: 0, AlignH: gui.AlignLeft}
	p1 := gui.NewPanel(100, 500)
	p1.SetPosition(10, 10)
	p1.SetBorders(1, 1, 1, 1)
	p1.SetBordersColor(&math32.Black)
	p1.SetPaddings(2, 2, 2, 2)
	ctx.Gui.Add(p1)

	// Horizontal box layout
	bl1 := gui.NewVBoxLayout()
	bl1.SetSpacing(4)
	p1.SetLayout(bl1)
	// Add button
	p1b1 := gui.NewButton("Add")
	p1b1.SetPosition(p1.Position().X+p1.Width()+10, p1.Position().Y)
	p1b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		child := gui.NewButton(fmt.Sprintf("child %d", len(p1.Children())))
		itemParams := params
		child.SetLayoutParams(&itemParams)
		p1.Add(child)
	})
	ctx.Gui.Add(p1b1)

	// Del button
	p1b2 := gui.NewButton("Del")
	p1b2.SetPosition(p1b1.Position().X+p1b1.Width()+10, p1b1.Position().Y)
	p1b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if len(p1.Children()) > 0 {
			p1.Remove(p1.Children()[0].(gui.IPanel))
		}
	})
	ctx.Gui.Add(p1b2)
	// Top
	p1c1 := gui.NewRadioButton("Top")
	p1c1.SetPosition(p1b1.Position().X, p1b1.Position().Y+p1b1.Height()+10)
	p1c1.SetGroup("alignH")
	p1c1.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignTop) })
	ctx.Gui.Add(p1c1)
	// Center
	p1c2 := gui.NewRadioButton("Center")
	p1c2.SetPosition(p1c1.Position().X+p1c1.Width()+10, p1c1.Position().Y)
	p1c2.SetGroup("alignH")
	p1c2.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignCenter) })
	ctx.Gui.Add(p1c2)
	// Bottom
	p1c3 := gui.NewRadioButton("Bottom")
	p1c3.SetPosition(p1c2.Position().X+p1c2.Width()+10, p1c2.Position().Y)
	p1c3.SetGroup("alignH")
	p1c3.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignBottom) })
	ctx.Gui.Add(p1c3)
	// Height
	p1c4 := gui.NewRadioButton("Height")
	p1c4.SetPosition(p1c3.Position().X+p1c3.Width()+10, p1c3.Position().Y)
	p1c4.SetGroup("alignH")
	p1c4.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignHeight) })
	ctx.Gui.Add(p1c4)

	// Align next item to Left
	p1c5 := gui.NewRadioButton("Left")
	p1c5.SetPosition(p1c1.Position().X, p1c1.Position().Y+30)
	p1c5.SetGroup("alignV")
	p1c5.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignLeft })
	ctx.Gui.Add(p1c5)
	// Align next item to Center
	p1c6 := gui.NewRadioButton("Center")
	p1c6.SetPosition(p1c5.Position().X+p1c5.Width()+10, p1c5.Position().Y)
	p1c6.SetGroup("alignV")
	p1c6.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignCenter })
	ctx.Gui.Add(p1c6)
	// Align next item to Bottom
	p1c7 := gui.NewRadioButton("Right")
	p1c7.SetPosition(p1c6.Position().X+p1c6.Width()+10, p1c6.Position().Y)
	p1c7.SetGroup("alignV")
	p1c7.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignRight })
	ctx.Gui.Add(p1c7)
	// Align next item to Width
	p1c8 := gui.NewRadioButton("Width")
	p1c8.SetPosition(p1c7.Position().X+p1c7.Width()+10, p1c7.Position().Y)
	p1c8.SetGroup("alignV")
	p1c8.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignWidth })
	ctx.Gui.Add(p1c8)
	// Expand factor for next item
	p1c9 := gui.NewRadioButton("Expand")
	p1c9.SetPosition(p1c8.Position().X+p1c8.Width()+10, p1c8.Position().Y)
	p1c9.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if p1c9.Value() {
			params.Expand = 1
		} else {
			params.Expand = 0
		}
	})
	ctx.Gui.Add(p1c9)
}

func (t *GuiLayoutVBox) Render(ctx *Context) {
}

