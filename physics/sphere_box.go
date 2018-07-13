package physics

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/engine/window"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/physics"
	"github.com/g3n/engine/physics/object"
	"math"
)

func init() {
	demos.Map["physics (experimental).sphere_box"] = &PhysicsSphereBox{}
}

type PhysicsSphereBox struct {
	sim *physics.Simulation
	rb  *object.Body
	rb2 *object.Body
	rb3 *object.Body
}

func (t *PhysicsSphereBox) Initialize(a *app.App) {

	// Subscribe to key events
	a.Window().Subscribe(window.OnKeyRepeat, t.onKey)
	a.Window().Subscribe(window.OnKeyDown, t.onKey)

	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)

	pl := light.NewPoint(math32.NewColor("white"), 1.0)
	pl.SetPosition(1, 0, 1)
	a.Scene().Add(pl)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.3)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	t.sim = physics.NewSimulation(a.Scene())

	sphereGeom := geometry.NewSphere(0.1, 16, 16, 0, math.Pi*2, 0, math.Pi)
	cubeGeom := geometry.NewCube(0.2)
	mat := material.NewPhong(&math32.Color{1, 1, 1})
	mat.SetTransparent(true)
	mat.SetOpacity(0.5)

	sphere := graphic.NewMesh(sphereGeom, mat)
	sphere.SetPosition(2, 0, 0)
	a.Scene().Add(sphere)
	t.rb2 = object.NewBody(sphere)
	//t.rb2.SetLinearDamping(0)
	t.sim.AddBody(t.rb2, "Sphere")
	t.rb2.SetVelocity(math32.NewVector3(-0.5, 0, 0))

	cube := graphic.NewMesh(cubeGeom, mat)
	cube.SetPosition(0, 0, 0)
	cube.SetRotation(0, math32.Pi*0.25, math32.Pi*0.25)
	a.Scene().Add(cube)
	t.rb3 = object.NewBody(cube)
	//t.rb3.SetLinearDamping(0)
	t.sim.AddBody(t.rb3, "Cube1")
	t.rb3.SetVelocity(math32.NewVector3(0.5, 0, 0))

}

func (t *PhysicsSphereBox) Render(a *app.App) {

	t.sim.Step(float32(a.FrameDelta().Seconds()))
}

func (t *PhysicsSphereBox) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	if kev.Action == window.Release {
		return
	}
	switch kev.Keycode {
	case window.KeyP:
		t.sim.SetPaused(!t.sim.Paused())
	case window.KeySpace:
		t.sim.SetPaused(false)
		t.sim.Step(0.016)
		t.sim.SetPaused(true)
	case window.Key1:
		t.rb2.ApplyVelocityDeltas(math32.NewVector3(-1, 0, 0), math32.NewVector3(0, 0, 1))
	case window.Key2:
		t.rb2.ApplyVelocityDeltas(math32.NewVector3(1, 0, 0), math32.NewVector3(0, 0, -1))
	}
}
