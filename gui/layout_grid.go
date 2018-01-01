package gui

import (
	"fmt"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["gui.layout_grid"] = &GuiLayoutGrid{}
}

type GuiLayoutGrid struct {
	pan     *gui.Panel      // main panel
	layout  *gui.GridLayout // panel layout
	bwidth  float32         // next child width
	bheight float32         // next child height
	sw      *gui.Slider     // slider for child width
	sh      *gui.Slider     // slider for child height
	colspan *int            // next child colspan
	ddcs    *gui.DropDown   // drop down for colspan
	menu    *gui.Menu       // child menu
}

func (t *GuiLayoutGrid) Initialize(app *g3nd.App) {

	// Creates menu with child options
	t.menu = gui.NewMenu()
	t.menu.SetBounded(false)
	t.menu.SetVisible(false)
	app.GuiPanel().Add(t.menu)
	options := []string{
		"Hide", "",
		"Align left", "Align center", "Align right", "",
		"Align top", "Align middle", "Align bottom", "",
		"Colspan 0", "Colspan 1", "Colspan 2", "Colspan 3", "Colspan 4", "Colspan 5",
	}
	for _, op := range options {
		if op != "" {
			t.menu.AddOption(op).SetId(op)
		} else {
			t.menu.AddSeparator()
		}
	}
	t.menu.Subscribe(gui.OnMouseOut, func(evname string, ev interface{}) {
		t.menu.SetVisible(false)
	})
	t.menu.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		child := t.menu.UserData().(*gui.ImageLabel)
		t.menu.SetVisible(false)
		opid := ev.(*gui.MenuItem).Id()
		// Get child layout parameters
		iparam := child.LayoutParams()
		var params *gui.GridLayoutParams
		if iparam == nil {
			params = &gui.GridLayoutParams{}
			child.SetLayoutParams(params)
		} else {
			params = iparam.(*gui.GridLayoutParams)
		}
		// Process menu options
		switch opid {
		case "Hide":
			child.SetVisible(false)
		case "Align left":
			params.AlignH = gui.AlignLeft
		case "Align center":
			params.AlignH = gui.AlignCenter
		case "Align right":
			params.AlignH = gui.AlignRight

		case "Align top":
			params.AlignV = gui.AlignTop
		case "Align middle":
			params.AlignV = gui.AlignCenter
		case "Align bottom":
			params.AlignV = gui.AlignBottom

		case "Colspan 0":
			params.ColSpan = 0
		case "Colspan 1":
			params.ColSpan = 1
		case "Colspan 2":
			params.ColSpan = 2
		case "Colspan 3":
			params.ColSpan = 3
		case "Colspan 4":
			params.ColSpan = 4
		case "Colspan 5":
			params.ColSpan = 5
		}
		// Force layout recalculation
		t.layout.Recalc(t.pan)
	})

	// Button to add child to panel
	b1 := gui.NewButton("Add")
	b1.SetPosition(10, 10)
	b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		child := gui.NewImageLabel(fmt.Sprintf("child %d", len(t.pan.Children())))
		child.SetBorders(1, 1, 1, 1)
		child.SetPaddings(2, 2, 2, 2)
		child.SetBgColor(&math32.Color{1, 1, 0})
		child.SetWidth(child.Width() + t.bwidth)
		child.SetHeight(child.Height() + t.bheight)
		child.Subscribe(gui.OnMouseDown, func(evname string, ev interface{}) {
			px := t.pan.Position().X + child.Position().X + child.Width()/2
			py := t.pan.Position().Y + child.Position().Y + child.Height()/2
			t.menu.SetPosition(px, py)
			t.menu.SetUserData(child)
			t.menu.SetVisible(true)
		})
		var params gui.GridLayoutParams
		if t.colspan != nil {
			params.ColSpan = *t.colspan
			child.SetLayoutParams(&params)
		}
		t.pan.Add(child)
		// Reset next child parameters
		t.colspan = nil
		t.ddcs.SelectPos(0)
	})
	app.GuiPanel().Add(b1)

	// Clear button
	b2 := gui.NewButton("Clear")
	b2.SetPosition(b1.Position().X+b1.Width()+10, 10)
	b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		t.pan.DisposeChildren(true)
		t.colspan = nil
		t.sw.SetValue(0)
		t.sh.SetValue(0)
	})
	app.GuiPanel().Add(b2)

	// Slider for child width
	t.sw = gui.NewHSlider(42, 22)
	t.sw.SetPosition(b2.Position().X+b2.Width()+10, 10)
	t.sw.SetText("width")
	t.sw.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.bwidth = t.sw.Value() * 100
	})
	app.GuiPanel().Add(t.sw)

	// Slider for child height
	t.sh = gui.NewHSlider(42, 22)
	t.sh.SetPosition(t.sw.Position().X+t.sw.Width()+10, 10)
	t.sh.SetText("height")
	t.sh.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.bheight = t.sh.Value() * 100
	})
	app.GuiPanel().Add(t.sh)

	// Colspan dropdown
	t.ddcs = gui.NewDropDown(100, gui.NewImageLabel("colspan"))
	t.ddcs.SetPosition(t.sh.Position().X+t.sh.Width()+10, 10)
	for i := 0; i <= 6; i++ {
		iml := gui.NewImageLabel(fmt.Sprintf("%d columns", i))
		iml.SetUserData(i)
		t.ddcs.Add(iml)
	}
	t.ddcs.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		sel := t.ddcs.Selected()
		v := sel.UserData().(int)
		if v != 0 {
			if t.colspan == nil {
				t.colspan = new(int)
				*t.colspan = v
			}
		} else {
			t.colspan = nil
		}
	})
	app.GuiPanel().Add(t.ddcs)

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
	app.GuiPanel().Add(dd1)

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
	app.GuiPanel().Add(dd2)

	// CheckBox for horizontal expansion
	cb1 := gui.NewCheckBox("ExpandH")
	cb1.SetPosition(dd2.Position().X+dd2.Width()+10, dd2.Position().Y)
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.layout.SetExpandH(cb1.Value())
	})
	app.GuiPanel().Add(cb1)

	// CheckBox for vertical expansion
	cb2 := gui.NewCheckBox("ExpandV")
	cb2.SetPosition(cb1.Position().X+cb1.Width()+10, cb1.Position().Y)
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.layout.SetExpandV(cb2.Value())
	})
	app.GuiPanel().Add(cb2)

	// Creates panel with grid layout
	t.pan = gui.NewPanel(600, 400)
	t.pan.SetPosition(10, dd1.Position().Y+dd1.Height()+10)
	t.pan.SetColor(math32.NewColor("white"))
	t.pan.SetBorders(1, 1, 1, 1)
	t.layout = gui.NewGridLayout(6)
	t.pan.SetLayout(t.layout)
	app.GuiPanel().Add(t.pan)
}

func (t *GuiLayoutGrid) Render(app *g3nd.App) {
}
