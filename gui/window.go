package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.window"] = &GuiWindow{}
}

type GuiWindow struct{}

func (t *GuiWindow) Initialize(a *app.App) {

	w1 := gui.NewWindow(300, 100)
	w1.SetPosition(10, 10)
	w1.SetResizable(true)
	w1.SetLayout(gui.NewFillLayout(true, true))
	c1 := gui.NewImageLabel(
		"This is the client area of the Window 1\n" +
			"This window does not have a title\n" +
			"It is resizable\n",
	)
	c1.SetColor(math32.NewColor("black"))
	c1.SetBgColor(math32.NewColor("white"))
	w1.Add(c1)
	a.GuiPanel().Add(w1)

	w2 := gui.NewWindow(300, 100)
	w2.SetTitle("Window2")
	w2.SetPosition(w1.Position().X+w1.Width()+50, 10)
	w2.SetResizable(true)
	w2.SetLayout(gui.NewFillLayout(true, true))
	c2 := gui.NewImageLabel(
		"This is the client area of the Window 2\n" +
			"This window has a title and a close button\n" +
			"It is resizable\n" +
			"It is movable by dragging its title",
	)
	c2.SetColor(math32.NewColor("black"))
	c2.SetBgColor(&math32.Color{0.7, 0.8, 0.9})
	w2.Add(c2)
	a.GuiPanel().Add(w2)

	w3 := gui.NewWindow(300, 100)
	w3.SetTitle("Window3")
	w3.SetCloseButton(false)
	w3.SetPosition(10, w1.Position().Y+w1.Height()+50)
	w3.SetResizable(true)
	w3.SetLayout(gui.NewFillLayout(true, true))
	c3 := gui.NewImageLabel(
		"This is the client area of the Window 3\n" +
			"This window has a title but no close button\n" +
			"It is resizable\n" +
			"It is movable by dragging its title",
	)
	c3.SetColor(math32.NewColor("black"))
	c3.SetBgColor(&math32.Color{0.8, 0.9, 0.9})
	w3.Add(c3)
	a.GuiPanel().Add(w3)
}

func (t *GuiWindow) Render(a *app.App) {
}
