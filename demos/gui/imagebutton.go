package gui

import (
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.imagebutton"] = &GuiImageButton{}
}

type GuiImageButton struct{}

// Start is called once at the start of the demo.
func (t *GuiImageButton) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// Large image button
	b1, err := gui.NewImageButton(a.DirData() + "/images/tiger1.jpg")
	if err != nil {
		panic(err)
	}
	b1.SetPosition(20, 20)
	a.DemoPanel().Add(b1)

	// Tiny image button
	b2, err := gui.NewImageButton(a.DirData() + "/images/ok.png")
	if err != nil {
		panic(err)
	}
	b2.SetPosition(b1.Panel.Position().X+b1.Panel.Width()+50, 20)
	a.DemoPanel().Add(b2)

	// Image button with text and multiple states
	b3, err := gui.NewImageButton(a.DirData() + "/images/blue_normal.png")
	if err != nil {
		panic(err)
	}
	b3.SetPosition(20, b1.Panel.Position().Y+b1.Panel.Height()+30)
	b3.SetText("LE TIGER")
	b3.SetFontSize(20)
	err = b3.SetImage(gui.ButtonOver, a.DirData()+"/images/blue_over.png")
	if err != nil {
		panic(err)
	}
	err = b3.SetImage(gui.ButtonPressed, a.DirData()+"/images/blue_pressed.png")
	if err != nil {
		panic(err)
	}

	bs := gui.ImageButtonStyle{}
	bs.Border = gui.RectBounds{0, 0, 0, 0}
	bs.Padding = gui.RectBounds{0, 0, 0, 0}
	bs.BorderColor = math32.Color4{0, 0, 0, 1}
	bs.BgColor = math32.Color4{0, 0, 0, 0}
	bs.FgColor = math32.Color4{0.85, 0.85, 0.85, 1}
	bss := gui.ImageButtonStyles{
		Normal:   bs,
		Over:     bs,
		Focus:    bs,
		Pressed:  bs,
		Disabled: bs,
	}

	b3.SetStyles(&bss)
	a.DemoPanel().Add(b3)

	// Image button with icon
	b4, err := gui.NewImageButton(a.DirData() + "/images/sprite0.png")
	if err != nil {
		panic(err)
	}
	b4.SetPosition(b2.Panel.Position().X+b2.Panel.Width()+50, 20)
	b4.SetIcon(icon.CheckBoxOutlineBlank)
	b4.SetFontSize(60)
	checked := false
	b4.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		if checked {
			checked = false
			b4.SetIcon(icon.CheckBoxOutlineBlank)
		} else {
			checked = true
			b4.SetIcon(icon.CheckBox)
		}
	})
	a.DemoPanel().Add(b4)

}

// Update is called every frame.
func (t *GuiImageButton) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiImageButton) Cleanup(a *app.App) {}
