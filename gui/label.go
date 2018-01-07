package gui

import (
	"strings"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.label"] = &GuiLabel{}
}

type GuiLabel struct{}

func (t *GuiLabel) Initialize(a *app.App) {

	l1 := gui.NewLabel("label1")
	l1.SetPosition(10, 10)
	a.GuiPanel().Add(l1)

	l2 := gui.NewLabel("label2")
	l2.SetPosition(60, 10)
	l2.SetBorders(1, 1, 1, 1)
	l2.SetBordersColor(math32.NewColor("black"))
	l2.SetPaddings(2, 2, 2, 2)
	a.GuiPanel().Add(l2)

	l3 := gui.NewLabel("label3")
	l3.SetPosition(120, 10)
	l3.SetBgColor(math32.NewColor("green"))
	l3.SetBorders(1, 1, 1, 1)
	l3.SetPaddings(4, 6, 4, 6)
	a.GuiPanel().Add(l3)

	l4 := gui.NewLabel("label4")
	l4.SetPosition(200, 10)
	l4.SetBgColor(math32.NewColor("blue"))
	l4.SetColor(math32.NewColor("white"))
	l4.SetBorders(1, 1, 1, 1)
	l4.SetPaddings(4, 20, 4, 20)
	l4.SetFontSize(20)
	a.GuiPanel().Add(l4)

	l5 := gui.NewLabel("label5")
	l5.SetPosition(320, 10)
	l5.SetFontSize(28)
	l5.SetColor(math32.NewColor("red"))
	l5.SetBorders(1, 1, 1, 1)
	l5.SetBordersColor(math32.NewColor("white"))
	l5.SetPaddings(4, 20, 4, 20)
	l5.SetSize(100, 100)
	a.GuiPanel().Add(l5)

	l6 := gui.NewLabel("label6")
	l6.SetPosition(450, 10)
	l6.SetColor(math32.NewColor("red"))
	l6.SetBorders(1, 1, 1, 1)
	l6.SetBordersColor(math32.NewColor("white"))
	l6.SetPaddings(4, 20, 4, 20)
	l6.SetSize(100, 100)
	l6.SetFontSize(28)
	a.GuiPanel().Add(l6)

	lines := []string{
		"We are merely picking up pebbles on the beach",
		"while the great ocean of truth",
		"lays completely undiscovered before us.",
	}
	l7 := gui.NewLabel(strings.Join(lines, "\n"))
	l7.SetPosition(10, 120)
	l7.SetBordersColor(math32.NewColor("red"))
	l7.SetBgColor(math32.NewColor("green"))
	l7.SetColor(math32.NewColor("blue"))
	l7.SetBorders(10, 4, 10, 4)
	l7.SetPaddings(4, 20, 4, 20)
	l7.SetFontSize(22)
	a.GuiPanel().Add(l7)
}

func (t *GuiLabel) Render(a *app.App) {

}
