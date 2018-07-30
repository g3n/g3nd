package gui

import (
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/engine/gui"
)

func init() {
	demos.Map["gui.custom_cursors"] = &CustomCursors{}
}

type CustomCursors struct {
	cursors []int
	current int
}

func (t *CustomCursors) Initialize(a *app.App) {

	t.cursors = make([]int, 2)
	t.current = 0

	instructions := gui.NewLabel("Click to change cursor!")
	instructions.SetPosition(50,50)
	a.GuiPanel().Add(instructions)

	var err error
	t.cursors[0], err = a.Window().Manager().CreateCursor(a.DirData() + "/images/gopher_cursor.png", 0, 0)
	t.cursors[1], err = a.Window().Manager().CreateCursor(a.DirData() + "/images/gauntlet_cursor.png", 0, 0)
	if err != nil {
		a.Log().Fatal("Error creating cursor: %s", err)
	}

	a.Window().SetCustomCursor(t.cursors[t.current])

	a.GuiPanel().Subscribe(gui.OnMouseDown, func(s string, i interface{}) {
		t.current += 1
		if t.current > len(t.cursors) - 1 {
			t.current = 0
		}
		a.Window().SetCustomCursor(t.cursors[t.current])
	})

}

func (t *CustomCursors) Render(a *app.App) {

}
