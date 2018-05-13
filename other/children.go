package other

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"

)

func init() {
	demos.Map["other.children"] = &Children{}
}

type Children struct {
	circ1 *graphic.Mesh
	circ2 *graphic.Mesh
	circ3 *graphic.Mesh
	count float32
}

func (t *Children) Initialize(a *app.App) {

	t.circ1 = graphic.NewMesh(
		geometry.NewCircle(0.5, 3),
		material.NewStandard(&math32.Color{0, 1, 0}),
	)
	t.circ1.SetPositionX(0)
	a.Scene().Add(t.circ1)

	t.circ2 = graphic.NewMesh(
		geometry.NewCircle(0.3, 5),
		material.NewStandard(&math32.Color{0, 0, 1}),
	)
	t.circ2.SetPositionX(1.4)
	t.circ1.Add(t.circ2)

	t.circ3 = graphic.NewMesh(
		geometry.NewCircle(0.2, 6),
		material.NewStandard(&math32.Color{1, 0, 0}),
	)
	t.circ3.SetPositionX(0.6)
	t.circ2.Add(t.circ3)
}

func (t *Children) Render(a *app.App) {

	// Rotate at 1 rotation each 5 seconds
	delta := a.FrameDeltaSeconds() * 2 * math32.Pi / 5
	t.circ1.AddRotationZ(delta)
	t.circ1.SetPositionX(math32.Sin(float32(t.count)))
	t.count += delta

	t.circ2.AddRotationZ(2 * delta)
	t.circ3.AddRotationZ(-4 * delta)
}
