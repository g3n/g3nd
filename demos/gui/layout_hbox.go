package gui

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.layout_hbox"] = &GuiLayoutHBox{}
}

type GuiLayoutHBox struct{}

// Start is called once at the start of the demo.
func (t *GuiLayoutHBox) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// Adds reset size button
	var p1 *gui.Panel
	const pwidth = 32
	const pheight = 32
	rb := gui.NewButton("Reset size")
	rb.SetPosition(10, 10)
	rb.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		p1.SetSize(pwidth, pheight)
	})
	a.DemoPanel().Add(rb)

	// Horizontal panel
	params := gui.HBoxLayoutParams{Expand: 0, AlignV: gui.AlignTop}
	p1 = gui.NewPanel(pwidth, pheight)
	p1.SetPosition(rb.Position().X, rb.Position().Y+rb.Height()+4)
	p1.SetBorders(1, 1, 1, 1)
	p1.SetBordersColor(math32.NewColor("Black"))
	p1.SetPaddings(2, 2, 2, 2)
	a.DemoPanel().Add(p1)
	// Sets Horizontal box layout
	bl1 := gui.NewHBoxLayout()
	bl1.SetSpacing(4)
	p1.SetLayout(bl1)

	const bposy = 140
	// Add button
	p1b1 := gui.NewButton("Add")
	p1b1.SetPosition(10, bposy)
	p1b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		child := gui.NewButton(fmt.Sprintf("child %d", len(p1.Children())))
		offs := rand.Int31n(30)
		child.SetHeight(child.Height() + float32(offs))
		itemParams := params
		child.SetLayoutParams(&itemParams)
		p1.Add(child)
	})
	a.DemoPanel().Add(p1b1)
	// Del button
	p1b2 := gui.NewButton("Del")
	p1b2.SetPosition(p1b1.Position().X+p1b1.Width()+10, p1b1.Position().Y)
	p1b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if len(p1.Children()) > 0 {
			p1.Remove(p1.Children()[0].(gui.IPanel))
		}
	})
	a.DemoPanel().Add(p1b2)
	// Auto height checkbox
	cb1 := gui.NewCheckBox("Auto height")
	cb1.SetPosition(p1b2.Position().X+p1b2.Width()+10, p1b2.Position().Y)
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		bl1.SetAutoHeight(cb1.Value())
	})
	a.DemoPanel().Add(cb1)
	cb1.SetValue(true)
	// Auto width checkbox
	cb2 := gui.NewCheckBox("Auto width")
	cb2.SetPosition(cb1.Position().X+cb1.Width()+10, cb1.Position().Y)
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) { bl1.SetAutoWidth(cb2.Value()) })
	cb2.SetValue(true)
	a.DemoPanel().Add(cb2)
	// X left
	p1c1 := gui.NewRadioButton("Left")
	p1c1.SetPosition(p1b2.Position().X+p1b2.Width()+10, cb1.Position().Y+30)
	p1c1.SetGroup("alignH")
	p1c1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c1.Value() {
			bl1.SetAlignH(gui.AlignLeft)
		}
	})
	a.DemoPanel().Add(p1c1)
	// X center
	p1c2 := gui.NewRadioButton("Center")
	p1c2.SetPosition(p1c1.Position().X+p1c1.Width()+10, p1c1.Position().Y)
	p1c2.SetGroup("alignH")
	p1c2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c2.Value() {
			bl1.SetAlignH(gui.AlignCenter)
		}
	})
	a.DemoPanel().Add(p1c2)
	// X right
	p1c3 := gui.NewRadioButton("Right")
	p1c3.SetPosition(p1c2.Position().X+p1c2.Width()+10, p1c1.Position().Y)
	p1c3.SetGroup("alignH")
	p1c3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c3.Value() {
			bl1.SetAlignH(gui.AlignRight)
		}
	})
	a.DemoPanel().Add(p1c3)
	// X right
	p1c4 := gui.NewRadioButton("Width")
	p1c4.SetPosition(p1c3.Position().X+p1c3.Width()+10, p1c3.Position().Y)
	p1c4.SetGroup("alignH")
	p1c4.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c4.Value() {
			bl1.SetAlignH(gui.AlignWidth)
		}
	})
	a.DemoPanel().Add(p1c4)

	// Align next item to Top
	p1c5 := gui.NewRadioButton("Top")
	p1c5.SetPosition(p1b2.Position().X+p1b2.Width()+10, p1c1.Position().Y+30)
	p1c5.SetGroup("alignV")
	p1c5.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c5.Value() {
			params.AlignV = gui.AlignTop
		}
	})
	a.DemoPanel().Add(p1c5)
	// Align next item to Center
	p1c6 := gui.NewRadioButton("Center")
	p1c6.SetPosition(p1c5.Position().X+p1c5.Width()+10, p1c5.Position().Y)
	p1c6.SetGroup("alignV")
	p1c6.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c6.Value() {
			params.AlignV = gui.AlignCenter
		}
	})
	a.DemoPanel().Add(p1c6)
	// Align next item to Bottom
	p1c7 := gui.NewRadioButton("Bottom")
	p1c7.SetPosition(p1c6.Position().X+p1c6.Width()+10, p1c6.Position().Y)
	p1c7.SetGroup("alignV")
	p1c7.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c7.Value() {
			params.AlignV = gui.AlignBottom
		}
	})
	a.DemoPanel().Add(p1c7)
	// Align next item to Bottom
	p1c8 := gui.NewRadioButton("Height")
	p1c8.SetPosition(p1c7.Position().X+p1c7.Width()+10, p1c7.Position().Y)
	p1c8.SetGroup("alignV")
	p1c8.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if p1c8.Value() {
			params.AlignV = gui.AlignHeight
		}
	})
	a.DemoPanel().Add(p1c8)
	// Expand factor for next item
	cbexp := gui.NewCheckBox("Expand")
	cbexp.SetPosition(p1c8.Position().X+p1c8.Width()+10, p1c8.Position().Y)
	cbexp.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if cbexp.Value() {
			params.Expand = 1
		} else {
			params.Expand = 0
		}
	})
	a.DemoPanel().Add(cbexp)
}

// Update is called every frame.
func (t *GuiLayoutHBox) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiLayoutHBox) Cleanup(a *app.App) {}
