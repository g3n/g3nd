package physics

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/engine/window"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/geometry"
	"math"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/physics"
	"github.com/g3n/engine/physics/object"
)

func init() {
	demos.Map["physics.spheres"] = &PhysicsSpheres{}
}

type PhysicsSpheres struct {
	sim *physics.Simulation
}

func (t *PhysicsSpheres) Initialize(a *app.App) {

	// Subscribe to key events
	a.Window().Subscribe(window.OnKeyRepeat, t.onKey)
	a.Window().Subscribe(window.OnKeyDown, t.onKey)

	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)

	pl := light.NewPoint(math32.NewColor("white"), 1.0)
	pl.SetPosition(1, 0, 1)
	a.Scene().Add(pl)

	t.sim = physics.NewSimulation(a.Scene())
	gravity := physics.NewConstantForceField(&math32.Vector3{0,-0.98,0})
	// //gravity := physics.NewAttractorForceField(&math32.Vector3{0.1,1,0}, 1)
	t.sim.AddForceField(gravity)

	// Creates sphere 1
	sphereGeom := geometry.NewSphere(0.1, 16, 16, 0, math.Pi*2, 0, math.Pi)
	mat := material.NewPhong(&math32.Color{1, 1, 1})
	mat.SetWireframe(true)

	//sphere1 := graphic.NewMesh(sphereGeom, mat)
	//a.Scene().Add(sphere1)
	//t.rb = object.NewBody(sphere1)
	//t.sim.AddBody(t.rb, "Sphere1")


	floorGeom := geometry.NewBox(0.5, 0.5, 0.5)
	floor := graphic.NewMesh(floorGeom, mat)
	floor.SetPosition(0,-0.2,0)
	a.Scene().Add(floor)
	rb3 := object.NewBody(floor)
	rb3.SetBodyType(object.Static)
	t.sim.AddBody(rb3, "Floor")

	sphere2 := graphic.NewMesh(sphereGeom, mat)
	sphere2.SetPosition(0, 1, 0)
	a.Scene().Add(sphere2)
	rb2 := object.NewBody(sphere2)
	t.sim.AddBody(rb2, "Sphere2")
	//rb2.SetVelocity(math32.NewVector3(-0.1, 0, 0))
}

func (t *PhysicsSpheres) Render(a *app.App) {

	t.sim.Step(float32(a.FrameDelta().Seconds()))
}

func (t *PhysicsSpheres) onKey(evname string, ev interface{}) {

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
		// TODO
	case window.Key2:
		// TODO
	}
}
