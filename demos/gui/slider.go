package gui

import (
	"fmt"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.slider"] = &GuiSlider{}
}

type GuiSlider struct{}

// Start is called once at the start of the demo.
func (t *GuiSlider) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	// Slider 1
	s1 := gui.NewHSlider(400, 32)
	s1.SetPosition(10, 10)
	s1.SetValue(0.2)
	s1.SetText(fmt.Sprintf("%1.2f", s1.Value()))
	s1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		a.Log().Info("Slider 1 OnChange: %v", s1.Value())
		s1.SetText(fmt.Sprintf("%1.2f", s1.Value()))
	})
	a.DemoPanel().Add(s1)

	// Slider 2
	s2 := gui.NewHSlider(400, 32)
	s2.SetPosition(10, s1.Position().Y+s1.Height()+10)
	s2.SetValue(0.4)
	s2.SetEnabled(false)
	s2.SetText(fmt.Sprintf("%1.2f", s2.Value()))
	s2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		a.Log().Info("Slider 2 OnChange: %v", s2.Value())
		s2.SetText(fmt.Sprintf("%1.2f", s2.Value()))
	})
	a.DemoPanel().Add(s2)
	// Enable control
	cb1 := gui.NewCheckBox("Enabled")
	cb1.SetPosition(s2.Position().X+s2.Width()+10, s2.Position().Y)
	cb1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		s2.SetEnabled(cb1.Value())
	})
	a.DemoPanel().Add(cb1)

	// Slider 3
	s3 := gui.NewVSlider(64, 400)
	s3.SetPosition(10, s2.Position().Y+s2.Height()+20)
	s3.SetValue(0.5)
	s3.SetText(fmt.Sprintf("%1.2f", s3.Value()))
	s3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		a.Log().Info("Slider 3 OnChange: %v", s3.Value())
		s3.SetText(fmt.Sprintf("%1.2f", s3.Value()))
	})
	a.DemoPanel().Add(s3)
}

// Update is called every frame.
func (t *GuiSlider) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiSlider) Cleanup(a *app.App) {}
