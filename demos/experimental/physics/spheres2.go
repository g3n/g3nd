package physics

import (
	"github.com/g3n/engine/experimental/physics"
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
	"math"
	"time"
)

func init() {
	app.DemoMap["physics-experimental.spheres2"] = &PhysicsSpheres2{}
}

type PhysicsSpheres2 struct {
	app        *app.App
	sim        *physics.Simulation
	sphereGeom *geometry.Sphere
	matSphere  *material.Standard
}

// Start is called once at the start of the demo.
func (t *PhysicsSpheres2) Start(a *app.App) {

	t.app = a

	// Subscribe to key events
	a.Subscribe(window.OnKeyRepeat, t.onKey)
	a.Subscribe(window.OnKeyDown, t.onKey)

	axis := graphic.NewAxisHelper(1)
	a.Scene().Add(axis)

	pl := light.NewPoint(math32.NewColor("white"), 1.0)
	pl.SetPosition(1, 0, 1)
	a.Scene().Add(pl)

	// Add directional light from top
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.3)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	// Add directional light from top
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.3)
	l3.SetPosition(0.1, 0, 0.1)
	a.Scene().Add(l3)

	t.sim = physics.NewSimulation(a.Scene())
	gravity := physics.NewConstantForceField(&math32.Vector3{0, -0.98, 0})
	// //gravity := physics.NewAttractorForceField(&math32.Vector3{0.1,1,0}, 1)
	t.sim.AddForceField(gravity)

	// Creates sphere 1
	t.sphereGeom = geometry.NewSphere(0.1, 16, 16, 0, math.Pi*2, 0, math.Pi)

	texfileG := a.DirData() + "/images/ground2.jpg"
	texG, err := texture.NewTexture2DFromImage(texfileG)
	texG.SetRepeat(10, 10)
	texG.SetWrapS(gls.REPEAT)
	texG.SetWrapT(gls.REPEAT)
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}

	mat := material.NewPhong(&math32.Color{1, 1, 1})
	mat.SetTransparent(true)
	mat.SetOpacity(0.5)
	mat.AddTexture(texG)
	//mat.SetWireframe(true)

	//sphere1 := graphic.NewMesh(sphereGeom, mat)
	//a.Scene().Add(sphere1)
	//t.rb = object.NewBody(sphere1)
	//t.sim.AddBody(t.rb, "Sphere1")

	floorGeom := geometry.NewBox(10, 0.5, 10)
	floor := graphic.NewMesh(floorGeom, mat)
	floor.SetPosition(3, -0.2, 0)
	a.Scene().Add(floor)
	floorBody := object.NewBody(floor)
	floorBody.SetBodyType(object.Static)
	t.sim.AddBody(floorBody, "Floor")

	// Creates texture 3
	texfile := a.DirData() + "/images/uvgrid.jpg"
	tex3, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	//tex3.SetFlipY(false)
	// Creates sphere 3
	t.matSphere = material.NewStandard(&math32.Color{1, 1, 1})
	t.matSphere.AddTexture(tex3)

	sphere2 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere2.SetPosition(0, 1, -0.02)
	a.Scene().Add(sphere2)
	rb2 := object.NewBody(sphere2)
	t.sim.AddBody(rb2, "Sphere2")

	sphere3 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere3.SetPosition(0.05, 1.2, 0.05)
	a.Scene().Add(sphere3)
	rb3 := object.NewBody(sphere3)
	t.sim.AddBody(rb3, "Sphere3")

	sphere4 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere4.SetPosition(-0.05, 1.4, 0)
	a.Scene().Add(sphere4)
	rb4 := object.NewBody(sphere4)
	t.sim.AddBody(rb4, "Sphere4")
}

func (t *PhysicsSpheres2) ThrowBall() {

	camPos := t.app.Camera().GetCamera().Position()
	camTarget := t.app.Camera().GetCamera().Target()
	throwDir := math32.NewVec3().SubVectors(&camTarget, &camPos).SetLength(3)

	sphere := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere.SetPositionVec(&camPos)
	t.app.Scene().Add(sphere)
	rb := object.NewBody(sphere)
	rb.SetVelocity(throwDir)
	t.sim.AddBody(rb, "Sphere4")
}

func (t *PhysicsSpheres2) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	switch kev.Key {
	case window.KeyP:
		t.sim.SetPaused(!t.sim.Paused())
	case window.KeyO:
		t.sim.SetPaused(false)
		t.sim.Step(0.016)
		t.sim.SetPaused(true)
	case window.KeySpace:
		t.ThrowBall()
	case window.Key1:
		// TODO
	case window.Key2:
		// TODO
	}
}

// Update is called every frame.
func (t *PhysicsSpheres2) Update(a *app.App, deltaTime time.Duration) {

	t.sim.Step(float32(deltaTime.Seconds()))
}

// Cleanup is called once at the end of the demo.
func (t *PhysicsSpheres2) Cleanup(a *app.App) {}
