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
			cl1.SetScaleX(10, &math32.Color{0.8, 0.8, 0.8})
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

	//data1 := []float32{0.1, 0.5, 0.6}
	//cl1.AddGraph("g1", "Graph1", &math32.Color{0, 0, 1}, data1)

	//data2 := []float32{0.2, 0.8, 0.9, 0.6, 0.61, 0.63, 0.68, 0.63, 0.54}
	//cl1.AddGraph("g2", "Graph2", &math32.Color{1, 0, 0}, data2)

	//cl2 := gui.NewChartLine(500, 200)
	//cl2.SetBorders(2, 2, 2, 2)
	//cl2.SetBordersColor(&math32.Black)
	//cl2.SetColor(&math32.White)
	//cl2.SetPosition(10, 10+200+10)
	//cl2.SetScaleX(10, &math32.Color{0.8, 0.8, 0.8})
	//cl2.AddGraph("g1", "Graph1", &math32.Color{0, 0, 1}, data1)
	//ctx.Gui.Add(cl2)

	// Label
	//l1 := gui.NewLabel("This is the Graph1")
	//l1.SetPosition(10, 10)
	//cl.Add(l1)

	//// Panel
	//b1 := gui.NewButton("button 1")
	//b1.SetPosition(10, 10)
	//cl.Add(b1)
}

func (t *GuiChart) Render(ctx *Context) {

}
