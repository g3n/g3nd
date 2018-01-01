package material

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/demos"
	"github.com/g3n/g3nd/g3nd"
)

type Boxmulti2 struct {
	box *graphic.Mesh
}

func init() {
	demos.Map["material.boxmulti2"] = &Boxmulti2{}
}

func (t *Boxmulti2) Initialize(app *g3nd.App) {

	// Front directional light
	l1 := light.NewDirectional(&math32.Color{0.4, 0.4, 0.4}, 1.0)
	l1.SetPosition(0, 0, 1)
	app.Scene().Add(l1)

	// Axis helper
	axis := graphic.NewAxisHelper(2)
	app.Scene().Add(axis)

	// Creates textures
	tex0, err := texture.NewTexture2DFromImage(app.DirData() + "/images/checkerboard.jpg")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}
	tex1, err := texture.NewTexture2DFromImage(app.DirData() + "/images/brick1.jpg")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}
	tex2, err := texture.NewTexture2DFromImage(app.DirData() + "/images/wall1.jpg")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}
	tex3, err := texture.NewTexture2DFromImage(app.DirData() + "/images/uvgrid.jpg")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}
	tex4, err := texture.NewTexture2DFromImage(app.DirData() + "/images/moss.png")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}
	tex5, err := texture.NewTexture2DFromImage(app.DirData() + "/images/tiger1.jpg")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}

	mat0 := material.NewStandard(&math32.Color{1, 1, 1})
	mat0.AddTexture(tex0)
	mat1 := material.NewStandard(&math32.Color{1, 1, 1})
	mat1.AddTexture(tex1)
	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	mat2.AddTexture(tex2)
	mat3 := material.NewStandard(&math32.Color{1, 1, 1})
	mat3.AddTexture(tex3)
	mat4 := material.NewStandard(&math32.Color{1, 1, 1})
	mat4.AddTexture(tex4)
	mat5 := material.NewStandard(&math32.Color{1, 1, 1})
	mat5.AddTexture(tex5)

	geom := geometry.NewBox(1, 1, 1, 2, 2, 2)
	t.box = graphic.NewMesh(geom, nil)
	t.box.AddGroupMaterial(mat0, 0)
	t.box.AddGroupMaterial(mat1, 1)
	t.box.AddGroupMaterial(mat2, 2)
	t.box.AddGroupMaterial(mat3, 3)
	t.box.AddGroupMaterial(mat4, 4)
	t.box.AddGroupMaterial(mat5, 5)

	app.Scene().Add(t.box)
}

func (t *Boxmulti2) Render(app *g3nd.App) {

	// Rotate at 1 rotation each 6 seconds
	delta := app.FrameDeltaSeconds() * 2 * math32.Pi / 6
	t.box.AddRotationY(delta)
}
