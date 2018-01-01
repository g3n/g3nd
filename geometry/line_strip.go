package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

func init() {
	demos.Map["geometry.line_strip"] = &LineStrip{}
}

type LineStrip struct{}

func (t *LineStrip) Initialize(app *g3nd.App) {

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
	geom1.AddVBO(gls.NewVBO().
		AddAttrib("VertexPosition", 3).
		AddAttrib("VertexColor", 3).
		SetBuffer(buffer),
	)

	// Creates basic material
	mat1 := material.NewBasic()
	mat1.SetLineWidth(2.0)

	// Creates line strip with the specified geometry and material
	lines1 := graphic.NewLineStrip(geom1, mat1)
	lines1.SetPosition(-0.6, 0, 0)
	app.Scene().Add(lines1)

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
	geom2.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(vertices))
	geom2.AddVBO(gls.NewVBO().AddAttrib("VertexColor", 3).SetBuffer(colors))

	// Creates basic material
	mat2 := material.NewBasic()
	mat2.SetLineWidth(2.0)

	// Creates line strip with the specified geometry and material
	lines2 := graphic.NewLineStrip(geom2, mat2)
	lines2.SetPosition(0.6, 0, 0)
	app.Scene().Add(lines2)
}

func (t *LineStrip) Render(app *g3nd.App) {
}
