package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type Boxmulti struct {
	box *graphic.Mesh
}

func init() {
	TestMap["material.boxmulti"] = &Boxmulti{}
}

func (t *Boxmulti) Initialize(ctx *Context) {

	// Front directional light
	l1 := light.NewDirectional(math32.NewColor(0.4, 0.4, 0.4), 1.0)
	l1.SetPosition(0, 0, 1)
	ctx.Scene.Add(l1)

	// Axis helper
	axis := graphic.NewAxisHelper(1)
	ctx.Scene.Add(axis)

	// Creates box geometry
	geom := geometry.NewBox(1, 1, 1, 1, 1, 1)

	// Creates box materials
	mat0 := material.NewStandard(math32.NewColor(1, 0, 0))
	mat1 := material.NewStandard(math32.NewColor(0, 1, 0))
	mat2 := material.NewStandard(math32.NewColor(0, 0, 1))
	mat3 := material.NewStandard(math32.NewColor(1, 1, 0))
	mat4 := material.NewStandard(math32.NewColor(0, 1, 1))
	mat5 := material.NewStandard(math32.NewColor(1, 0, 1))

	// Creates box mesh and add materials for the groups
	t.box = graphic.NewMesh(geom, nil)
	t.box.AddGroupMaterial(mat0, 0)
	t.box.AddGroupMaterial(mat1, 1)
	t.box.AddGroupMaterial(mat2, 2)
	t.box.AddGroupMaterial(mat3, 3)
	t.box.AddGroupMaterial(mat4, 4)
	t.box.AddGroupMaterial(mat5, 5)
	ctx.Scene.Add(t.box)
}

func (t *Boxmulti) Render(ctx *Context) {

	t.box.AddRotationY(0.01)
}

