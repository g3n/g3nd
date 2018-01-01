package geometry

import (
	"math"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["geometry.circle"] = &Circle{}
}

type Circle struct {
	circ1 *graphic.Mesh
	circ2 *graphic.Mesh
	circ3 *graphic.Mesh
}

func (t *Circle) Initialize(app *g3nd.App) {

	geom1 := geometry.NewCircle(1, 3, 0, 2*math.Pi)
	mat1 := material.NewStandard(&math32.Color{0, 1, 0})
	mat1.SetWireframe(true)
	t.circ1 = graphic.NewMesh(geom1, mat1)
	t.circ1.SetPosition(-1.5, 0, 0)
	app.Scene().Add(t.circ1)

	geom2 := geometry.NewCircle(1, 50, 0, 2*math.Pi)
	mat2 := material.NewStandard(&math32.Color{0, 0, 1})
	t.circ2 = graphic.NewMesh(geom2, mat2)
	t.circ2.SetPosition(0, 0, -0.5)
	app.Scene().Add(t.circ2)

	geom3 := geometry.NewCircle(1, 5, 0, 2*math.Pi)
	mat3 := material.NewStandard(&math32.Color{1, 0, 0})
	t.circ3 = graphic.NewMesh(geom3, mat3)
	t.circ3.SetPosition(1.5, 0, -1.0)
	app.Scene().Add(t.circ3)
}

func (t *Circle) Render(app *g3nd.App) {

	// Rotate at 1 rotation each 5 seconds
	delta := app.FrameDeltaSeconds() * 2 * math32.Pi / 5
	t.circ1.AddRotationZ(delta)
	t.circ3.AddRotationZ(-delta)
}
