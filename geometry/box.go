package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["geometry.box"] = &Box{}
}

type Box struct {
	box     *graphic.Mesh
	normals *graphic.NormalsHelper
}

func (t *Box) Initialize(a *app.App) {

	// Add box
	geom := geometry.NewSegmentedCube(1,2)
	mat := material.NewStandard(&math32.Color{0.5, 0, 0})
	mat.SetWireframe(false)
	t.box = graphic.NewMesh(geom, mat)
	a.Scene().Add(t.box)

	// Add normals helper
	t.normals = graphic.NewNormalsHelper(t.box, 0.5, &math32.Color{0, 0, 1}, 1)
	a.Scene().Add(t.normals)

	// Adds directional light
	l1 := light.NewDirectional(&math32.Color{0.4, 0.4, 0.4}, 1.0)
	l1.SetPosition(0, 0, 1)
	a.Scene().Add(l1)
}

func (t *Box) Render(a *app.App) {

	// Rotate the box at 1 rotation each 4 seconds
	t.box.RotateY(a.FrameDeltaSeconds() * 2 * math32.Pi / 4)
	t.normals.Update()
}
