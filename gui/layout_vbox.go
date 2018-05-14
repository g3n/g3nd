package gui

import (
	"fmt"
	"math/rand"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.layout_vbox"] = &GuiLayoutVBox{}
}

type GuiLayoutVBox struct{}

func (t *GuiLayoutVBox) Initialize(a *app.App) {

	// Adds reset size button
	var p1 *gui.Panel
	const pwidth = 32
	const pheight = 32
	rb := gui.NewButton("Reset size")
	rb.SetPosition(10, 10)
	rb.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		p1.SetSize(pwidth, pheight)
	})
	a.GuiPanel().Add(rb)

	// Vertical panel
	params := gui.VBoxLayoutParams{Expand: 0, AlignH: gui.AlignLeft}
	p1 = gui.NewPanel(pwidth, pheight)
	p1.SetPosition(rb.Position().X, rb.Position().Y+rb.Height()+4)
	p1.SetBorders(1, 1, 1, 1)
	p1.SetBordersColor(math32.NewColor("black"))
	p1.SetPaddings(2, 2, 2, 2)
	a.GuiPanel().Add(p1)
	// Horizontal box layout
	bl1 := gui.NewVBoxLayout()
	bl1.SetSpacing(4)
	p1.SetLayout(bl1)

	const bposx = 140
	// Add button
	p1b1 := gui.NewButton("Add")
	p1b1.SetPosition(bposx, p1.Position().Y)
	p1b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		child := gui.NewButton(fmt.Sprintf("child %d", len(p1.Children())))
		offs := rand.Int31n(30)
		child.SetWidth(child.Width() + float32(offs))
		itemParams := params
		child.SetLayoutParams(&itemParams)
		p1.Add(child)
	})
	a.GuiPanel().Add(p1b1)

	// Del button
	p1b2 := gui.NewButton("Del")
	p1b2.SetPosition(p1b1.Position().X+p1b1.Width()+10, p1b1.Position().Y)
	p1b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if len(p1.Children()) > 0 {
			p1.Remove(p1.Children()[0].(gui.IPanel))
		}
	})
	a.GuiPanel().Add(p1b2)
	// Auto height checkbox
	cb1 := gui.NewCheckBox("Auto height")
	cb1.SetPosition(p1b2.Position().X+p1b2.Width()+10, p1b2.Position().Y)
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		bl1.SetAutoHeight(cb1.Value())
	})
	cb1.SetValue(true)
	a.GuiPanel().Add(cb1)
	// Auto width checkbox
	cb2 := gui.NewCheckBox("Auto width")
	cb2.SetPosition(cb1.Position().X+cb1.Width()+10, cb1.Position().Y)
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) { bl1.SetAutoWidth(cb2.Value()) })
	cb2.SetValue(true)
	a.GuiPanel().Add(cb2)
	// Top
	p1c1 := gui.NewRadioButton("Top")
	p1c1.SetPosition(cb1.Position().X, cb1.Position().Y+cb1.Height()+10)
	p1c1.SetGroup("alignH")
	p1c1.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignTop) })
	a.GuiPanel().Add(p1c1)
	// Center
	p1c2 := gui.NewRadioButton("Center")
	p1c2.SetPosition(p1c1.Position().X+p1c1.Width()+10, p1c1.Position().Y)
	p1c2.SetGroup("alignH")
	p1c2.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignCenter) })
	a.GuiPanel().Add(p1c2)
	// Bottom
	p1c3 := gui.NewRadioButton("Bottom")
	p1c3.SetPosition(p1c2.Position().X+p1c2.Width()+10, p1c2.Position().Y)
	p1c3.SetGroup("alignH")
	p1c3.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignBottom) })
	a.GuiPanel().Add(p1c3)
	// Height
	p1c4 := gui.NewRadioButton("Height")
	p1c4.SetPosition(p1c3.Position().X+p1c3.Width()+10, p1c3.Position().Y)
	p1c4.SetGroup("alignH")
	p1c4.Subscribe(gui.OnClick, func(evname string, ev interface{}) { bl1.SetAlignV(gui.AlignHeight) })
	a.GuiPanel().Add(p1c4)

	// Align next item to Left
	p1c5 := gui.NewRadioButton("Left")
	p1c5.SetPosition(p1c1.Position().X, p1c1.Position().Y+30)
	p1c5.SetGroup("alignV")
	p1c5.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignLeft })
	a.GuiPanel().Add(p1c5)
	// Align next item to Center
	p1c6 := gui.NewRadioButton("Center")
	p1c6.SetPosition(p1c5.Position().X+p1c5.Width()+10, p1c5.Position().Y)
	p1c6.SetGroup("alignV")
	p1c6.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignCenter })
	a.GuiPanel().Add(p1c6)
	// Align next item to Bottom
	p1c7 := gui.NewRadioButton("Right")
	p1c7.SetPosition(p1c6.Position().X+p1c6.Width()+10, p1c6.Position().Y)
	p1c7.SetGroup("alignV")
	p1c7.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignRight })
	a.GuiPanel().Add(p1c7)
	// Align next item to Width
	p1c8 := gui.NewRadioButton("Width")
	p1c8.SetPosition(p1c7.Position().X+p1c7.Width()+10, p1c7.Position().Y)
	p1c8.SetGroup("alignV")
	p1c8.Subscribe(gui.OnClick, func(evname string, ev interface{}) { params.AlignH = gui.AlignWidth })
	a.GuiPanel().Add(p1c8)
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
	a.GuiPanel().Add(p1c9)
}

func (t *GuiLayoutVBox) Render(a *app.App) {
}
