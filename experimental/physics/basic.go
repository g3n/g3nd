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
	"github.com/g3n/engine/experimental/physics"
	"github.com/g3n/engine/experimental/physics/object"
	"math"
	"github.com/g3n/engine/experimental/physics/shape"
)

func init() {
	demos.Map["physics-experimental.basic"] = &PhysicsBasic{}
}

type PhysicsBasic struct {
	sim *physics.Simulation
	rb  *object.Body
	rb2 *object.Body
	rb3 *object.Body
}

func (t *PhysicsBasic) Initialize(a *app.App) {

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
	//gravity := physics.NewConstantForceField(&math32.Vector3{0,-9.8,0})
	// //gravity := physics.NewAttractorForceField(&math32.Vector3{0.1,1,0}, 1)
	//t.sim.AddForceField(gravity)

	// Creates sphere 1
	sphereGeom := geometry.NewSphere(0.1, 16, 16, 0, math.Pi*2, 0, math.Pi)
	mat := material.NewPhong(&math32.Color{1, 1, 1})
	mat.SetWireframe(true)

	sphere1 := graphic.NewMesh(sphereGeom, mat)
	a.Scene().Add(sphere1)
	t.rb = object.NewBody(sphere1)
	t.rb.SetShape(shape.NewSphere(0.1))  // COMMENT THIS ON BOTH SPHERES TO TEST Convex-Convex collision!
	t.sim.AddBody(t.rb, "Sphere1")

	sphere2 := graphic.NewMesh(sphereGeom, mat)
	sphere2.SetPosition(1, 0, 0)
	a.Scene().Add(sphere2)
	t.rb2 = object.NewBody(sphere2)
	t.rb2.SetShape(shape.NewSphere(0.1)) // COMMENT THIS ON BOTH SPHERES TO TEST Convex-Convex collision!
	t.sim.AddBody(t.rb2, "Sphere2")
	t.rb2.SetVelocity(math32.NewVector3(-0.5, 0, 0))
	t.rb2.SetAngularVelocity(math32.NewVector3(0, 0, 1))

	//cubeGeom := geometry.NewCube(0.2)
	//cube1 := graphic.NewMesh(cubeGeom, mat)
	//a.Scene().Add(cube1)
	//t.rb3 = object.NewBody(cube1)
	//t.sim.AddBody(t.rb3, "Cube1")

}

func (t *PhysicsBasic) Render(a *app.App) {

	t.sim.Step(float32(a.FrameDelta().Seconds()))
}

func (t *PhysicsBasic) onKey(evname string, ev interface{}) {

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
