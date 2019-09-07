package gui

import (
	"strings"
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.imagelabel"] = &GuiImageLabel{}
}

type GuiImageLabel struct{}

// Start is called once at the start of the demo.
func (t *GuiImageLabel) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	l1 := gui.NewImageLabel("label1")
	l1.SetPosition(10, 10)
	a.DemoPanel().Add(l1)

	l2 := gui.NewImageLabel("label2")
	l2.SetPosition(60, 10)
	l2.SetBorders(1, 1, 1, 1)
	l2.SetBordersColor(math32.NewColor("black"))
	l2.SetPaddings(2, 2, 2, 2)
	a.DemoPanel().Add(l2)

	l3 := gui.NewImageLabel("label3")
	l3.SetPosition(120, 10)
	l3.SetBgColor(math32.NewColor("green"))
	l3.SetBorders(1, 1, 1, 1)
	l3.SetPaddings(4, 6, 4, 6)
	a.DemoPanel().Add(l3)

	l4 := gui.NewImageLabel("label4")
	l4.SetPosition(200, 10)
	l4.SetBgColor(math32.NewColor("blue"))
	//l4.SetBgAlpha(1)
	l4.SetColor(math32.NewColor("white"))
	l4.SetBorders(1, 1, 1, 1)
	l4.SetPaddings(4, 20, 4, 20)
	l4.SetFontSize(20)
	a.DemoPanel().Add(l4)

	l5 := gui.NewImageLabel("label5")
	l5.SetPosition(320, 10)
	l5.SetFontSize(28)
	l5.SetColor(math32.NewColor("red"))
	l5.SetBorders(1, 1, 1, 1)
	l5.SetBordersColor(math32.NewColor("white"))
	l5.SetPaddings(4, 20, 4, 20)
	l5.SetSize(100, 100)
	a.DemoPanel().Add(l5)

	l6 := gui.NewLabel("label6")
	l6.SetPosition(450, 10)
	l6.SetColor(math32.NewColor("red"))
	l6.SetBorders(1, 1, 1, 1)
	l6.SetBordersColor(math32.NewColor("white"))
	l6.SetPaddings(4, 20, 4, 20)
	l6.SetSize(100, 100)
	l6.SetFontSize(28)
	a.DemoPanel().Add(l6)

	lines := []string{
		"We are merely picking up pebbles on the beach",
		"while the great ocean of truth",
		"lays completely undiscovered before us.",
	}
	l7 := gui.NewImageLabel(strings.Join(lines, "\n"))
	l7.SetPosition(10, 120)
	l7.SetBordersColor(math32.NewColor("red"))
	l7.SetBgColor(math32.NewColor("green"))
	l7.SetColor(math32.NewColor("blue"))
	l7.SetBorders(10, 4, 10, 4)
	l7.SetPaddings(4, 20, 4, 20)
	l7.SetFontSize(22)
	a.DemoPanel().Add(l7)

	l8 := gui.NewImageLabel("label8")
	l8.SetPosition(10, l7.Position().Y+l7.Height()+30)
	l8.SetIcon(string(icon.ArrowBack))
	a.DemoPanel().Add(l8)

	l9 := gui.NewImageLabel("label9")
	l9.SetIcon(string(icon.ArrowDownward))
	l9.SetPosition(100, l8.Position().Y)
	l9.SetBorders(1, 1, 1, 1)
	l9.SetBordersColor(math32.NewColor("black"))
	l9.SetPaddings(2, 2, 2, 2)
	a.DemoPanel().Add(l9)

	l10 := gui.NewImageLabel("label10")
	l10.SetIcon(string(icon.ArrowDropDown))
	l10.SetPosition(200, l8.Position().Y)
	l10.SetBgColor(math32.NewColor("green"))
	l10.SetBorders(1, 1, 1, 1)
	l10.SetPaddings(4, 6, 4, 6)
	a.DemoPanel().Add(l10)

	l11 := gui.NewImageLabel("label11")
	l11.SetPosition(300, l8.Position().Y)
	l11.SetImageFromFile(a.DirData() + "/icons/add2.png")
	l11.SetBgColor(math32.NewColor("blue"))
	//l4.SetBgAlpha(1)
	l11.SetColor(math32.NewColor("white"))
	l11.SetBorders(1, 1, 1, 1)
	l11.SetPaddings(4, 20, 4, 20)
	l11.SetFontSize(20)
	a.DemoPanel().Add(l11)

	l12 := gui.NewImageLabel("label12")
	img, err := gui.NewImage(a.DirData() + "/images/tiger1.jpg")
	if err != nil {
		a.Log().Fatal("%s", err)
	}
	img.SetContentAspectWidth(64)
	l12.SetImage(img)
	l12.SetPosition(460, l8.Position().Y)
	l12.SetFontSize(28)
	l12.SetColor(math32.NewColor("red"))
	l12.SetBorders(1, 1, 1, 1)
	l12.SetBordersColor(math32.NewColor("white"))
	l12.SetPaddings(4, 20, 4, 20)
	l12.SetSize(100, 100)
	a.DemoPanel().Add(l12)
}

// Update is called every frame.
func (t *GuiImageLabel) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiImageLabel) Cleanup(a *app.App) {}
