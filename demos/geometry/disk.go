package geometry

import (
	"time"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["geometry.disk"] = &Disk{}
}

type Disk struct {
	disk1 *graphic.Mesh
	disk2 *graphic.Mesh
	disk3 *graphic.Mesh
}

// Start is called once at the start of the demo.
func (t *Disk) Start(a *app.App) {

	geom1 := geometry.NewDisk(1, 3)
	mat1 := material.NewStandard(&math32.Color{0, 1, 0})
	mat1.SetWireframe(true)
	t.disk1 = graphic.NewMesh(geom1, mat1)
	t.disk1.SetPosition(-1.5, 0, 0)
	a.Scene().Add(t.disk1)

	geom2 := geometry.NewDisk(1, 50)
	mat2 := material.NewStandard(&math32.Color{0, 0, 1})
	t.disk2 = graphic.NewMesh(geom2, mat2)
	t.disk2.SetPosition(0, 0, -0.5)
	a.Scene().Add(t.disk2)

	geom3 := geometry.NewDisk(1, 5)
	mat3 := material.NewStandard(&math32.Color{1, 0, 0})
	t.disk3 = graphic.NewMesh(geom3, mat3)
	t.disk3.SetPosition(1.5, 0, -1.0)
	a.Scene().Add(t.disk3)
}

// Update is called every frame.
func (t *Disk) Update(a *app.App, deltaTime time.Duration) {

	// Rotate at 1 rotation each 5 seconds
	delta := float32(deltaTime.Seconds()) * 2 * math32.Pi / 5
	t.disk1.RotateZ(delta)
	t.disk3.RotateZ(-delta)
}

// Cleanup is called once at the end of the demo.
func (t *Disk) Cleanup(a *app.App) {}
