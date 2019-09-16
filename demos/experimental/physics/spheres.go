package physics

import (
	"github.com/g3n/engine/experimental/collision/shape"
	"github.com/g3n/engine/experimental/physics"
	"github.com/g3n/engine/experimental/physics/object"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util"
	"github.com/g3n/engine/window"
	"github.com/g3n/g3nd/app"
	"math"
	"time"
)

func init() {
	app.DemoMap["physics-experimental.spheres"] = &PhysicsSpheres{}
}

type PhysicsSpheres struct {
	sim *physics.Simulation
	app *app.App

	sphereGeom *geometry.Sphere
	matSphere  *material.Standard

	anim   *texture.Animator
	sprite *graphic.Sprite

	attractorOn bool
	gravity     *physics.ConstantForceField
	attractor   *physics.AttractorForceField
}

// Start is called once at the start of the demo.
func (t *PhysicsSpheres) Start(a *app.App) {

	t.app = a

	// Subscribe to key events
	a.Subscribe(window.OnKeyRepeat, t.onKey)
	a.Subscribe(window.OnKeyDown, t.onKey)

	//a.Camera().GetCamera().SetPosition
	// LookAt

	axis := util.NewAxisHelper(1)
	a.Scene().Add(axis)

	pl := light.NewPoint(math32.NewColor("white"), 1.0)
	pl.SetPosition(0, 1, 0)
	a.Scene().Add(pl)

	// Add directional light from top
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.5)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	// Create simulation and force fields
	t.sim = physics.NewSimulation(a.Scene())
	t.gravity = physics.NewConstantForceField(&math32.Vector3{0, -0.98, 0})
	t.attractor = physics.NewAttractorForceField(&math32.Vector3{0, 1, 0}, 1)
	t.sim.AddForceField(t.gravity)

	// Create sprite texture and animator
	tex2, err := texture.NewTexture2DFromImage(a.DirData() + "/images/smoke30.png")
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}
	t.anim = texture.NewAnimator(tex2, 6, 5)
	t.anim.SetDispTime(2 * 16666 * time.Microsecond)
	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	mat2.AddTexture(tex2)
	mat2.SetOpacity(0.5)
	mat2.SetTransparent(true)
	t.sprite = graphic.NewSprite(2, 2, mat2)
	t.sprite.SetPosition(0, 1, 0)
	t.sprite.SetVisible(false)
	a.Scene().Add(t.sprite)

	// Create sphere geometry
	t.sphereGeom = geometry.NewSphere(0.1, 16, 16, 0, math.Pi*2, 0, math.Pi)

	texfileG := a.DirData() + "/images/ground2.jpg"
	texG, err := texture.NewTexture2DFromImage(texfileG)
	texG.SetRepeat(100, 100)
	texG.SetWrapS(gls.REPEAT)
	texG.SetWrapT(gls.REPEAT)
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}

	mat := material.NewPhong(&math32.Color{1, 1, 1})
	mat.SetTransparent(true)
	mat.SetOpacity(0.5)
	mat.AddTexture(texG)

	floorGeom := geometry.NewPlane(100, 100, 1, 1)
	floor := graphic.NewMesh(floorGeom, mat)
	floor.SetPosition(0, 0, 0)
	floor.SetRotation(-math32.Pi/2, 0, 0)
	a.Scene().Add(floor)
	floorBody := object.NewBody(floor)
	floorBody.SetShape(shape.NewPlane())
	floorBody.SetBodyType(object.Static)
	t.sim.AddBody(floorBody, "Floor")

	// Create sphere texture
	texfile := a.DirData() + "/images/uvgrid.jpg"
	tex3, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		a.Log().Fatal("Error loading texture: %s", err)
	}

	// Create sphere material
	t.matSphere = material.NewStandard(&math32.Color{1, 1, 1})
	t.matSphere.AddTexture(tex3)

	sphere2 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere2.SetPosition(0, 1, -0.02)
	a.Scene().Add(sphere2)
	rb2 := object.NewBody(sphere2)
	rb2.SetShape(shape.NewSphere(0.1))
	t.sim.AddBody(rb2, "Sphere2")

	sphere3 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere3.SetPosition(0.05, 1.2, 0.05)
	a.Scene().Add(sphere3)
	rb3 := object.NewBody(sphere3)
	rb3.SetShape(shape.NewSphere(0.1))
	t.sim.AddBody(rb3, "Sphere3")

	sphere4 := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere4.SetPosition(-0.05, 1.4, 0)
	a.Scene().Add(sphere4)
	rb4 := object.NewBody(sphere4)
	rb4.SetShape(shape.NewSphere(0.1))
	t.sim.AddBody(rb4, "Sphere4")
}

func (t *PhysicsSpheres) ThrowBall() {

	// Obtain throw direction from camera position and target
	camPos := t.app.Camera().Position()
	camTarget := t.app.Orbit().Target()
	throwDir := math32.NewVec3().SubVectors(&camTarget, &camPos).SetLength(3)

	// Create sphere rigid body
	sphere := graphic.NewMesh(t.sphereGeom, t.matSphere)
	sphere.SetPositionVec(&camPos)
	t.app.Scene().Add(sphere)
	rb := object.NewBody(sphere)
	rb.SetShape(shape.NewSphere(0.1))
	rb.SetVelocity(throwDir)
	t.sim.AddBody(rb, "Sphere")
}

func (t *PhysicsSpheres) onKey(evname string, ev interface{}) {

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
	case window.KeyA:
		if t.attractorOn {
			t.sim.AddForceField(t.gravity)
			t.sim.RemoveForceField(t.attractor)
			t.sprite.SetVisible(false)
			t.attractorOn = false
		} else {
			t.sim.RemoveForceField(t.gravity)
			t.sim.AddForceField(t.attractor)
			t.sprite.SetVisible(true)
			t.attractorOn = true
		}
	case window.Key2:
		// TODO
	}
}

// Update is called every frame.
func (t *PhysicsSpheres) Update(a *app.App, deltaTime time.Duration) {

	t.sim.Step(float32(deltaTime.Seconds()))
	t.anim.Update(time.Now())
}

// Cleanup is called once at the end of the demo.
func (t *PhysicsSpheres) Cleanup(a *app.App) {}
