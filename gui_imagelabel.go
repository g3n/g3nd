package main

import (
	"strings"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["gui.imagelabel"] = &GuiImageLabel{}
}

type GuiImageLabel struct{}

func (t *GuiImageLabel) Initialize(ctx *Context) {

	//axis := graphic.NewAxisHelper(1)
	//ctx.Scene.Add(axis)

	l1 := gui.NewImageLabel("label1")
	l1.SetPosition(10, 10)
	ctx.Gui.Add(l1)

	l2 := gui.NewImageLabel("label2")
	l2.SetPosition(60, 10)
	l2.SetBorders(1, 1, 1, 1)
	l2.SetBordersColor(math32.NewColor("black"))
	l2.SetPaddings(2, 2, 2, 2)
	ctx.Gui.Add(l2)

	l3 := gui.NewImageLabel("label3")
	l3.SetPosition(120, 10)
	l3.SetBgColor(math32.NewColor("green"))
	l3.SetBorders(1, 1, 1, 1)
	l3.SetPaddings(4, 6, 4, 6)
	ctx.Gui.Add(l3)

	l4 := gui.NewImageLabel("label4")
	l4.SetPosition(200, 10)
	l4.SetBgColor(math32.NewColor("blue"))
	//l4.SetBgAlpha(1)
	l4.SetColor(math32.NewColor("white"))
	l4.SetBorders(1, 1, 1, 1)
	l4.SetPaddings(4, 20, 4, 20)
	l4.SetFontSize(20)
	ctx.Gui.Add(l4)

	l5 := gui.NewImageLabel("label5")
	l5.SetPosition(320, 10)
	l5.SetFontSize(28)
	l5.SetColor(math32.NewColor("red"))
	l5.SetBorders(1, 1, 1, 1)
	l5.SetBordersColor(math32.NewColor("white"))
	l5.SetPaddings(4, 20, 4, 20)
	l5.SetSize(100, 100)
	ctx.Gui.Add(l5)

	l6 := gui.NewLabel("label6")
	l6.SetPosition(450, 10)
	l6.SetColor(math32.NewColor("red"))
	l6.SetBorders(1, 1, 1, 1)
	l6.SetBordersColor(math32.NewColor("white"))
	l6.SetPaddings(4, 20, 4, 20)
	l6.SetSize(100, 100)
	l6.SetFontSize(28)
	ctx.Gui.Add(l6)

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
	ctx.Gui.Add(l7)

	l8 := gui.NewImageLabel("label8")
	l8.SetPosition(10, l7.Position().Y+l7.Height()+30)
	l8.SetIcon(string(icon.ArrowBack))
	ctx.Gui.Add(l8)

	l9 := gui.NewImageLabel("label9")
	l9.SetIcon(string(icon.ArrowDownward))
	l9.SetPosition(100, l8.Position().Y)
	l9.SetBorders(1, 1, 1, 1)
	l9.SetBordersColor(math32.NewColor("black"))
	l9.SetPaddings(2, 2, 2, 2)
	ctx.Gui.Add(l9)

	l10 := gui.NewImageLabel("label10")
	l10.SetIcon(string(icon.ArrowDropDown))
	l10.SetPosition(200, l8.Position().Y)
	l10.SetBgColor(math32.NewColor("green"))
	l10.SetBorders(1, 1, 1, 1)
	l10.SetPaddings(4, 6, 4, 6)
	ctx.Gui.Add(l10)

	l11 := gui.NewImageLabel("label11")
	l11.SetPosition(300, l8.Position().Y)
	l11.SetImageFromFile(ctx.DirData + "/icons/add2.png")
	l11.SetBgColor(math32.NewColor("blue"))
	//l4.SetBgAlpha(1)
	l11.SetColor(math32.NewColor("white"))
	l11.SetBorders(1, 1, 1, 1)
	l11.SetPaddings(4, 20, 4, 20)
	l11.SetFontSize(20)
	ctx.Gui.Add(l11)

	l12 := gui.NewImageLabel("label12")
	img, err := gui.NewImage(ctx.DirData + "/images/tiger1.jpg")
	if err != nil {
		log.Fatal("%s", err)
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
	ctx.Gui.Add(l12)
}

func (t *GuiImageLabel) Render(ctx *Context) {
}
