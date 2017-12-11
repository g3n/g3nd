package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

type Blending struct {
	texbg *texture.Texture2D
}

func init() {
	TestMap["material.blending"] = &Blending{}
}

func (t *Blending) Initialize(ctx *Context) {

	ctx.CamPersp.SetPositionZ(600)
	ctx.AmbLight.SetIntensity(2)

	// Creates checker board textures for background
	c1 := &math32.Color{0.7, 0.7, 0.7}
	c2 := &math32.Color{0.3, 0.3, 0.3}
	t.texbg = texture.NewBoard(16, 16, c1, c2, c2, c1, 1)
	t.texbg.SetWrapS(gls.REPEAT)
	t.texbg.SetWrapT(gls.REPEAT)
	t.texbg.SetRepeat(64, 64)

	// Creates background plane
	matbg := material.NewStandard(&math32.Color{1, 1, 1})
	matbg.SetPolygonOffset(1, 1)
	matbg.AddTexture(t.texbg)
	geombg := geometry.NewPlane(4000, 3000, 1, 1)
	meshbg := graphic.NewMesh(geombg, matbg)
	meshbg.SetPosition(0, 0, -1)
	ctx.Scene.Add(meshbg)

	// Builds list of textures
	texnames := []string{
		"uvgrid.jpg", "sprite0.jpg",
		"sprite0.png", "lensflare0.png",
		"lensflare0_alpha.png",
	}
	texlist := []*texture.Texture2D{}
	for _, tname := range texnames {
		tex, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/" + tname)
		if err != nil {
			log.Fatal("Error loading texture: %s", err)
		}
		texlist = append(texlist, tex)
	}

	blendings := []struct {
		blending string
		value    material.Blending
	}{
		{"NoBlending", material.BlendingNone},
		{"NormalBlending", material.BlendingNormal},
		{"AdditiveBlending", material.BlendingAdditive},
		{"SubtractiveBlending", material.BlendingSubtractive},
		{"MultiplyBlending", material.BlendingMultiply},
	}

	// This geometry will be shared by several meshes
	// For each mesh which uses this geometry we need to increment its refcount
	geo1 := geometry.NewPlane(100, 100, 1, 1)
	defer geo1.Dispose()

	// Internal function go generate a row of images
	var addImageRow = func(tex *texture.Texture2D, y int) {
		for i := 0; i < len(blendings); i++ {
			material := material.NewPhong(&math32.Color{1, 1, 1})
			material.SetOpacity(1)
			material.AddTexture(tex)
			material.SetBlending(blendings[i].value)
			x := (float32(i) - float32(len(blendings))/2) * 110
			mesh := graphic.NewMesh(geo1.Incref(), material)
			mesh.SetPosition(x, float32(y), 0)
			ctx.Scene.Add(mesh)
		}
	}

	addImageRow(texlist[0], 300)
	addImageRow(texlist[1], 150)
	addImageRow(texlist[2], 0)
	addImageRow(texlist[3], -150)
	addImageRow(texlist[4], -300)
}

func (t *Blending) Render(ctx *Context) {

	time := float32(ctx.Win.GetTime() * 0.003)
	rx, ry := t.texbg.Repeat()
	ox := math32.Mod(time*rx, 1.0)
	oy := math32.Mod(time*ry, 1.0)
	t.texbg.SetOffset(-ox, oy)
}
