package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func init() {
	TestMap["geometry.lines"] = &Lines{}
}

type Lines struct{}

func (t *Lines) Initialize(ctx *Context) {

	// Creates geometry
	geom := geometry.NewGeometry()
	vertices := math32.NewArrayF32(0, 16)
	vertices.Append(
		-0.5, 0.0, 0.0,
		0.5, 0.0, 0.0,
		0.0, -0.5, 0.0,
		0.0, 0.5, 0.0,
		0.0, 0.0, -0.5,
		0.0, 0.0, 0.5,
	)
	colors := math32.NewArrayF32(0, 16)
	colors.Append(
		1.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
	)
	geom.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(vertices))
	geom.AddVBO(gls.NewVBO().AddAttrib("VertexColor", 3).SetBuffer(colors))

	// Creates basic material
	mat := material.NewBasic()
	mat.SetLineWidth(2.0)

	// Creates lines with the specified geometry and material
	lines1 := graphic.NewLines(geom, mat)
	ctx.Scene.Add(lines1)
}

func (t *Lines) Render(ctx *Context) {
}
