package geometry

import (
	"math"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["geometry.cylinder"] = &Cylinder{}
}

type Cylinder struct {
	mesh    *graphic.Mesh
	normals *graphic.NormalsHelper
}

func (t *Cylinder) Initialize(a *app.App) {

	// Add directional red light from right
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	l1.SetPosition(1, 0, 0)
	a.Scene().Add(l1)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1.0)
	l2.SetPosition(0, 1, 0)
	a.Scene().Add(l2)

	// Add directional blue light from front
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	l3.SetPosition(0, 0, 1)
	a.Scene().Add(l3)

	// Left cylinder
	geom1 := geometry.NewCylinder(0.8, 0.8, 2, 16, 2, 0, 2*math.Pi, true, true)
	mat1 := material.NewPhong(&math32.Color{0, 1, 0})
	mat1.SetWireframe(true)
	mat1.SetSide(material.SideDouble)
	t.mesh = graphic.NewMesh(geom1, mat1)
	t.mesh.SetPosition(-2, 0, 0)
	a.Scene().Add(t.mesh)

	// Middle cylinder
	geom2 := geometry.NewCylinder(0.8, 0.8, 2, 32, 16, 0, 2*math.Pi, false, true)
	mat2 := material.NewPhong(&math32.Color{1, 1, 1})
	mat2.SetSide(material.SideDouble)
	mesh := graphic.NewMesh(geom2, mat2)
	mesh.SetPosition(0, 0, 0)
	a.Scene().Add(mesh)

	// Right cylinder
	geom3 := geometry.NewCylinder(0.4, 0.8, 2, 32, 1, 0, 2*math.Pi, false, true)
	mat3 := material.NewStandard(&math32.Color{1, 1, 1})
	mat3.SetSide(material.SideDouble)
	mesh3 := graphic.NewMesh(geom3, mat3)
	mesh3.SetPosition(2, 0, 0)
	a.Scene().Add(mesh3)

	// Adds axis helper
	axis := graphic.NewAxisHelper(2)
	a.Scene().Add(axis)

	// Adds normals helper
	t.normals = graphic.NewNormalsHelper(t.mesh, 0.5, &math32.Color{0, 1, 0}, 1)
	a.Scene().Add(t.normals)
}

func (t *Cylinder) Render(a *app.App) {

	t.mesh.RotateY(0.005)
	t.normals.Update()
}
