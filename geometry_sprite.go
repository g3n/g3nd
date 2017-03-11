package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type Sprite struct {
	s1 *graphic.Sprite
	s2 *graphic.Sprite
	s3 *graphic.Sprite
	s4 *graphic.Sprite
}

func init() {
	TestMap["geometry.sprite"] = &Sprite{}
}

func (t *Sprite) Initialize(ctx *Context) {

	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	// Sprite 1
	mat1 := material.NewStandard(&math32.Color{1, 0, 0})
	t.s1 = graphic.NewSprite(1, 1, mat1)
	ctx.Scene.Add(t.s1)

	// Sprite 2
	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	mat2.SetWireframe(true)
	mat2.SetEmissiveColor(&math32.Color{0.5, 0.5, 0.5})
	t.s2 = graphic.NewSprite(1, 1, mat2)
	t.s2.SetPosition(2, 0, -2)
	ctx.Scene.Add(t.s2)

	// Sprite 3
	mat3 := material.NewStandard(&math32.Color{0, 0, 1})
	t.s3 = graphic.NewSprite(2, 1, mat3)
	t.s3.SetPosition(-2, 0, -3)
	ctx.Scene.Add(t.s3)

	// Sprite 4
	mat4 := material.NewStandard(&math32.Color{1, 1, 1})
	mat4.SetOpacity(1)
	t.s4 = graphic.NewSprite(1, 1, mat4)
	t.s4.SetPosition(0, 1, -1)
	t.s4.SetScale(2, 1, 1)
	ctx.Scene.Add(t.s4)
}

func (t *Sprite) Render(ctx *Context) {

	t.s1.AddRotationZ(-0.01)
	t.s3.AddRotationZ(0.01)
}
