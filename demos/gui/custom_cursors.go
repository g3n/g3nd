package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["gui.custom_cursors"] = &CustomCursors{}
}

type CustomCursors struct {
	cursors []window.Cursor
	current int
}

// Start is called once at the start of the demo.
func (t *CustomCursors) Start(a *app.App) {

	// Show and enable demo panel
	a.DemoPanel().SetRenderable(true)
	a.DemoPanel().SetEnabled(true)

	t.cursors = make([]window.Cursor, 2)
	t.current = 0

	instructions := gui.NewLabel("Click to change cursor!")
	instructions.SetPosition(50, 50)
	a.DemoPanel().Add(instructions)

	var err error
	t.cursors[0], err = a.CreateCursor(a.DirData()+"/images/gopher_cursor.png", 0, 0)
	t.cursors[1], err = a.CreateCursor(a.DirData()+"/images/gauntlet_cursor.png", 0, 0)
	if err != nil {
		a.Log().Fatal("Error creating cursor: %s", err)
	}

	a.SetCursor(t.cursors[t.current])

	// Change cursor when clicking
	a.DemoPanel().SubscribeID(gui.OnMouseDown, a, func(s string, i interface{}) {
		t.current += 1
		if t.current > len(t.cursors)-1 {
			t.current = 0
		}
		a.SetCursor(t.cursors[t.current])
	})

}

// Update is called every frame.
func (t *CustomCursors) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *CustomCursors) Cleanup(a *app.App) {}
