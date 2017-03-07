package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"math"
)

func init() {
	TestMap["texture.cylinder"] = &TextureCylinder{}
}

type TextureCylinder struct {
	mesh1 *graphic.Mesh
	mesh2 *graphic.Mesh
	mesh3 *graphic.Mesh
}

func (t *TextureCylinder) Initialize(ctx *Context) {

	// Add directional red light from right
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	l1.SetPosition(1, 0, 0)
	ctx.Scene.Add(l1)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1.0)
	l2.SetPosition(0, 1, 0)
	ctx.Scene.Add(l2)

	// Add directional blue light from front
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	l3.SetPosition(0, 0, 1)
	ctx.Scene.Add(l3)

	// Left cylinder
	tex, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/brick1.jpg")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	geom1 := geometry.NewCylinder(0.8, 0.8, 2, 16, 2, 0, 2*math.Pi, true, true)
	mat1 := material.NewPhong(math32.NewColor(1, 1, 1))
	mat1.SetSide(material.SideDouble)
	mat1.AddTexture(tex)
	t.mesh1 = graphic.NewMesh(geom1, mat1)
	t.mesh1.SetPosition(-2, 0, 0)
	ctx.Scene.Add(t.mesh1)

	// Middle cylinder
	tex, err = texture.NewTexture2DFromImage(ctx.DirData + "/images/moss.png")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	geom2 := geometry.NewCylinder(0.8, 0.8, 2, 32, 16, 0, 2*math.Pi, false, true)
	mat2 := material.NewPhong(math32.NewColor(1, 1, 1))
	mat2.SetSide(material.SideDouble)
	mat2.AddTexture(tex)
	t.mesh2 = graphic.NewMesh(geom2, mat2)
	t.mesh2.SetPosition(0, 0, 0)
	ctx.Scene.Add(t.mesh2)

	// Right cylinder
	tex, err = texture.NewTexture2DFromImage(ctx.DirData + "/images/checkerboard.jpg")
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	geom3 := geometry.NewCylinder(0.4, 0.8, 2, 32, 1, 0, 2*math.Pi, false, true)
	mat3 := material.NewStandard(math32.NewColor(1, 1, 1))
	mat3.SetSide(material.SideDouble)
	mat3.AddTexture(tex)
	t.mesh3 = graphic.NewMesh(geom3, mat3)
	t.mesh3.SetPosition(2, 0, 0)
	ctx.Scene.Add(t.mesh3)

	// Adds axis helper
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)
}

func (t *TextureCylinder) Render(ctx *Context) {

	t.mesh1.AddRotationY(0.005)
	t.mesh2.AddRotationY(-0.004)
	t.mesh3.AddRotationY(0.003)
}

