package main

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
)

type GuiPanelRounded struct {
	p1 *gui.Panel
	p2 *gui.Panel
	p3 *gui.Panel
	p4 *gui.Panel
	p5 *gui.Panel
	p6 *gui.Panel
}

func init() {
	TestMap["gui.panel_rounded"] = &GuiPanelRounded{}
}

func (t *GuiPanelRounded) Initialize(ctx *Context) {

	// Panel 1
	t.p1 = gui.NewPanel(100, 100)
	t.p1.SetPosition(10, 10)
	t.p1.SetMargins(2, 2, 2, 2)
	t.p1.SetBorders(2, 2, 2, 2)
	t.p1.SetPaddings(2, 2, 2, 2)
	t.p1.SetBordersColor(math32.NewColor("black"))
	t.p1.SetColor(math32.NewColor("yellow"))
	t.p1.SetPaddingsColor(math32.NewColor("green"))
	ctx.Gui.Add(t.p1)

	// Panel 2
	t.p2 = gui.NewPanel(200, 100)
	t.p2.SetPosition(t.p1.Position().X+t.p1.Width()+10, t.p1.Position().Y)
	t.p2.SetMargins(2, 2, 2, 2)
	t.p2.SetBorders(4, 4, 4, 4)
	t.p2.SetBordersColor(math32.NewColor("black"))
	t.p2.SetPaddings(2, 2, 2, 2)
	t.p2.SetColor(math32.NewColor("green"))
	t.p2.SetPaddingsColor(math32.NewColor("white"))
	ctx.Gui.Add(t.p2)

	// Panel 3
	t.p3 = gui.NewPanel(100, 200)
	t.p3.SetPosition(t.p2.Position().X+t.p2.Width()+10, t.p2.Position().Y)
	t.p3.SetRoundness(0.5, 0.5, 0.5, 0.5)
	t.p3.SetMargins(2, 2, 2, 2)
	t.p3.SetBorders(4, 4, 4, 4)
	t.p3.SetBordersColor(math32.NewColor("black"))
	t.p3.SetPaddings(2, 2, 2, 2)
	t.p3.SetColor(math32.NewColor("yellow"))
	t.p3.SetPaddingsColor(math32.NewColor("blue"))
	ctx.Gui.Add(t.p3)

	// Panel 4
	t.p4 = gui.NewPanel(200, 100)
	t.p4.SetPosition(t.p1.Position().X, t.p3.Position().Y+t.p3.Height()+10)
	t.p4.SetBorders(2, 2, 2, 2)
	t.p4.SetBordersColor(math32.NewColor("red"))
	t.p4.SetPaddings(2, 2, 2, 2)
	t.p4.SetColor(math32.NewColor("white"))
	t.p4.SetPaddingsColor(math32.NewColor("green"))
	ctx.Gui.Add(t.p4)

	// Panel 5
	t.p5 = gui.NewPanel(200, 100)
	t.p5.SetPosition(t.p4.Position().X+t.p4.Width()+10, t.p4.Position().Y)
	t.p5.SetBorders(2, 2, 2, 2)
	t.p5.SetBordersColor(math32.NewColor("black"))
	t.p5.SetPaddings(2, 2, 2, 2)
	t.p5.SetColor(math32.NewColor("green"))
	t.p5.SetPaddingsColor(math32.NewColor("yellow"))
	ctx.Gui.Add(t.p5)

	// Panel 5
	t.p6 = gui.NewPanel(100, 200)
	t.p6.SetPosition(t.p5.Position().X+t.p5.Width()+10, t.p5.Position().Y)
	t.p6.SetBorders(2, 2, 2, 2)
	t.p6.SetBordersColor(math32.NewColor("black"))
	t.p6.SetPaddings(2, 2, 2, 2)
	t.p6.SetColor(math32.NewColor("purple"))
	t.p6.SetPaddingsColor(math32.NewColor("yellow"))
	ctx.Gui.Add(t.p6)
}

func (t *GuiPanelRounded) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	if kev.Action == window.Release {
		return
	}
	switch kev.Keycode {
	case window.KeyB:
		for _, ichild := range t.p3.Children() {
			pan := ichild.(*gui.Panel)
			pan.SetBounded(!pan.Bounded())
		}
	}

}

func (t *GuiPanelRounded) Render(ctx *Context) {

	//	time := ctx.Win.GetTime()
	//	delta := float32((math.Sin(time) + 1.0) / 2)
	//	t.p1.SetRoundness(delta, delta, delta, delta)
	//	t.p2.SetRoundness(delta, delta, delta, delta)
	//	t.p3.SetRoundness(delta, delta, delta, delta)
	//	t.p4.SetRoundness(delta, delta, 0, 0)
	//	t.p5.SetRoundness(0, delta, delta, 0)
	//	t.p6.SetRoundness(0, delta, 0, delta*2)
	//	//	maxWidth := 200
	//	//	maxHeight := 100
	//	//	t.p1.SetSize(float32(delta)*float32(maxWidth), float32(delta)*float32(maxHeight))
	//	//	t.p2.SetContentSize(float32(delta)*float32(maxWidth), float32(delta)*float32(maxHeight))
	//	//	t.p3.SetPosition(50+60*math32.Sin(float32(time)), t.p3.Position().Y)
}
