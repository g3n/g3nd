package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"

	"math"
)

type Sphere struct {
	sphere1 *graphic.Mesh
	sphere2 *graphic.Mesh
	normals *graphic.NormalsHelper
}

func init() {
	demos.Map["geometry.sphere"] = &Sphere{}
}

func (t *Sphere) Initialize(a *app.App) {

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
	axis := graphic.NewAxisHelper(2)
	a.Scene().Add(axis)

	// Adds normals helper
	t.normals = graphic.NewNormalsHelper(t.sphere1, 0.5, &math32.Color{0, 1, 0})
	a.Scene().Add(t.normals)
}

func (t *Sphere) Render(a *app.App) {

	t.sphere1.AddRotationY(0.005)
	t.normals.Update()
	t.sphere2.AddRotationY(-0.005)
}
