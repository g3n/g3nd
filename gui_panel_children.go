package main

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
)

type GuiPanelChildren struct {
	p0 *gui.Panel
	p1 *gui.Panel
	p2 *gui.Panel
}

func init() {
	TestMap["gui.panel_children"] = &GuiPanelChildren{}
}

func (t *GuiPanelChildren) Initialize(ctx *Context) {

	// Panel 0
	t.p0 = gui.NewPanel(500, 300)
	t.p0.SetPosition(0, 0)
	t.p0.SetMargins(10, 10, 10, 10)
	t.p0.SetBorders(6, 6, 6, 6)
	t.p0.SetBordersColor(&math32.Red)
	t.p0.SetPaddings(8, 8, 8, 8)
	t.p0.SetColor(&math32.White)
	t.p0.SetPaddingsColor(&math32.Blue)
	l0 := gui.NewLabel("P0")
	t.p0.Add(l0)
	ctx.Gui.Add(t.p0)

	// Panel 1
	t.p1 = gui.NewPanel(400, 200)
	t.p1.SetPosition(20, 20)
	t.p1.SetBorders(6, 6, 6, 6)
	t.p1.SetBordersColor(&math32.Red)
	t.p1.SetPaddings(8, 8, 8, 8)
	t.p1.SetColor(&math32.White)
	t.p1.SetPaddingsColor(&math32.Green)
	t.p1.SetName("p1")
	ctx.Gui.Add(t.p0)
	t.p0.Add(t.p1)
	l1 := gui.NewLabel("P1(ASDW): Child of P0")
	t.p1.Add(l1)

	// Panel2
	t.p2 = gui.NewPanel(200, 100)
	t.p2.SetPosition(20, 20)
	t.p2.SetColor(&math32.Color{0.9, 0.9, 0.95})
	t.p2.SetBorders(6, 6, 6, 4)
	t.p2.SetBordersColor(&math32.Red)
	t.p2.SetPaddings(8, 8, 8, 8)
	t.p2.SetPaddingsColor(&math32.Black)
	l2 := gui.NewLabel("P2(JKLI): Child of P1")
	t.p2.Add(l2)
	t.p1.Add(t.p2)

	// Subscribe to key events
	ctx.Win.Subscribe(window.OnKeyDown, t.onKey)
	ctx.Win.Subscribe(window.OnKeyRepeat, t.onKey)
}

func (t *GuiPanelChildren) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	if kev.Action == window.Release {
		return
	}
	const step = 2
	switch kev.Keycode {
	// Move panel 1
	case window.KeyA:
		px := t.p1.Position().X
		px -= step
		t.p1.SetPositionX(px)
	case window.KeyD:
		px := t.p1.Position().X
		px += step
		t.p1.SetPositionX(px)
	case window.KeyW:
		py := t.p1.Position().Y
		py -= step
		t.p1.SetPositionY(py)
	case window.KeyS:
		py := t.p1.Position().Y
		py += step
		t.p1.SetPositionY(py)
	// Move panel 2
	case window.KeyJ:
		px := t.p2.Position().X
		px -= step
		t.p2.SetPositionX(px)
	case window.KeyL:
		px := t.p2.Position().X
		px += step
		t.p2.SetPositionX(px)
	case window.KeyI:
		py := t.p2.Position().Y
		py -= step
		t.p2.SetPositionY(py)
	case window.KeyK:
		py := t.p2.Position().Y
		py += step
		t.p2.SetPositionY(py)
	}
}

func (t *GuiPanelChildren) Render(ctx *Context) {

}
