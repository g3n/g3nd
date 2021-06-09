// TODO: UVs, Caps

package geometry

import (
	"math"
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
	app.DemoMap["geometry.tube"] = &Tube{}
}

type Tube struct {
	mesh *graphic.Mesh
}

// Start is called once at the start of the demo.
func (t *Tube) Start(a *app.App) {

	a.Camera().SetPosition(0, 10, 80)

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

	var path []math32.Vector3
	var i float64
	for i = 0; i <= 100; i++ {
		x := math.Cos(math.Pi*2*i/10) * 3
		y := i / 3
		z := math.Sin(math.Pi*2*i/10) * 3
		path = append(path, *math32.NewVector3(float32(x), float32(y), float32(z)))
	}
	c := geometry.NewTube(path, 1, 8, true)

	mat := material.NewStandard(&math32.Color{1, 1, 1})
	mat.SetSide(material.SideDouble)
	t.mesh = graphic.NewMesh(c, mat)
	a.Scene().Add(t.mesh)

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)
}

// Update is called every frame.
func (t *Tube) Update(a *app.App, deltaTime time.Duration) {
	t.mesh.RotateY(0.01)
}

// Cleanup is called once at the end of the demo.
func (t *Tube) Cleanup(a *app.App) {}
