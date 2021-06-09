package geometry

import (
	"time"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["geometry.box"] = &Box{}
}

type Box struct {
	box     *graphic.Mesh
	normals *helper.Normals
}

// Start is called once at the start of the demo.
func (t *Box) Start(a *app.App) {

	// Add box
	geom := geometry.NewSegmentedCube(1, 2)
	mat := material.NewStandard(&math32.Color{0.5, 0, 0})
	mat.SetWireframe(false)
	t.box = graphic.NewMesh(geom, mat)
	a.Scene().Add(t.box)

	// Add normals helper
	t.normals = helper.NewNormals(t.box, 0.5, &math32.Color{0, 0, 1}, 1)
	a.Scene().Add(t.normals)

	// Adds directional light
	l1 := light.NewDirectional(&math32.Color{0.4, 0.4, 0.4}, 1.0)
	l1.SetPosition(0, 0, 1)
	a.Scene().Add(l1)
}

// Update is called every frame.
func (t *Box) Update(a *app.App, deltaTime time.Duration) {

	// Rotate the box at 1 rotation each 4 seconds
	t.box.RotateY(float32(deltaTime.Seconds()) * 2 * math32.Pi / 4)
	t.normals.Update()
}

// Cleanup is called once at the end of the demo.
func (t *Box) Cleanup(a *app.App) {}
