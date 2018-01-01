package gui

import (
	"fmt"
	"math"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["gui.chart"] = &GuiChart{}
}

type GuiChart struct{}

func (t *GuiChart) Initialize(app *g3nd.App) {

	// Creates Chart panel
	chart := gui.NewChart(0, 0)
	chart.SetMargins(10, 10, 10, 10)
	chart.SetBorders(2, 2, 2, 2)
	chart.SetBordersColor(math32.NewColor("black"))
	chart.SetColor(math32.NewColor("white"))
	chart.SetPaddings(8, 8, 8, 8)
	chart.SetPosition(0, 0)
	chart.SetFormatX("%3.1f")
	chart.SetFormatY("%2.1f")
	app.GuiPanel().Add(chart)

	minY := float32(-12.0)
	maxY := float32(12.0)
	startX := 0
	firstX := float32(0)
	stepX := float32(2.0)
	countStepX := float32(20)
	chart.SetRangeY(minY, maxY)

	// Graph1
	var g1 *gui.Graph
	data1 := make([]float32, 0)
	var x float32
	for x = 0; x < 2*math.Pi*10; x += 0.1 {
		data1 = append(data1, 10*math32.Sin(x)*math32.Sin(x/10))
	}
	cbG1 := gui.NewCheckBox("Graph1")
	cbG1.SetPosition(10, 10)
	cbG1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbG1.Value() {
			g1 = chart.AddLineGraph(&math32.Color{0, 0, 1}, data1)
			g1.SetLineWidth(2.0)
		} else {
			chart.RemoveGraph(g1)
			g1 = nil
		}
	})
	app.GuiPanel().Add(cbG1)
	cbG1.SetValue(true)

	// Graph2
	var g2 *gui.Graph
	data2 := make([]float32, 0)
	var x2 float32
	for x2 = 0; x2 < 2*math.Pi*10; x2 += 0.1 {
		data2 = append(data2, -2+5*math32.Cos(x2/3))
	}
	cbG2 := gui.NewCheckBox("Graph2")
	cbG2.SetPosition(cbG1.Position().X+cbG1.Width()+10, cbG1.Position().Y)
	cbG2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbG2.Value() {
			g2 = chart.AddLineGraph(&math32.Color{1, 0, 0}, data2)
			g2.SetLineWidth(2.0)
		} else {
			chart.RemoveGraph(g2)
			g2 = nil
		}
	})
	app.GuiPanel().Add(cbG2)

	// Title checkbox
	cbTitle := gui.NewCheckBox("Chart Title")
	cbTitle.SetPosition(cbG1.Position().X, cbG1.Position().Y+cbG1.Height()+10)
	cbTitle.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbTitle.Value() {
			chart.SetTitle("Chart Title", 16)
		} else {
			chart.SetTitle("", 0)
		}
	})
	cbTitle.SetValue(true)
	app.GuiPanel().Add(cbTitle)

	// X Scale checkbox
	cbX := gui.NewCheckBox("X Scale")
	cbX.SetPosition(cbTitle.Position().X+cbTitle.Width()+10, cbTitle.Position().Y)
	cbX.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbX.Value() {
			chart.SetScaleX(5, &math32.Color{0.8, 0.8, 0.8})
			chart.SetFontSizeX(13)
			chart.SetRangeX(firstX, stepX, countStepX)
		} else {
			chart.ClearScaleX()
		}
	})
	cbX.SetValue(true)
	app.GuiPanel().Add(cbX)

	// Y Scale checkbox
	cbY := gui.NewCheckBox("Y Scale")
	cbY.SetPosition(cbX.Position().X+cbX.Width()+10, cbX.Position().Y)
	cbY.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		if cbY.Value() {
			chart.SetScaleY(5, &math32.Color{0.8, 0.8, 0.8})
			chart.SetFontSizeY(13)
		} else {
			chart.ClearScaleY()
		}
	})
	cbY.SetValue(true)
	app.GuiPanel().Add(cbY)

	// startX ranger
	rStartx := newRanger(100, 20, 0, 2*math.Pi*10, "startX:%2.0f")
	rStartx.SetPosition(cbTitle.Position().X, cbTitle.Position().Y+cbTitle.Height()+10)
	rStartx.SetValue(0)
	rStartx.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		firstX = math32.Round(rStartx.Value())
		startX = int(firstX * 10)
		chart.SetRangeX(firstX, stepX, countStepX)
		if startX >= len(data1) {
			return
		}
		if g1 != nil {
			g1.SetData(data1[startX:])
		}
		if g2 != nil {
			g2.SetData(data2[startX:])
		}
	})
	app.GuiPanel().Add(rStartx)

	// step and countStepX ranger
	rStepx := newRanger(100, 20, 1, 20, "stepX:%2.0f")
	rStepx.SetPosition(rStartx.Position().X+rStartx.Width()+10, rStartx.Position().Y)
	rStepx.SetValue(float32(stepX))
	rStepx.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		stepX = float32(int(rStepx.Value()))
		countStepX = 10 * stepX
		chart.SetRangeX(firstX, stepX, countStepX)
	})
	app.GuiPanel().Add(rStepx)

	// minY ranger
	rMiny := newRanger(100, 20, -20, 20, "minY:%3.1f")
	rMiny.SetPosition(rStepx.Position().X+rStepx.Width()+10, rStepx.Position().Y)
	rMiny.SetValue(minY)
	rMiny.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		minY = rMiny.Value()
		chart.SetRangeY(minY, maxY)
	})
	app.GuiPanel().Add(rMiny)

	// maxY ranger
	rMaxy := newRanger(100, 20, -20, 20, "maxY:%3.1f")
	rMaxy.SetPosition(rMiny.Position().X+rMiny.Width()+10, rMiny.Position().Y)
	rMaxy.SetValue(maxY)
	rMaxy.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		maxY = rMaxy.Value()
		chart.SetRangeY(minY, maxY)
	})
	app.GuiPanel().Add(rMaxy)

	// Y auto range checkbox
	cbAutoy := gui.NewCheckBox("Y auto")
	cbAutoy.SetPosition(rMaxy.Position().X+rMaxy.Width()+10, rMaxy.Position().Y)
	cbAutoy.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if cbAutoy.Value() {
			chart.SetRangeYauto(true)
			minY, maxY = chart.RangeY()
			rMiny.SetEnabled(false)
			rMaxy.SetEnabled(false)
			rMiny.SetValue(minY)
			rMaxy.SetValue(maxY)
		} else {
			chart.SetRangeYauto(false)
			rMiny.SetEnabled(true)
			rMaxy.SetEnabled(true)

		}
	})
	app.GuiPanel().Add(cbAutoy)

	// Sets chart position and size
	chartY := cbAutoy.Position().Y + cbAutoy.Height() + 10
	chart.SetPosition(0, chartY)
	chart.SetSize(app.GuiPanel().ContentWidth(), app.GuiPanel().ContentHeight()-chartY)
	app.GuiPanel().Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		chart.SetSize(app.GuiPanel().ContentWidth(), app.GuiPanel().ContentHeight()-chartY)
	})
}

func (t *GuiChart) Render(app *g3nd.App) {

}

type ranger struct {
	*gui.Slider
	min    float32
	max    float32
	val    float32
	format string
}

func newRanger(width, height, min, max float32, format string) *ranger {

	r := new(ranger)
	r.Slider = gui.NewHSlider(width, height)
	r.min = min
	r.max = max
	r.format = format

	r.Slider.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		r.SetText(fmt.Sprintf(r.format, r.Value()))
	})
	return r
}

func (r *ranger) Value() float32 {

	return r.min + (r.max-r.min)*r.Slider.Value()
}

func (r *ranger) SetValue(val float32) {

	if val < r.min {
		r.val = r.min
	} else if val > r.max {
		r.val = r.max
	} else {
		r.val = val
	}
	sv := (r.val - r.min) / (r.max - r.min)
	r.SetText(fmt.Sprintf(r.format, r.val))
	r.Slider.SetValue(sv)
}
