package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["gui.chart"] = &GuiChart{}
}

type GuiChart struct{}

func (t *GuiChart) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	// Creates ChartLine panel
	cl1 := gui.NewChartLine(500, 200)
	cl1.SetMargins(10, 10, 10, 10)
	cl1.SetBorders(10, 10, 10, 10)
	cl1.SetBordersColor(&math32.Black)
	cl1.SetColor(&math32.White)
	cl1.SetPaddings(8, 8, 8, 8)
	cl1.SetPaddingsColor(&math32.Green)
	cl1.SetPosition(0, 0)
	ctx.Gui.Add(cl1)

	// Title checkbox
	cbTitle := gui.NewCheckBox("Chart Title")
	cbTitle.SetPosition(10, cl1.Height()+10)
	cbTitle.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbTitle.Value() {
			title := gui.NewLabel("Chart Title")
			title.SetBgColor(&math32.White)
			title.SetFontSize(16)
			cl1.SetTitle(title)
		} else {
			cl1.SetTitle(nil)
		}
	})
	ctx.Gui.Add(cbTitle)

	// X Scale checkbox
	cbX := gui.NewCheckBox("X Scale")
	cbX.SetPosition(cbTitle.Position().X+cbTitle.Width()+10, cbTitle.Position().Y)
	cbX.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbX.Value() {
			cl1.SetScaleX(4, &math32.Color{1.8, 0.8, 0.8})
		} else {
			cl1.ClearScaleX()
		}
	})
	ctx.Gui.Add(cbX)

	// Y Scale checkbox
	cbY := gui.NewCheckBox("Y Scale")
	cbY.SetPosition(cbX.Position().X+cbX.Width()+10, cbX.Position().Y)
	cbY.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbY.Value() {
			cl1.SetScaleY(10, &math32.Color{0.8, 0.8, 0.8})
		} else {
			cl1.ClearScaleY()
		}
	})
	ctx.Gui.Add(cbY)

	// Graph1
	var g1 *gui.LineGraph
	cbG1 := gui.NewCheckBox("Graph1")
	cbG1.SetPosition(cbY.Position().X+cbY.Width()+10, cbY.Position().Y)
	cbG1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbG1.Value() {
			data1 := []float32{0, 1, 1.5, 3, 20, 5, -8, 7, 7.5, 9, 9.5}
			cl1.SetRangeX(0, len(data1), 0, 1)
			cl1.SetRangeY(0, 12)
			g1 = cl1.AddGraph(&math32.Color{0, 0, 1}, data1)
			g1.SetLineWidth(2.0)
			cl1.SetScaleY(10, &math32.Color{0.8, 0.8, 0.8})
		} else {
			cl1.RemoveGraph(g1)
		}
	})
	ctx.Gui.Add(cbG1)

}

func (t *GuiChart) Render(ctx *Context) {

}
