package gui

import (
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["gui.panel_modal"] = &GuiPanelModal{}
}

type GuiPanelModal struct {
	panels []gui.IPanel
}

// Start is called once at the start of the demo.
func (t *GuiPanelModal) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	p1 := t.createPanel(a)
	p1.SetPosition(10, 10)
	a.DemoPanel().Add(p1)
	t.panels = append(t.panels, p1)

	p2 := t.createPanel(a)
	p2.SetPosition(p1.Position().X, p1.Position().Y+p1.Height()+40)
	a.DemoPanel().Add(p2)
	t.panels = append(t.panels, p2)

	p3 := t.createPanel(a)
	p3.SetPosition(p2.Position().X, p2.Position().Y+p2.Height()+40)
	a.DemoPanel().Add(p3)
	t.panels = append(t.panels, p3)
}

func (t *GuiPanelModal) setModal(a *app.App, p gui.IPanel) {

	gui.Manager().SetModal(p)
	// If no modal panels, sets all panels color to white
	if p == nil {
		for _, ipan := range t.panels {
			ipan.GetPanel().SetColor4(&math32.Color4{})
		}
		return
	}
	// Sets the modal panel color to white and others to gray
	for _, ipan := range t.panels {
		if ipan == p {
			ipan.GetPanel().SetColor4(&math32.Color4{})
			continue
		}
		ipan.GetPanel().SetColor(math32.NewColor("gray"))
	}
}

func (t *GuiPanelModal) createPanel(a *app.App) *gui.Panel {

	// Panel
	p := gui.NewPanel(0, 0)
	p.SetBorders(1, 1, 1, 1)
	p.SetPaddings(8, 8, 8, 8)
	layout := gui.NewHBoxLayout()
	layout.SetSpacing(10)
	layout.SetAlignH(gui.AlignCenter)
	layout.SetAutoHeight(true)
	layout.SetAutoWidth(true)
	p.SetLayout(layout)

	// Button 1
	b1 := gui.NewButton("Set Modal")
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		t.setModal(a, p)
	})
	p.Add(b1)

	// Button 2
	b2 := gui.NewButton("Clear Modal")
	b2.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		t.setModal(a, nil)
	})
	p.Add(b2)

	// Button 3
	b3 := gui.NewButton("button")
	p.Add(b3)

	// Checkbox
	cb1 := gui.NewCheckBox("checkbox")
	cb1.SetLayoutParams(&gui.HBoxLayoutParams{Expand: 0, AlignV: gui.AlignCenter})
	p.Add(cb1)

	// Dropdown
	dd1 := gui.NewDropDown(100, gui.NewImageLabel("dropdown"))
	dd1.SetLayoutParams(&gui.HBoxLayoutParams{Expand: 0, AlignV: gui.AlignCenter})
	p.Add(dd1)

	// Edit
	e1 := gui.NewEdit(100, "edit text")
	e1.SetLayoutParams(&gui.HBoxLayoutParams{Expand: 0, AlignV: gui.AlignCenter})
	p.Add(e1)

	return p
}

// Update is called every frame.
func (t *GuiPanelModal) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *GuiPanelModal) Cleanup(a *app.App) {}
