package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"fmt"
)

func init() {
	demos.Map["gui.scroller"] = &Scroller{}
}

type Scroller struct{}

func (t *Scroller) Initialize(app *app.App) {

	// Scroller2

	img, _ := gui.NewImage(app.DirData() + "/images/uvgrid.jpg")
	imgOriginalSize := float32(512)
	img.SetSize(imgOriginalSize, imgOriginalSize)
	scroller := gui.NewScroller(300, 380, gui.ScrollBoth, img)
	scroller.SetPosition(10, 10)
	app.GuiPanel().Add(scroller)

	// ScrollMode radio buttons

	scrollMode1 := gui.NewRadioButton("ScrollNone")
	scrollMode1.SetPosition(10, 420)
	scrollMode1.Label.SetFontSize(14)
	scrollMode1.SetGroup("ScrollMode")
	scrollMode1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if scrollMode1.Value() == true {
			app.Log().Info("ScrollNone")
			scroller.SetScrollMode(gui.ScrollNone)
		}
	})
	app.GuiPanel().Add(scrollMode1)

	scrollMode2 := gui.NewRadioButton("ScrollVertical")
	scrollMode2.SetPosition(10, 445)
	scrollMode2.Label.SetFontSize(14)
	scrollMode2.SetGroup("ScrollMode")
	scrollMode2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if scrollMode2.Value() == true {
			app.Log().Info("ScrollVertical")
			scroller.SetScrollMode(gui.ScrollVertical)
		}
	})
	app.GuiPanel().Add(scrollMode2)

	scrollMode3 := gui.NewRadioButton("ScrollHorizontal")
	scrollMode3.SetPosition(10, 470)
	scrollMode3.Label.SetFontSize(14)
	scrollMode3.SetGroup("ScrollMode")
	scrollMode3.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if scrollMode3.Value() == true {
			app.Log().Info("ScrollHorizontal")
			scroller.SetScrollMode(gui.ScrollHorizontal)
		}
	})
	app.GuiPanel().Add(scrollMode3)

	scrollMode4 := gui.NewRadioButton("ScrollBoth")
	scrollMode4.SetPosition(10, 495)
	scrollMode4.Label.SetFontSize(14)
	scrollMode4.SetGroup("ScrollMode")
	scrollMode4.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if scrollMode4.Value() == true {
			app.Log().Info("ScrollBoth")
			scroller.SetScrollMode(gui.ScrollBoth)
		}
	})
	app.GuiPanel().Add(scrollMode4)

	// Interlocking style radio buttons

	interlocking1 := gui.NewRadioButton("ScrollbarInterlockingVertical")
	interlocking1.SetPosition(200, 420)
	interlocking1.Label.SetFontSize(14)
	interlocking1.SetGroup("ScrollbarInterlocking")
	interlocking1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if interlocking1.Value() == true {
			app.Log().Info("ScrollbarInterlockingVertical")
			scroller.SetScrollbarInterlocking(gui.ScrollbarInterlockingVertical)
		}
	})
	app.GuiPanel().Add(interlocking1)

	interlocking2 := gui.NewRadioButton("ScrollbarInterlockingHorizontal")
	interlocking2.SetPosition(200, 445)
	interlocking2.Label.SetFontSize(14)
	interlocking2.SetGroup("ScrollbarInterlocking")
	interlocking2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if interlocking2.Value() == true {
			app.Log().Info("ScrollbarInterlockingHorizontal")
			scroller.SetScrollbarInterlocking(gui.ScrollbarInterlockingHorizontal)
		}
	})
	app.GuiPanel().Add(interlocking2)

	interlocking3 := gui.NewRadioButton("ScrollbarInterlockingNone")
	interlocking3.SetPosition(200, 470)
	interlocking3.Label.SetFontSize(14)
	interlocking3.SetGroup("ScrollbarInterlocking")
	interlocking3.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if interlocking3.Value() == true {
			app.Log().Info("ScrollbarInterlockingNone")
			scroller.SetScrollbarInterlocking(gui.ScrollbarInterlockingNone)
		}
	})
	app.GuiPanel().Add(interlocking3)

	// CoverCorner Checkbox
	cb1 := gui.NewCheckBox("CornerCovered")
	cb1.SetPosition(200, 495)
	cb1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("CornerCovered Checkbox OnChange: %v", cb1.Value())
		scroller.SetCornerCovered(cb1.Value())
	})
	cb1.SetValue(true)
	app.GuiPanel().Add(cb1)

	// Vertical scrollbar options

	lv := gui.NewLabel("Vertical Scrollbar")
	lv.SetPosition(330, 25)
	app.GuiPanel().Add(lv)

	v1 := gui.NewCheckBox("Left/Right")
	v1.SetPosition(350, 50)
	v1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Left/Right Checkbox OnChange: %v", v1.Value())
		if v1.Value() == true {
			scroller.SetVerticalScrollbarPosition(gui.ScrollbarLeft)
		} else {
			scroller.SetVerticalScrollbarPosition(gui.ScrollbarRight)
		}
	})
	app.GuiPanel().Add(v1)

	v2 := gui.NewCheckBox("OverlapContent")
	v2.SetPosition(350, 75)
	v2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Vertical OverlapContent Checkbox OnChange: %v", v2.Value())
		scroller.SetVerticalScrollbarOverlapping(v2.Value())
	})
	app.GuiPanel().Add(v2)

	v3 := gui.NewCheckBox("AutoSizeButton")
	v3.SetPosition(350, 100)
	v3.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Vertical AutoSizeButton Checkbox OnChange: %v", v3.Value())
		scroller.SetVerticalScrollbarAutoSizeButton(v3.Value())
	})
	app.GuiPanel().Add(v3)

	vb := gui.NewHSlider(100, 28)
	vb.SetPosition(350, 125)
	vb.SetValue(0.16)
	vb.SetText(fmt.Sprintf("Broadness: %v", vb.Value() * vb.ContentWidth()))
	vb.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		//app.Log().Info("Slider 1 OnChange: %v", s1.Value())
		broadness := vb.Value() * vb.ContentWidth()
		vb.SetText(fmt.Sprintf("Broadness: %v", broadness))
		scroller.SetVerticalScrollbarBroadness(broadness)
	})
	app.GuiPanel().Add(vb)

	// Horizontal scrollbar options

	lh := gui.NewLabel("Horizontal Scrollbar")
	lh.SetPosition(330, 175)
	app.GuiPanel().Add(lh)

	h1 := gui.NewCheckBox("Top/Bottom")
	h1.SetPosition(350, 200)
	h1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Top/Bottom Checkbox OnChange: %v", h1.Value())
		if h1.Value() == true {
			scroller.SetHorizontalScrollbarPosition(gui.ScrollbarTop)
		} else {
			scroller.SetHorizontalScrollbarPosition(gui.ScrollbarBottom)
		}
	})
	app.GuiPanel().Add(h1)

	h2 := gui.NewCheckBox("OverlapContent")
	h2.SetPosition(350, 225)
	h2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Horizontal OverlapContent Checkbox OnChange: %v", h2.Value())
		scroller.SetHorizontalScrollbarOverlapping(h2.Value())
	})
	app.GuiPanel().Add(h2)

	h3 := gui.NewCheckBox("AutoSizeButton")
	h3.SetPosition(350, 250)
	h3.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Horizontal AutoSizeButton Checkbox OnChange: %v", h3.Value())
		scroller.SetHorizontalScrollbarAutoSizeButton(h3.Value())
	})
	app.GuiPanel().Add(h3)

	hb := gui.NewHSlider(100, 28)
	hb.SetPosition(350, 275)
	hb.SetValue(0.16)
	hb.SetText(fmt.Sprintf("Broadness: %v", hb.Value() * hb.ContentWidth()))
	hb.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		//app.Log().Info("Slider 1 OnChange: %v", s1.Value())
		broadness := hb.Value() * hb.ContentWidth()
		hb.SetText(fmt.Sprintf("Broadness: %v", broadness))
		scroller.SetHorizontalScrollbarBroadness(broadness)
	})
	app.GuiPanel().Add(hb)

	// Content area size controls

	cw := gui.NewHSlider(200, 28)
	cw.SetPosition(350, 330)
	cw.SetValue(1)
	cw.SetText(fmt.Sprintf("Content Width: %v", cw.Value() * imgOriginalSize))
	cw.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		//app.Log().Info("Slider 1 OnChange: %v", s1.Value())
		width := cw.Value() * imgOriginalSize
		cw.SetText(fmt.Sprintf("Content Width: %v", width))
		img.SetWidth(width)
		scroller.Update()
	})
	app.GuiPanel().Add(cw)

	ch := gui.NewHSlider(200, 28)
	ch.SetPosition(350, 370)
	ch.SetValue(1)
	ch.SetText(fmt.Sprintf("Content Height: %v", ch.Value() * imgOriginalSize))
	ch.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		//app.Log().Info("Slider 1 OnChange: %v", s1.Value())
		height := ch.Value() * imgOriginalSize
		ch.SetText(fmt.Sprintf("Content Height: %v", height))
		img.SetHeight(height)
		scroller.Update()
	})
	app.GuiPanel().Add(ch)

	// TODO Fix Radio Button start value bug
	//interlocking2.SetValue(true)

	// TODO add dropdown with style choices
}

func (t *Scroller) Render(a *app.App) {
}
