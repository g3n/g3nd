package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util"
	"github.com/g3n/g3nd/app"
	"time"

	"math"
)

func init() {
	app.DemoMap["geometry.sphere"] = &Sphere{}
}

type Sphere struct {
	sphere1 *graphic.Mesh
	sphere2 *graphic.Mesh
	normals *util.NormalsHelper
}

// Start is called once at the start of the demo.
func (t *Sphere) Start(a *app.App) {

	// Add directional red light from right
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	l1.SetPosition(0.1, 0, 0)
	a.Scene().Add(l1)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1.0)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	// Add directional blue light from front
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	l3.SetPosition(0, 0, 0.1)
	a.Scene().Add(l3)

	// Creates sphere 1
	geom1 := geometry.NewSphere(1, 16, 16, 0, math.Pi*2, 0, math.Pi)
	mat1 := material.NewStandard(&math32.Color{0, 0, 0.6})
	mat1.SetWireframe(true)
	mat1.SetSide(material.SideDouble)
	t.sphere1 = graphic.NewMesh(geom1, mat1)
	t.sphere1.SetPosition(-1.5, 0, 0)
	a.Scene().Add(t.sphere1)

	// Creates sphere 2
	geom2 := geometry.NewSphere(1, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	mat2.SetWireframe(false)
	mat2.SetSide(material.SideDouble)
	t.sphere2 = graphic.NewMesh(geom2, mat2)
	t.sphere2.SetPosition(1.5, 0, 0)
	a.Scene().Add(t.sphere2)

	// Add axis helper
	axis := util.NewAxisHelper(2)
	a.Scene().Add(axis)

	// Adds normals helper
	t.normals = util.NewNormalsHelper(t.sphere1, 0.5, &math32.Color{0, 1, 0}, 1)
	a.Scene().Add(t.normals)
}

// Update is called every frame.
func (t *Sphere) Update(a *app.App, deltaTime time.Duration) {

	t.sphere1.RotateY(0.005)
	t.normals.Update()
	t.sphere2.RotateY(-0.005)
}

// Cleanup is called once at the end of the demo.
func (t *Sphere) Cleanup(a *app.App) {}
