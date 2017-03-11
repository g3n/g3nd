package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"math"
)

type Texsphere struct {
	sphere1 *graphic.Mesh
	sphere2 *graphic.Mesh
	sphere3 *graphic.Mesh
	sphere4 *graphic.Mesh
}

func init() {
	TestMap["texture.sphere"] = &Texsphere{}
}

func (t *Texsphere) Initialize(ctx *Context) {

	// Adds directional front light
	dir1 := light.NewDirectional(math32.NewColor(1, 1, 1), 1.0)
	dir1.SetPosition(0, 0, 100)
	ctx.Scene.Add(dir1)

	// Adds directional top light
	dir2 := light.NewDirectional(math32.NewColor(1, 1, 1), 1.0)
	dir2.SetPosition(0, 100, 0)
	ctx.Scene.Add(dir2)

	// Creates texture 1
	texfile := ctx.DirData + "/images/checkerboard.jpg"
	tex1, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	tex1.SetWrapS(gls.REPEAT)
	tex1.SetWrapT(gls.REPEAT)
	tex1.SetRepeat(2, 2)
	// Creates sphere 1
	geom1 := geometry.NewSphere(1, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat1 := material.NewStandard(math32.NewColor(1, 1, 1))
	mat1.AddTexture(tex1)
	t.sphere1 = graphic.NewMesh(geom1, mat1)
	t.sphere1.SetPosition(-1.1, 1.1, 0)
	ctx.Scene.Add(t.sphere1)

	// Creates texture 2
	texfile = ctx.DirData + "/images/earth_clouds_big.jpg"
	tex2, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	tex2.SetFlipY(false)
	// Creates sphere 2
	geom2 := geometry.NewSphere(1, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat2 := material.NewPhong(math32.NewColor(1, 1, 1))
	mat2.AddTexture(tex2)
	t.sphere2 = graphic.NewMesh(geom2, mat2)
	t.sphere2.SetPosition(1.1, 1.1, 0)
	ctx.Scene.Add(t.sphere2)

	// Creates texture 3
	texfile = ctx.DirData + "/images/uvgrid.jpg"
	tex3, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	tex3.SetFlipY(false)
	// Creates sphere 3
	geom3 := geometry.NewSphere(1, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat3 := material.NewStandard(math32.NewColor(1, 1, 1))
	mat3.AddTexture(tex3)
	t.sphere3 = graphic.NewMesh(geom3, mat3)
	t.sphere3.SetPosition(-1.1, -1.1, 0)
	ctx.Scene.Add(t.sphere3)

	// Creates texture 4
	texfile = ctx.DirData + "/images/brick1.jpg"
	tex4, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		log.Fatal("Error loading texture: %s", err)
	}
	// Creates sphere 4
	geom4 := geometry.NewSphere(1, 32, 32, 0, math.Pi, 0, math.Pi/2)
	mat4 := material.NewPhong(math32.NewColor(1, 1, 1))
	mat4.AddTexture(tex4)
	mat4.SetSide(material.SideDouble)
	t.sphere4 = graphic.NewMesh(geom4, mat4)
	t.sphere4.SetPosition(1.1, -1.1, 0)
	ctx.Scene.Add(t.sphere4)

	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)
}

func (t *Texsphere) Render(ctx *Context) {

	t.sphere1.AddRotationY(0.01)
	t.sphere2.AddRotationY(-0.01)
	t.sphere3.AddRotationY(0.01)
	t.sphere4.AddRotationY(-0.01)
}
