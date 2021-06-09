package gui

import (
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.panel_children"] = &GuiPanelChildren{}
}

type GuiPanelChildren struct {
	p0 *gui.Panel
	p1 *gui.Panel
	p2 *gui.Panel
	p3 *gui.Panel
}

// Start is called once at the start of the demo.
func (t *GuiPanelChildren) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

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
	l0.SetColor(math32.NewColor("black"))
	t.p0.Add(l0)
	a.DemoPanel().Add(t.p0)

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
	l1.SetColor(math32.NewColor("black"))
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
	l2.SetColor(math32.NewColor("black"))
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
	l3.SetColor(math32.NewColor("black"))
	t.p3.Add(l3)
	t.p1.Add(t.p3)
	t.p1.SetTopChild(t.p2)

	// Subscribe to key events
	a.SubscribeID(window.OnKeyDown, a, t.onKey)
	a.SubscribeID(window.OnKeyRepeat, a, t.onKey)
}

func (t *GuiPanelChildren) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	const step = 2
	switch kev.Key {
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

// Update is called every frame.
func (t *GuiPanelChildren) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiPanelChildren) Cleanup(a *app.App) {}
