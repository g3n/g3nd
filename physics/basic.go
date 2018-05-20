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
)

func init() {
	demos.Map["physics.basic"] = &PhysicsBasic{}
}

type PhysicsBasic struct {
	sim *physics.Simulation
	rb  *physics.Particle
}

func (t *PhysicsBasic) Initialize(a *app.App) {

	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)

	pl := light.NewPoint(math32.NewColor("white"), 1.0)
	pl.SetPosition(1, 0, 0)
	a.Scene().Add(pl)

	t.sim = physics.NewSimulation()
	gravity := physics.NewConstantForceField(&math32.Vector3{0,-9.8,0})
	//gravity := physics.NewAttractorForceField(&math32.Vector3{0.1,1,0}, 1)
	t.sim.AddForceField(gravity)

	// Creates sphere 1
	sphereGeom := geometry.NewSphere(0.1, 16, 16, 0, math.Pi*2, 0, math.Pi)
	mat := material.NewPhong(&math32.Color{1, 1, 1})
	sphere1 := graphic.NewMesh(sphereGeom, mat)
	a.Scene().Add(sphere1)

	t.rb = physics.NewParticle(sphere1)
	t.sim.AddParticle(t.rb)

}

func (t *PhysicsBasic) Render(a *app.App) {

	t.sim.Step(a.FrameDelta())
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
