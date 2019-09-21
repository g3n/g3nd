package light

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	eutil "github.com/g3n/engine/util"
	"github.com/g3n/g3nd/app"
	"time"

	"github.com/g3n/g3nd/util"
	"math"
)

func init() {
	app.DemoMap["light.point"] = &PointLight{}
}

type PointLight struct {
	vl    *util.PointLightMesh
	hl    *util.PointLightMesh
	count float64
}

// Start is called once at the start of the demo.
func (t *PointLight) Start(a *app.App) {

	// Create spheres
	sphereGeom := geometry.NewSphere(0.5, 32, 16)
	sphere1 := graphic.NewMesh(sphereGeom, material.NewStandard(&math32.Color{0, 0, 0.6}))
	sphere1.SetPositionX(1)
	a.Scene().Add(sphere1)
	sphere2 := graphic.NewMesh(sphereGeom, material.NewStandard(&math32.Color{0, 0.5, 0.0}))
	sphere2.SetPositionX(-1)
	a.Scene().Add(sphere2)

	// Create box
	boxGeom := geometry.NewCube(4)
	boxMat := material.NewStandard(&math32.Color{1, 1, 1})
	boxMat.SetSide(material.SideBack)
	a.Scene().Add(graphic.NewMesh(boxGeom, boxMat))

	// Create axis helper
	axis := eutil.NewAxisHelper(1)
	a.Scene().Add(axis)

	// Create vertical point light
	t.vl = util.NewPointLightMesh(&math32.Color{1, 1, 1})
	a.Scene().Add(t.vl.Mesh)

	// Create horizontal point light
	t.hl = util.NewPointLightMesh(&math32.Color{1, 1, 1})
	a.Scene().Add(t.hl.Mesh)

	// Add controls
	if a.ControlFolder() == nil {
		return
	}
	g := a.ControlFolder().AddGroup("Show lights")
	cb1 := g.AddCheckBox("Horizontal").SetValue(t.hl.Mesh.Visible())
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.hl.Mesh.SetVisible(!t.hl.Mesh.Visible())
	})
	cb2 := g.AddCheckBox("Vertical").SetValue(t.vl.Mesh.Visible())
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.vl.Mesh.SetVisible(!t.vl.Mesh.Visible())
	})
}

// Update is called every frame.
func (t *PointLight) Update(a *app.App, deltaTime time.Duration) {

	t.vl.SetPosition(0, 1.5*float32(math.Sin(t.count)), 0)
	t.hl.SetPosition(1.5*float32(math.Sin(t.count)), 1, 0)
	t.count += 0.02 // TODO use deltaTime
}

// Cleanup is called once at the end of the demo.
func (t *PointLight) Cleanup(a *app.App) {}
