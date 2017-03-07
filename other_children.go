package main

import (
	"math"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/graphic"
)

func init() {
	TestMap["other.children"] = &Children{}
}

type Children struct {
	circ1 *graphic.Mesh
	circ2 *graphic.Mesh
	circ3 *graphic.Mesh
	count float64
}

func (t *Children) Initialize(ctx *Context) {

	t.circ1 = graphic.NewMesh(
		geometry.NewCircle(0.5, 3, 0, 2*math.Pi),
		material.NewStandard(math32.NewColor(0,1,0)),
	)
	t.circ1.SetPositionX(0)
	ctx.Scene.Add(t.circ1)

	t.circ2 = graphic.NewMesh(
		geometry.NewCircle(0.3, 5, 0, 2*math.Pi),
		material.NewStandard(math32.NewColor(0,0,1)),
	)
	t.circ2.SetPositionX(1.4)
	t.circ1.Add(t.circ2)

	t.circ3 = graphic.NewMesh(
		geometry.NewCircle(0.2, 6, 0, 2*math.Pi),
		material.NewStandard(math32.NewColor(1,0,0)),
	)
	t.circ3.SetPositionX(0.6)
	t.circ2.Add(t.circ3)
}

func (t *Children) Render(ctx *Context) {

	t.circ1.AddRotationZ(0.01)
	t.circ1.SetPositionX(math32.Sin(float32(t.count)))
	t.count += 0.01

	t.circ2.AddRotationZ(0.02)
	t.circ3.AddRotationZ(-0.05)
}

