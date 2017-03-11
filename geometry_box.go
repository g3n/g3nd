package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type Box struct {
	box     *graphic.Mesh
	normals *graphic.NormalsHelper
}

func init() {
	TestMap["geometry.box"] = &Box{}
}

func (t *Box) Initialize(ctx *Context) {

	// Add box
	geom := geometry.NewBox(1, 1, 1, 2, 2, 2)
	mat := material.NewStandard(math32.NewColor(0.5, 0, 0))
	mat.SetWireframe(false)
	t.box = graphic.NewMesh(geom, mat)
	ctx.Scene.Add(t.box)

	// Add normals helper
	t.normals = graphic.NewNormalsHelper(t.box, 0.5, &math32.Color{0, 0, 1}, 1)
	ctx.Scene.Add(t.normals)

	// Adds directional light
	l1 := light.NewDirectional(math32.NewColor(0.4, 0.4, 0.4), 1.0)
	l1.SetPosition(0, 0, 1)
	ctx.Scene.Add(l1)
}

func (t *Box) Render(ctx *Context) {

	t.box.AddRotationY(0.01)
	t.normals.Update()
}
