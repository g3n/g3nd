package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["geometry.line_strip"] = &LineStrip{}
}

type LineStrip struct{}

// Start is called once at the start of the demo.
func (t *LineStrip) Start(a *app.App) {

	//
	// Creates line strip geometry using ONE interlaced buffer for vertices and colors
	//
	geom1 := geometry.NewGeometry()
	buffer := math32.NewArrayF32(0, 0)
	buffer.Append(
		0.0, 0.0, 0.0, 1.0, 0.0, 0.0,
		0.5, 0.0, 0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, -0.5, 0.0, 0.0, 1.0,
		0.0, 0.5, 0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 0.5, 0.0, 1.0, 0.0,
		-0.5, 0.0, 0.0, 0.0, 0.0, 1.0,
		0.0, -0.5, 0.0, 1.0, 0.0, 0.0,
	)
	geom1.AddVBO(gls.NewVBO(buffer).
		AddAttrib(gls.VertexPosition).
		AddAttrib(gls.VertexColor),
	)

	// Creates basic material
	mat1 := material.NewBasic()

	// Creates line strip with the specified geometry and material
	lines1 := graphic.NewLineStrip(geom1, mat1)
	lines1.SetPosition(-0.6, 0, 0)
	a.Scene().Add(lines1)

	//
	// Creates line strip geometry using TWO buffers: one for vertices and one for colors
	//
	geom2 := geometry.NewGeometry()
	vertices := math32.NewArrayF32(0, 32)
	vertices.Append(
		0.0, 0.0, 0.0,
		0.5, 0.0, 0.0,
		0.0, 0.0, -0.5,
		0.0, 0.5, 0.0,
		0.0, 0.0, 0.5,
		-0.5, 0.0, 0.0,
		0.0, -0.5, 0.0,
	)
	colors := math32.NewArrayF32(0, 32)
	colors.Append(
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
		1.0, 0.0, 0.0,
	)
	geom2.AddVBO(gls.NewVBO(vertices).AddAttrib(gls.VertexPosition))
	geom2.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))

	// Creates basic material
	mat2 := material.NewBasic()

	// Creates line strip with the specified geometry and material
	lines2 := graphic.NewLineStrip(geom2, mat2)
	lines2.SetPosition(0.6, 0, 0)
	a.Scene().Add(lines2)
}

// Update is called every frame.
func (t *LineStrip) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *LineStrip) Cleanup(a *app.App) {}
