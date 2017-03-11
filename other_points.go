package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"math/rand"
)

type Points2 struct {
	points *graphic.Points
}

func init() {
	TestMap["other.points"] = &Points2{}
}

func (t *Points2) Initialize(ctx *Context) {

	ctx.GS.ClearColor(0, 0, 0, 1)
	ctx.CamPersp.SetPositionZ(5)

	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	// Load textures for the sprites
	spnames := []string{
		"snowflake1.png", "snowflake2.png",
		"snowflake3.png", "snowflake4.png",
		"snowflake5.png",
	}
	sprites := []*texture.Texture2D{}
	for _, name := range spnames {
		tex, err := texture.NewTexture2DFromImage(ctx.DirData + "/images/" + name)
		if err != nil {
			log.Fatal("Error loading texture: %s", err)
		}
		sprites = append(sprites, tex)
	}

	// Creates geometry with random points
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	numPoints := 10000
	coord := float32(10)
	for i := 0; i < numPoints; i++ {
		var vertex math32.Vector3
		vertex.Set(
			rand.Float32()*coord-coord/2,
			rand.Float32()*coord-coord/2,
			rand.Float32()*coord-coord/2,
		)
		positions.AppendVector3(&vertex)
	}
	geom.AddVBO(gls.NewVBO().AddAttrib("VertexPosition", 3).SetBuffer(positions))

	t.points = graphic.NewPoints(geom, nil)
	start := 0
	count := numPoints / len(sprites)
	for _, tex := range sprites {
		mat := material.NewPoint(math32.NewColor(1, 1, 1))
		mat.AddTexture(tex)
		mat.SetSize(1000)
		mat.SetBlending(material.BlendingAdditive)
		mat.SetDepthTest(false)
		t.points.AddMaterial(t.points, mat, start, count)
		start += count
	}
	ctx.Scene.Add(t.points)
}

func (t *Points2) Render(ctx *Context) {

	t.points.AddRotationY(0.005)
}
