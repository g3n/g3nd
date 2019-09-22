package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["geometry.points"] = &Points{}
}

type Points struct{}

// Start is called once at the start of the demo.
func (t *Points) Start(a *app.App) {

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)

	// Creates geometry
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	positions.Append(
		1, 1, 1,
		1, 1, -1,
		-1, 1, -1,
		-1, 1, 1,
		1, -1, 1,
		1, -1, -1,
		-1, -1, -1,
		-1, -1, 1,
		0, 1, 0,
		0, 0, -1,
		0, -1, 0,
		1, 0, 0,
		-1, 0, 0,
		0, 0, 1,
		0, 0, 0,
	)
	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	positions = nil // Positions cannot be used after transfering to VBO

	// Creates point material
	mat := material.NewPoint(&math32.Color{0, 0, 0})
	mat.SetSize(50)

	// Creates points mesh
	mesh := graphic.NewPoints(geom, mat)
	mesh.SetScale(1, 1, 1)
	a.Scene().Add(mesh)
}

// Update is called every frame.
func (t *Points) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *Points) Cleanup(a *app.App) {}
