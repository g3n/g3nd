package other

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
	"time"

	"math/rand"
)

func init() {
	app.DemoMap["other.points"] = &Points2{}
}

type Points2 struct {
	points *graphic.Points
}

// Start is called once at the start of the demo.
func (t *Points2) Start(a *app.App) {

	a.Gls().ClearColor(0, 0, 0, 1)

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)

	// Load textures for the sprites
	spnames := []string{"snowflake1.png", "snowflake2.png", "snowflake3.png", "snowflake4.png", "snowflake5.png"}
	sprites := []*texture.Texture2D{}
	for _, name := range spnames {
		tex, err := texture.NewTexture2DFromImage(a.DirData() + "/images/" + name)
		if err != nil {
			a.Log().Fatal("Error loading texture: %s", err)
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
	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))

	t.points = graphic.NewPoints(geom, nil)
	start := 0
	count := numPoints / len(sprites)
	for _, tex := range sprites {
		mat := material.NewPoint(&math32.Color{1, 1, 1})
		mat.SetTransparent(true)
		mat.SetOpacity(0.6)
		mat.AddTexture(tex)
		mat.SetSize(60)
		mat.SetBlending(material.BlendAdditive)
		mat.SetDepthMask(false)
		t.points.AddMaterial(t.points, mat, start, count)
		start += count
	}
	a.Scene().Add(t.points)
}

// Update is called every frame.
func (t *Points2) Update(a *app.App, deltaTime time.Duration) {

	rps := float32(deltaTime.Seconds()) * 2 * math32.Pi
	t.points.RotateY(rps * 0.05)
}

// Cleanup is called once at the end of the demo.
func (t *Points2) Cleanup(a *app.App) {}
