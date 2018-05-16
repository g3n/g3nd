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

type Torus struct {
	torus1  *graphic.Mesh
	normals *graphic.NormalsHelper
}

func init() {
	demos.Map["geometry.torus"] = &Torus{}
}

func (t *Torus) Initialize(a *app.App) {

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

	// Add torus at upper-left
	geom1 := geometry.NewTorus(1, 0.25, 8, 8, 2*math.Pi)
	mat1 := material.NewStandard(&math32.Color{0, 0, 0.5})
	t.torus1 = graphic.NewMesh(geom1, mat1)
	mat1.SetWireframe(true)
	mat1.SetSide(material.SideDouble)
	t.torus1.SetPosition(-2, 1.5, 0)
	a.Scene().Add(t.torus1)

	// Add torus at upper-right
	geom2 := geometry.NewTorus(1, 0.25, 32, 32, 2*math.Pi)
	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	torus2 := graphic.NewMesh(geom2, mat2)
	torus2.SetPosition(2, 1.5, 0)
	a.Scene().Add(torus2)

	// Add torus at bottom-left
	geom3 := geometry.NewTorus(1, 0.25, 32, 32, 2*math.Pi)
	mat3 := material.NewStandard(&math32.Color{0.5, 0.5, 0.5})
	torus3 := graphic.NewMesh(geom3, mat3)
	torus3.SetPosition(-2, -1.5, 0)
	a.Scene().Add(torus3)

	// Add torus at bottom-right
	geom4 := geometry.NewTorus(1, 0.25, 64, 64, 3*math.Pi/2)
	mat4 := material.NewStandard(&math32.Color{0.5, 0.5, 0.5})
	mat4.SetSide(material.SideDouble)
	torus4 := graphic.NewMesh(geom4, mat4)
	torus4.SetPosition(2, -1.5, 0)
	a.Scene().Add(torus4)

	// Adds axis helper
	axis := graphic.NewAxisHelper(2)
	a.Scene().Add(axis)

	// Adds normals helper
	t.normals = graphic.NewNormalsHelper(t.torus1, 0.5, &math32.Color{0, 1, 0})
	a.Scene().Add(t.normals)
}

func (t *Torus) Render(a *app.App) {

	// Rotate at 1 rotation each 5 seconds
	delta := a.FrameDeltaSeconds() * 2 * math32.Pi / 5
	t.torus1.AddRotationZ(delta)
	t.normals.Update()
}
