package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["geometry.plane"] = &Plane{}
}

type Plane struct {
	plane1 *graphic.Mesh
	plane2 *graphic.Mesh
	plane3 *graphic.Mesh
}

func (t *Plane) Initialize(app *g3nd.App) {

	// Adds directional RED light from +X
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1)
	l1.SetPosition(1, 0, 0)
	app.Scene().Add(l1)

	// Adds directional GREEN light from +Y
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1)
	l2.SetPosition(0, 1, 0)
	app.Scene().Add(l2)

	// Adds directional BLUE light from +Z
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1)
	l3.SetPosition(0, 0, 1)
	app.Scene().Add(l3)

	// Adds axis helper
	ah := graphic.NewAxisHelper(1.0)
	app.Scene().Add(ah)

	// Adds square plane, double sided at left
	plane1_geom := geometry.NewPlane(1, 1, 1, 1)
	plane1_mat := material.NewStandard(&math32.Color{1, 1, 1})
	plane1_mat.SetWireframe(false)
	plane1_mat.SetSide(material.SideDouble)
	t.plane1 = graphic.NewMesh(plane1_geom, plane1_mat)
	t.plane1.SetPositionX(-1)
	app.Scene().Add(t.plane1)

	// Adds rectangular plane, single sided at center
	plane2_geom := geometry.NewPlane(0.5, 1, 1, 1)
	plane2_mat := material.NewStandard(&math32.Color{1, 1, 1})
	plane2_mat.SetWireframe(false)
	plane2_mat.SetSide(material.SideFront)
	t.plane2 = graphic.NewMesh(plane2_geom, plane2_mat)
	app.Scene().Add(t.plane2)

	// Adds rectangular plane, single side at right
	plane3_geom := geometry.NewPlane(0.5, 1, 1, 1)
	plane3_mat := material.NewStandard(&math32.Color{1, 1, 1})
	plane3_mat.SetWireframe(false)
	plane3_mat.SetSide(material.SideBack)
	t.plane3 = graphic.NewMesh(plane3_geom, plane3_mat)
	t.plane3.SetPositionX(1)
	app.Scene().Add(t.plane3)

}

func (t *Plane) Render(app *g3nd.App) {

	// Rotate at 1 rotation each 5 seconds
	delta := app.FrameDeltaSeconds() * 2 * math32.Pi / 5
	t.plane1.AddRotationY(delta)
	t.plane2.AddRotationX(delta)
	t.plane3.AddRotationX(-delta)
}
