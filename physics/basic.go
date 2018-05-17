package physicssss

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/engine/window"
)

func init() {
	demos.Map["physics.basic"] = &PhysicsBasic{}
}

type PhysicsBasic struct {
	// TODO
}

func (t *PhysicsBasic) Initialize(a *app.App) {

	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)
}

func (t *PhysicsBasic) Render(a *app.App) {

}

func (t *PhysicsBasic) onKey(a *app.App, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	if kev.Action == window.Release {
		return
	}
	switch kev.Keycode {
	case window.Key1:
		// TODO
	case window.Key2:
		// TODO
	}
}
