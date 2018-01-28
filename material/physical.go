package material

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

type Physical struct {
}

func init() {
	demos.Map["material.physical"] = &Physical{}
}

func (t *Physical) Initialize(a *app.App) {

	// Add directional red light from right
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	l1.SetPosition(0.1, 0, 0)
	a.Scene().Add(l1)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1.0)
	l2.SetPosition(0, 0.1, 0)
	a.Scene().Add(l2)

	// Add directional blue light from front
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	l3.SetPosition(0, 0, 0.1)
	a.Scene().Add(l3)

	// Creates sphere 1
	geom1 := geometry.NewSphere(0.5, 16, 16, 0, math.Pi*2, 0, math.Pi)
	mat1 := material.NewPhysical()
	mat1.SetWireframe(false)
	mat1.SetSide(material.SideDouble)
	sphere := graphic.NewMesh(geom1, mat1)
	sphere.SetPosition(0, 0, 0)
	a.Scene().Add(sphere)

}

func (t *Physical) Render(a *app.App) {

}
