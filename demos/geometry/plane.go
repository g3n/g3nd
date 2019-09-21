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
)

func init() {
	app.DemoMap["geometry.plane"] = &Plane{}
}

type Plane struct {
	plane1 *graphic.Mesh
	plane2 *graphic.Mesh
	plane3 *graphic.Mesh
}

// Start is called once at the start of the demo.
func (t *Plane) Start(a *app.App) {

	// Adds directional RED light from +X
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1)
	l1.SetPosition(1, 0, 0)
	a.Scene().Add(l1)

	// Adds directional GREEN light from +Y
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1)
	l2.SetPosition(0, 1, 0)
	a.Scene().Add(l2)

	// Adds directional BLUE light from +Z
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1)
	l3.SetPosition(0, 0, 1)
	a.Scene().Add(l3)

	// Adds axis helper
	ah := util.NewAxisHelper(1.0)
	a.Scene().Add(ah)

	// Adds square plane, double sided at left
	plane1_geom := geometry.NewPlane(1, 1)
	plane1_mat := material.NewStandard(&math32.Color{1, 1, 1})
	plane1_mat.SetWireframe(false)
	plane1_mat.SetSide(material.SideDouble)
	t.plane1 = graphic.NewMesh(plane1_geom, plane1_mat)
	t.plane1.SetPositionX(-1)
	a.Scene().Add(t.plane1)

	// Adds rectangular plane, single sided at center
	plane2_geom := geometry.NewPlane(0.5, 1)
	plane2_mat := material.NewStandard(&math32.Color{1, 1, 1})
	plane2_mat.SetWireframe(false)
	plane2_mat.SetSide(material.SideFront)
	t.plane2 = graphic.NewMesh(plane2_geom, plane2_mat)
	a.Scene().Add(t.plane2)

	// Adds rectangular plane, single side at right
	plane3_geom := geometry.NewPlane(0.5, 1)
	plane3_mat := material.NewStandard(&math32.Color{1, 1, 1})
	plane3_mat.SetWireframe(false)
	plane3_mat.SetSide(material.SideBack)
	t.plane3 = graphic.NewMesh(plane3_geom, plane3_mat)
	t.plane3.SetPositionX(1)
	a.Scene().Add(t.plane3)

}

// Update is called every frame.
func (t *Plane) Update(a *app.App, deltaTime time.Duration) {

	// Rotate at 1 rotation each 5 seconds
	delta := float32(deltaTime.Seconds()) * 2 * math32.Pi / 5
	t.plane1.RotateY(delta)
	t.plane2.RotateX(delta)
	t.plane3.RotateX(-delta)
}

// Cleanup is called once at the end of the demo.
func (t *Plane) Cleanup(a *app.App) {}
