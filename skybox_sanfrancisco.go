package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

func init() {
	TestMap["skybox.sanfrancisco"] = &Skybox{}
}

type Skybox struct {
}

func (t *Skybox) Initialize(ctx *Context) {

	var textures = []string{
		"sanfrancisco/posx.jpg",
		"sanfrancisco/negx.jpg",
		"sanfrancisco/posy.jpg",
		"sanfrancisco/negy.jpg",
		"sanfrancisco/posz.jpg",
		"sanfrancisco/negz.jpg",
	}

	// Add axis helper
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	geom := geometry.NewBox(50, 50, 50, 2, 2, 2)
	skybox := graphic.NewMesh(geom, nil)
	for i := 0; i < 6; i++ {
		tex, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/" + textures[i])
		if err != nil {
			log.Fatal("Error loading texture: %s", err)
		}
		matFace := material.NewStandard(math32.NewColor(1, 1, 1))
		matFace.AddTexture(tex)
		matFace.SetSide(material.SideBack)
		skybox.AddGroupMaterial(matFace, i)
	}
	ctx.Scene.Add(skybox)
}

func (t *Skybox) Render(ctx *Context) {
}
