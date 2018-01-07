package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

type GuiPanelChildren struct {
	p0 *gui.Panel
	p1 *gui.Panel
	p2 *gui.Panel
	p3 *gui.Panel
}

func init() {
	demos.Map["gui.panel_children"] = &GuiPanelChildren{}
}

func (t *GuiPanelChildren) Initialize(a *app.App) {

	// Panel 0
	t.p0 = gui.NewPanel(500, 300)
	t.p0.SetPosition(0, 0)
	t.p0.SetMargins(10, 10, 10, 10)
	t.p0.SetBorders(6, 6, 6, 6)
	t.p0.SetBordersColor(math32.NewColor("red"))
	t.p0.SetPaddings(8, 8, 8, 8)
	t.p0.SetColor(math32.NewColor("white"))
	t.p0.SetPaddingsColor(math32.NewColor("blue"))
	l0 := gui.NewLabel("P0")
	t.p0.Add(l0)
	a.GuiPanel().Add(t.p0)

	// Panel 1
	t.p1 = gui.NewPanel(400, 200)
	t.p1.SetPosition(20, 20)
	t.p1.SetBorders(6, 6, 6, 6)
	t.p1.SetBordersColor(math32.NewColor("red"))
	t.p1.SetPaddings(8, 8, 8, 8)
	t.p1.SetColor(math32.NewColor("white"))
	t.p1.SetPaddingsColor(math32.NewColor("green"))
	t.p1.SetName("p1")
	//app.Gui().Add(t.p0)
	t.p0.Add(t.p1)
	l1 := gui.NewLabel("P1(ASDW): Child of P0")
	t.p1.Add(l1)

	// Panel2
	t.p2 = gui.NewPanel(240, 80)
	t.p2.SetPosition(20, 20)
	t.p2.SetColor(&math32.Color{0.9, 0.9, 0.95})
	t.p2.SetBorders(6, 6, 6, 4)
	t.p2.SetBordersColor(math32.NewColor("red"))
	t.p2.SetPaddings(8, 8, 8, 8)
	t.p2.SetPaddingsColor(math32.NewColor("black"))
	l2 := gui.NewLabel("P2(JKLI): Child of P1")
	t.p2.Add(l2)
	t.p1.Add(t.p2)

	// Panel3
	t.p3 = gui.NewPanel(240, 80)
	t.p3.SetPosition(20, 150)
	t.p3.SetColor(&math32.Color{0.9, 0.9, 0.95})
	t.p3.SetBorders(6, 6, 6, 4)
	t.p3.SetBordersColor(math32.NewColor("red"))
	t.p3.SetPaddings(8, 8, 8, 8)
	t.p3.SetPaddingsColor(math32.NewColor("black"))
	t.p3.SetBounded(false)
	l3 := gui.NewLabel("P3(FGHT): Child of P1 (unbounded)")
	t.p3.Add(l3)
	t.p1.Add(t.p3)
	t.p1.SetTopChild(t.p2)

	// Subscribe to key events
	a.Window().Subscribe(window.OnKeyDown, t.onKey)
	a.Window().Subscribe(window.OnKeyRepeat, t.onKey)
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
	// Move panel 3
	case window.KeyF:
		px := t.p3.Position().X
		px -= step
		t.p3.SetPositionX(px)
	case window.KeyH:
		px := t.p3.Position().X
		px += step
		t.p3.SetPositionX(px)
	case window.KeyT:
		py := t.p3.Position().Y
		py -= step
		t.p3.SetPositionY(py)
	case window.KeyG:
		py := t.p3.Position().Y
		py += step
		t.p3.SetPositionY(py)
	}
}

func (t *GuiPanelChildren) Render(a *app.App) {

}
