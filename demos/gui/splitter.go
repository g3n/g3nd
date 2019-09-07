package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["gui.splitter"] = &GuiSplitter{}
}

type GuiSplitter struct{}

// Start is called once at the start of the demo.
func (t *GuiSplitter) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	s1 := gui.NewHSplitter(400, 200)
	s1.SetPosition(10, 10)
	s1.P0.SetColor(&math32.Color{1, 0, 0})
	s1.P0.SetBorders(1, 0, 1, 1)
	s1.P0.SetBordersColor(math32.NewColor("black"))
	s1.P0.Add(gui.NewLabel("This is the left panel\nof the splitter"))

	s1.P1.SetColor4(&math32.Color4{0, 0, 0, 0})
	s1.P1.SetBorders(1, 1, 1, 0)
	s1.P1.SetBordersColor(math32.NewColor("black"))
	s1.P1.Add(gui.NewLabel("This is the right panel\nof the splitter"))
	s1.SetSplit(0.75)
	a.DemoPanel().Add(s1)

	s2 := gui.NewVSplitter(400, 200)
	s2.SetPosition(10, 300)
	s2.P0.SetBorders(1, 1, 0, 1)
	s2.P0.SetBordersColor(math32.NewColor("black"))
	s2.P0.SetColor(&math32.Color{0, 1, 0})
	s2.P0.Add(gui.NewLabel("This is the top panel\nof the splitter"))

	s2.P1.SetColor4(&math32.Color4{0, 0, 0, 0})
	s2.P1.SetBorders(0, 1, 1, 1)
	s2.P1.SetBordersColor(math32.NewColor("black"))
	s2.P1.Add(gui.NewLabel("This is the bottom panel\nof the splitter"))
	s2.SetSplit(0.25)
	a.DemoPanel().Add(s2)
}

// Update is called every frame.
func (t *GuiSplitter) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiSplitter) Cleanup(a *app.App) {}
