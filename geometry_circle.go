package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"math"
)

func init() {
	TestMap["geometry.circle"] = &Circle{}
}

type Circle struct {
	circ1 *graphic.Mesh
	circ2 *graphic.Mesh
	circ3 *graphic.Mesh
}

func (t *Circle) Initialize(ctx *Context) {

	geom1 := geometry.NewCircle(1, 3, 0, 2*math.Pi)
	mat1 := material.NewStandard(&math32.Color{0, 1, 0})
	mat1.SetWireframe(true)
	t.circ1 = graphic.NewMesh(geom1, mat1)
	t.circ1.SetPosition(-1.5, 0, 0)
	ctx.Scene.Add(t.circ1)

	geom2 := geometry.NewCircle(1, 50, 0, 2*math.Pi)
	mat2 := material.NewStandard(&math32.Color{0, 0, 1})
	t.circ2 = graphic.NewMesh(geom2, mat2)
	t.circ2.SetPosition(0, 0, -0.5)
	ctx.Scene.Add(t.circ2)

	geom3 := geometry.NewCircle(1, 5, 0, 2*math.Pi)
	mat3 := material.NewStandard(&math32.Color{1, 0, 0})
	t.circ3 = graphic.NewMesh(geom3, mat3)
	t.circ3.SetPosition(1.5, 0, -1.0)
	ctx.Scene.Add(t.circ3)
}

func (t *Circle) Render(ctx *Context) {

	t.circ1.AddRotationZ(0.01)
	t.circ3.AddRotationZ(-0.01)
}
